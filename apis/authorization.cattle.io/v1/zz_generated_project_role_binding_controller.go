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
	ProjectRoleBindingGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Kind:    "ProjectRoleBinding",
	}
	ProjectRoleBindingResource = metav1.APIResource{
		Name:         "projectrolebindings",
		SingularName: "projectrolebinding",
		Namespaced:   false,
		Kind:         ProjectRoleBindingGroupVersionKind.Kind,
	}
)

type ProjectRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRoleBinding
}

type ProjectRoleBindingHandlerFunc func(key string, obj *ProjectRoleBinding) error

type ProjectRoleBindingController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler ProjectRoleBindingHandlerFunc)
	Enqueue(namespace, name string)
	Start(threadiness int, ctx context.Context) error
}

type ProjectRoleBindingInterface interface {
	Create(*ProjectRoleBinding) (*ProjectRoleBinding, error)
	Get(name string, opts metav1.GetOptions) (*ProjectRoleBinding, error)
	Update(*ProjectRoleBinding) (*ProjectRoleBinding, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ProjectRoleBindingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ProjectRoleBindingController
}

type projectRoleBindingController struct {
	controller.GenericController
}

func (c *projectRoleBindingController) AddHandler(handler ProjectRoleBindingHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ProjectRoleBinding))
	})
}

type projectRoleBindingFactory struct {
}

func (c projectRoleBindingFactory) Object() runtime.Object {
	return &ProjectRoleBinding{}
}

func (c projectRoleBindingFactory) List() runtime.Object {
	return &ProjectRoleBindingList{}
}

func (s *projectRoleBindingClient) Controller() ProjectRoleBindingController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.projectRoleBindingControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ProjectRoleBindingGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &projectRoleBindingController{
		GenericController: genericController,
	}

	s.client.projectRoleBindingControllers[s.ns] = c

	return c
}

type projectRoleBindingClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   ProjectRoleBindingController
}

func (s *projectRoleBindingClient) Create(o *ProjectRoleBinding) (*ProjectRoleBinding, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ProjectRoleBinding), err
}

func (s *projectRoleBindingClient) Get(name string, opts metav1.GetOptions) (*ProjectRoleBinding, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ProjectRoleBinding), err
}

func (s *projectRoleBindingClient) Update(o *ProjectRoleBinding) (*ProjectRoleBinding, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ProjectRoleBinding), err
}

func (s *projectRoleBindingClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *projectRoleBindingClient) List(opts metav1.ListOptions) (*ProjectRoleBindingList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ProjectRoleBindingList), err
}

func (s *projectRoleBindingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *projectRoleBindingClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
