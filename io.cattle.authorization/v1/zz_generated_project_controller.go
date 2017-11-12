package v1

import (
	"sync"

	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

var (
	ProjectGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "io.cattle.authorization",
		Kind:    "Project",
	}
	ProjectResource = metav1.APIResource{
		Name:         "projects",
		SingularName: "project",
		Namespaced:   false,
		Kind:         ProjectGroupVersionKind.Kind,
	}
)

type ProjectList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Project
}

type ProjectHandlerFunc func(key string, obj *Project) error

type ProjectController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler ProjectHandlerFunc)
	Enqueue(namespace, name string)
	Start(threadiness int, ctx context.Context) error
}

type ProjectInterface interface {
	Create(*Project) (*Project, error)
	Get(name string, opts metav1.GetOptions) (*Project, error)
	Update(*Project) (*Project, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ProjectList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() (ProjectController, error)
}

type projectController struct {
	controller.GenericController
}

func (c *projectController) AddHandler(handler ProjectHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*Project))
	})
}

type projectFactory struct {
}

func (c projectFactory) Object() runtime.Object {
	return &Project{}
}

func (c projectFactory) List() runtime.Object {
	return &ProjectList{}
}

func NewProjectClient(namespace string, config rest.Config) (ProjectInterface, error) {
	objectClient, err := clientbase.NewObjectClient(namespace, config, &ProjectResource, ProjectGroupVersionKind, projectFactory{})
	return &projectClient{
		objectClient: objectClient,
	}, err
}

func (s *projectClient) Controller() (ProjectController, error) {
	s.Lock()
	defer s.Unlock()

	if s.controller != nil {
		return s.controller, nil
	}

	controller, err := controller.NewGenericController(ProjectGroupVersionKind.Kind+"Controller",
		s.objectClient)
	if err != nil {
		return nil, err
	}

	s.controller = &projectController{
		GenericController: controller,
	}
	return s.controller, nil
}

type projectClient struct {
	sync.Mutex
	objectClient *clientbase.ObjectClient
	controller   ProjectController
}

func (s *projectClient) Create(o *Project) (*Project, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*Project), err
}

func (s *projectClient) Get(name string, opts metav1.GetOptions) (*Project, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*Project), err
}

func (s *projectClient) Update(o *Project) (*Project, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*Project), err
}

func (s *projectClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *projectClient) List(opts metav1.ListOptions) (*ProjectList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ProjectList), err
}

func (s *projectClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *projectClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
