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
	RoleTemplateGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Kind:    "RoleTemplate",
	}
	RoleTemplateResource = metav1.APIResource{
		Name:         "roletemplates",
		SingularName: "roletemplate",
		Namespaced:   false,
		Kind:         RoleTemplateGroupVersionKind.Kind,
	}
)

type RoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleTemplate
}

type RoleTemplateHandlerFunc func(key string, obj *RoleTemplate) error

type RoleTemplateController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler RoleTemplateHandlerFunc)
	Enqueue(namespace, name string)
	Start(threadiness int, ctx context.Context) error
}

type RoleTemplateInterface interface {
	Create(*RoleTemplate) (*RoleTemplate, error)
	Get(name string, opts metav1.GetOptions) (*RoleTemplate, error)
	Update(*RoleTemplate) (*RoleTemplate, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*RoleTemplateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() RoleTemplateController
}

type roleTemplateController struct {
	controller.GenericController
}

func (c *roleTemplateController) AddHandler(handler RoleTemplateHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*RoleTemplate))
	})
}

type roleTemplateFactory struct {
}

func (c roleTemplateFactory) Object() runtime.Object {
	return &RoleTemplate{}
}

func (c roleTemplateFactory) List() runtime.Object {
	return &RoleTemplateList{}
}

func (s *roleTemplateClient) Controller() RoleTemplateController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.roleTemplateControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(RoleTemplateGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &roleTemplateController{
		GenericController: genericController,
	}

	s.client.roleTemplateControllers[s.ns] = c

	return c
}

type roleTemplateClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   RoleTemplateController
}

func (s *roleTemplateClient) Create(o *RoleTemplate) (*RoleTemplate, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*RoleTemplate), err
}

func (s *roleTemplateClient) Get(name string, opts metav1.GetOptions) (*RoleTemplate, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*RoleTemplate), err
}

func (s *roleTemplateClient) Update(o *RoleTemplate) (*RoleTemplate, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*RoleTemplate), err
}

func (s *roleTemplateClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *roleTemplateClient) List(opts metav1.ListOptions) (*RoleTemplateList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*RoleTemplateList), err
}

func (s *roleTemplateClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *roleTemplateClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
