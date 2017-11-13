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
	RoleTemplatesGetter
	PodSecurityPolicyTemplatesGetter
	ProjectRoleBindingsGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface

	projectControllers                   map[string]ProjectController
	roleTemplateControllers              map[string]RoleTemplateController
	podSecurityPolicyTemplateControllers map[string]PodSecurityPolicyTemplateController
	projectRoleBindingControllers        map[string]ProjectRoleBindingController
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

		projectControllers:                   map[string]ProjectController{},
		roleTemplateControllers:              map[string]RoleTemplateController{},
		podSecurityPolicyTemplateControllers: map[string]PodSecurityPolicyTemplateController{},
		projectRoleBindingControllers:        map[string]ProjectRoleBindingController{},
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

type RoleTemplatesGetter interface {
	RoleTemplates(namespace string) RoleTemplateInterface
}

func (c *Client) RoleTemplates(namespace string) RoleTemplateInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &RoleTemplateResource, RoleTemplateGroupVersionKind, roleTemplateFactory{})
	return &roleTemplateClient{
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

type ProjectRoleBindingsGetter interface {
	ProjectRoleBindings(namespace string) ProjectRoleBindingInterface
}

func (c *Client) ProjectRoleBindings(namespace string) ProjectRoleBindingInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &ProjectRoleBindingResource, ProjectRoleBindingGroupVersionKind, projectRoleBindingFactory{})
	return &projectRoleBindingClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
