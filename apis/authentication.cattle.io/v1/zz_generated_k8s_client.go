package v1

import (
	"sync"

	"github.com/rancher/norman/clientbase"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Interface interface {
	RESTClient() rest.Interface

	TokensGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface

	tokenControllers map[string]TokenController
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

		tokenControllers: map[string]TokenController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

type TokensGetter interface {
	Tokens(namespace string) TokenInterface
}

func (c *Client) Tokens(namespace string) TokenInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &TokenResource, TokenGroupVersionKind, tokenFactory{})
	return &tokenClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
