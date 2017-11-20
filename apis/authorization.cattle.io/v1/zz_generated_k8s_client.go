package v1

import (
	"sync"

	"github.com/rancher/norman/clientbase"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Interface interface {
	RESTClient() rest.Interface

	ProjectsGetter
	ProjectRoleTemplatesGetter
	PodSecurityPolicyTemplatesGetter
	ProjectRoleTemplateBindingsGetter
	ClusterRoleTemplatesGetter
	ClusterRoleTemplateBindingsGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface

	projectControllers                    map[string]ProjectController
	projectRoleTemplateControllers        map[string]ProjectRoleTemplateController
	podSecurityPolicyTemplateControllers  map[string]PodSecurityPolicyTemplateController
	projectRoleTemplateBindingControllers map[string]ProjectRoleTemplateBindingController
	clusterRoleTemplateControllers        map[string]ClusterRoleTemplateController
	clusterRoleTemplateBindingControllers map[string]ClusterRoleTemplateBindingController
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

		projectControllers:                    map[string]ProjectController{},
		projectRoleTemplateControllers:        map[string]ProjectRoleTemplateController{},
		podSecurityPolicyTemplateControllers:  map[string]PodSecurityPolicyTemplateController{},
		projectRoleTemplateBindingControllers: map[string]ProjectRoleTemplateBindingController{},
		clusterRoleTemplateControllers:        map[string]ClusterRoleTemplateController{},
		clusterRoleTemplateBindingControllers: map[string]ClusterRoleTemplateBindingController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

type ProjectsGetter interface {
	Projects(namespace string) ProjectInterface
}

func (c *Client) Projects(namespace string) ProjectInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ProjectResource, ProjectGroupVersionKind, projectFactory{})
	return &projectClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ProjectRoleTemplatesGetter interface {
	ProjectRoleTemplates(namespace string) ProjectRoleTemplateInterface
}

func (c *Client) ProjectRoleTemplates(namespace string) ProjectRoleTemplateInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ProjectRoleTemplateResource, ProjectRoleTemplateGroupVersionKind, projectRoleTemplateFactory{})
	return &projectRoleTemplateClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type PodSecurityPolicyTemplatesGetter interface {
	PodSecurityPolicyTemplates(namespace string) PodSecurityPolicyTemplateInterface
}

func (c *Client) PodSecurityPolicyTemplates(namespace string) PodSecurityPolicyTemplateInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &PodSecurityPolicyTemplateResource, PodSecurityPolicyTemplateGroupVersionKind, podSecurityPolicyTemplateFactory{})
	return &podSecurityPolicyTemplateClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ProjectRoleTemplateBindingsGetter interface {
	ProjectRoleTemplateBindings(namespace string) ProjectRoleTemplateBindingInterface
}

func (c *Client) ProjectRoleTemplateBindings(namespace string) ProjectRoleTemplateBindingInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ProjectRoleTemplateBindingResource, ProjectRoleTemplateBindingGroupVersionKind, projectRoleTemplateBindingFactory{})
	return &projectRoleTemplateBindingClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ClusterRoleTemplatesGetter interface {
	ClusterRoleTemplates(namespace string) ClusterRoleTemplateInterface
}

func (c *Client) ClusterRoleTemplates(namespace string) ClusterRoleTemplateInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ClusterRoleTemplateResource, ClusterRoleTemplateGroupVersionKind, clusterRoleTemplateFactory{})
	return &clusterRoleTemplateClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ClusterRoleTemplateBindingsGetter interface {
	ClusterRoleTemplateBindings(namespace string) ClusterRoleTemplateBindingInterface
}

func (c *Client) ClusterRoleTemplateBindings(namespace string) ClusterRoleTemplateBindingInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ClusterRoleTemplateBindingResource, ClusterRoleTemplateBindingGroupVersionKind, clusterRoleTemplateBindingFactory{})
	return &clusterRoleTemplateBindingClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
