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
	ClusterRoleTemplateBindingGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Kind:    "ClusterRoleTemplateBinding",
	}
	ClusterRoleTemplateBindingResource = metav1.APIResource{
		Name:         "clusterroletemplatebindings",
		SingularName: "clusterroletemplatebinding",
		Namespaced:   false,
		Kind:         ClusterRoleTemplateBindingGroupVersionKind.Kind,
	}
)

type ClusterRoleTemplateBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRoleTemplateBinding
}

type ClusterRoleTemplateBindingHandlerFunc func(key string, obj *ClusterRoleTemplateBinding) error

type ClusterRoleTemplateBindingController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler ClusterRoleTemplateBindingHandlerFunc)
	Enqueue(namespace, name string)
	Start(threadiness int, ctx context.Context) error
}

type ClusterRoleTemplateBindingInterface interface {
	Create(*ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error)
	Get(name string, opts metav1.GetOptions) (*ClusterRoleTemplateBinding, error)
	Update(*ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ClusterRoleTemplateBindingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ClusterRoleTemplateBindingController
}

type clusterRoleTemplateBindingController struct {
	controller.GenericController
}

func (c *clusterRoleTemplateBindingController) AddHandler(handler ClusterRoleTemplateBindingHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ClusterRoleTemplateBinding))
	})
}

type clusterRoleTemplateBindingFactory struct {
}

func (c clusterRoleTemplateBindingFactory) Object() runtime.Object {
	return &ClusterRoleTemplateBinding{}
}

func (c clusterRoleTemplateBindingFactory) List() runtime.Object {
	return &ClusterRoleTemplateBindingList{}
}

func (s *clusterRoleTemplateBindingClient) Controller() ClusterRoleTemplateBindingController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.clusterRoleTemplateBindingControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ClusterRoleTemplateBindingGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &clusterRoleTemplateBindingController{
		GenericController: genericController,
	}

	s.client.clusterRoleTemplateBindingControllers[s.ns] = c

	return c
}

type clusterRoleTemplateBindingClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   ClusterRoleTemplateBindingController
}

func (s *clusterRoleTemplateBindingClient) Create(o *ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ClusterRoleTemplateBinding), err
}

func (s *clusterRoleTemplateBindingClient) Get(name string, opts metav1.GetOptions) (*ClusterRoleTemplateBinding, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ClusterRoleTemplateBinding), err
}

func (s *clusterRoleTemplateBindingClient) Update(o *ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ClusterRoleTemplateBinding), err
}

func (s *clusterRoleTemplateBindingClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *clusterRoleTemplateBindingClient) List(opts metav1.ListOptions) (*ClusterRoleTemplateBindingList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ClusterRoleTemplateBindingList), err
}

func (s *clusterRoleTemplateBindingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *clusterRoleTemplateBindingClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
