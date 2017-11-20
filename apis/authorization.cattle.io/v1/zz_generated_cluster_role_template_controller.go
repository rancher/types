package v1

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ClusterRoleTemplateGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Kind:    "ClusterRoleTemplate",
	}
	ClusterRoleTemplateResource = metav1.APIResource{
		Name:         "clusterroletemplates",
		SingularName: "clusterroletemplate",
		Namespaced:   false,
		Kind:         ClusterRoleTemplateGroupVersionKind.Kind,
	}
)

type ClusterRoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRoleTemplate
}

type ClusterRoleTemplateHandlerFunc func(key string, obj *ClusterRoleTemplate) error

type ClusterRoleTemplateController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler ClusterRoleTemplateHandlerFunc)
	Enqueue(namespace, name string)
	Start(threadiness int, ctx context.Context) error
}

type ClusterRoleTemplateInterface interface {
	Create(*ClusterRoleTemplate) (*ClusterRoleTemplate, error)
	Get(name string, opts metav1.GetOptions) (*ClusterRoleTemplate, error)
	Update(*ClusterRoleTemplate) (*ClusterRoleTemplate, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ClusterRoleTemplateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ClusterRoleTemplateController
}

type clusterRoleTemplateController struct {
	controller.GenericController
}

func (c *clusterRoleTemplateController) AddHandler(handler ClusterRoleTemplateHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ClusterRoleTemplate))
	})
}

type clusterRoleTemplateFactory struct {
}

func (c clusterRoleTemplateFactory) Object() runtime.Object {
	return &ClusterRoleTemplate{}
}

func (c clusterRoleTemplateFactory) List() runtime.Object {
	return &ClusterRoleTemplateList{}
}

func (s *clusterRoleTemplateClient) Controller() ClusterRoleTemplateController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.clusterRoleTemplateControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ClusterRoleTemplateGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &clusterRoleTemplateController{
		GenericController: genericController,
	}

	s.client.clusterRoleTemplateControllers[s.ns] = c

	return c
}

type clusterRoleTemplateClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   ClusterRoleTemplateController
}

func (s *clusterRoleTemplateClient) Create(o *ClusterRoleTemplate) (*ClusterRoleTemplate, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ClusterRoleTemplate), err
}

func (s *clusterRoleTemplateClient) Get(name string, opts metav1.GetOptions) (*ClusterRoleTemplate, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ClusterRoleTemplate), err
}

func (s *clusterRoleTemplateClient) Update(o *ClusterRoleTemplate) (*ClusterRoleTemplate, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ClusterRoleTemplate), err
}

func (s *clusterRoleTemplateClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *clusterRoleTemplateClient) List(opts metav1.ListOptions) (*ClusterRoleTemplateList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ClusterRoleTemplateList), err
}

func (s *clusterRoleTemplateClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *clusterRoleTemplateClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
