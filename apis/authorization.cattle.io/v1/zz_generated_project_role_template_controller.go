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
	ProjectRoleTemplateGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Kind:    "ProjectRoleTemplate",
	}
	ProjectRoleTemplateResource = metav1.APIResource{
		Name:         "projectroletemplates",
		SingularName: "projectroletemplate",
		Namespaced:   false,
		Kind:         ProjectRoleTemplateGroupVersionKind.Kind,
	}
)

type ProjectRoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRoleTemplate
}

type ProjectRoleTemplateHandlerFunc func(key string, obj *ProjectRoleTemplate) error

type ProjectRoleTemplateController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler ProjectRoleTemplateHandlerFunc)
	Enqueue(namespace, name string)
	Start(threadiness int, ctx context.Context) error
}

type ProjectRoleTemplateInterface interface {
	Create(*ProjectRoleTemplate) (*ProjectRoleTemplate, error)
	Get(name string, opts metav1.GetOptions) (*ProjectRoleTemplate, error)
	Update(*ProjectRoleTemplate) (*ProjectRoleTemplate, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ProjectRoleTemplateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ProjectRoleTemplateController
}

type projectRoleTemplateController struct {
	controller.GenericController
}

func (c *projectRoleTemplateController) AddHandler(handler ProjectRoleTemplateHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ProjectRoleTemplate))
	})
}

type projectRoleTemplateFactory struct {
}

func (c projectRoleTemplateFactory) Object() runtime.Object {
	return &ProjectRoleTemplate{}
}

func (c projectRoleTemplateFactory) List() runtime.Object {
	return &ProjectRoleTemplateList{}
}

func (s *projectRoleTemplateClient) Controller() ProjectRoleTemplateController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.projectRoleTemplateControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ProjectRoleTemplateGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &projectRoleTemplateController{
		GenericController: genericController,
	}

	s.client.projectRoleTemplateControllers[s.ns] = c

	return c
}

type projectRoleTemplateClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   ProjectRoleTemplateController
}

func (s *projectRoleTemplateClient) Create(o *ProjectRoleTemplate) (*ProjectRoleTemplate, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ProjectRoleTemplate), err
}

func (s *projectRoleTemplateClient) Get(name string, opts metav1.GetOptions) (*ProjectRoleTemplate, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ProjectRoleTemplate), err
}

func (s *projectRoleTemplateClient) Update(o *ProjectRoleTemplate) (*ProjectRoleTemplate, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ProjectRoleTemplate), err
}

func (s *projectRoleTemplateClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *projectRoleTemplateClient) List(opts metav1.ListOptions) (*ProjectRoleTemplateList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ProjectRoleTemplateList), err
}

func (s *projectRoleTemplateClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *projectRoleTemplateClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
