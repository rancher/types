package v1

import (
	"context"
	"sync"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Interface interface {
	RESTClient() rest.Interface
	controller.Starter

	ClustersGetter
	ClusterNodesGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	clusterControllers     map[string]ClusterController
	clusterNodeControllers map[string]ClusterNodeController
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

		clusterControllers:     map[string]ClusterController{},
		clusterNodeControllers: map[string]ClusterNodeController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

func (c *Client) Sync(ctx context.Context) error {
	return controller.Sync(ctx, c.starters...)
}

func (c *Client) Start(ctx context.Context, threadiness int) error {
	return controller.Start(ctx, threadiness, c.starters...)
}

type ClustersGetter interface {
	Clusters(namespace string) ClusterInterface
}

func (c *Client) Clusters(namespace string) ClusterInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ClusterResource, ClusterGroupVersionKind, clusterFactory{})
	return &clusterClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ClusterNodesGetter interface {
	ClusterNodes(namespace string) ClusterNodeInterface
}

func (c *Client) ClusterNodes(namespace string) ClusterNodeInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ClusterNodeResource, ClusterNodeGroupVersionKind, clusterNodeFactory{})
	return &clusterNodeClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
