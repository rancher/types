package v1

import (
	"sync"

	"github.com/rancher/norman/clientbase"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Interface interface {
	RESTClient() rest.Interface

	WorkloadsGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface

	workloadControllers map[string]WorkloadController
}

func NewForConfig(config rest.Config) (Interface, error) {
	if config.NegotiatedSerializer == nil {
		configConfig := dynamic.ContentConfig()
		config.NegotiatedSerializer = configConfig.NegotiatedSerializer
	}

	restClient, err := rest.UnversionedRESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{
		restClient: restClient,

		workloadControllers: map[string]WorkloadController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

type WorkloadsGetter interface {
	Workloads(namespace string) WorkloadInterface
}

func (c *Client) Workloads(namespace string) WorkloadInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &WorkloadResource, WorkloadGroupVersionKind, workloadFactory{})
	return &workloadClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
