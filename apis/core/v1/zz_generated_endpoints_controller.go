package v1

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	EndpointsGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Endpoints",
	}
	EndpointsResource = metav1.APIResource{
		Name:         "endpoints",
		SingularName: "endpoints",
		Namespaced:   true,

		Kind: EndpointsGroupVersionKind.Kind,
	}
)

type EndpointsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.Endpoints
}

type EndpointsHandlerFunc func(key string, obj *v1.Endpoints) error

type EndpointsLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.Endpoints, err error)
	Get(namespace, name string) (*v1.Endpoints, error)
}

type EndpointsController interface {
	Informer() cache.SharedIndexInformer
	Lister() EndpointsLister
	AddHandler(name string, handler EndpointsHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type EndpointsInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*v1.Endpoints) (*v1.Endpoints, error)
	GetNamespace(name, namespace string, opts metav1.GetOptions) (*v1.Endpoints, error)
	Get(name string, opts metav1.GetOptions) (*v1.Endpoints, error)
	Update(*v1.Endpoints) (*v1.Endpoints, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*EndpointsList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() EndpointsController
	AddHandler(name string, sync EndpointsHandlerFunc)
	AddLifecycle(name string, lifecycle EndpointsLifecycle)
}

type endpointsLister struct {
	controller *endpointsController
}

func (l *endpointsLister) List(namespace string, selector labels.Selector) (ret []*v1.Endpoints, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.Endpoints))
	})
	return
}

func (l *endpointsLister) Get(namespace, name string) (*v1.Endpoints, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    EndpointsGroupVersionKind.Group,
			Resource: "endpoints",
		}, name)
	}
	return obj.(*v1.Endpoints), nil
}

type endpointsController struct {
	controller.GenericController
}

func (c *endpointsController) Lister() EndpointsLister {
	return &endpointsLister{
		controller: c,
	}
}

func (c *endpointsController) AddHandler(name string, handler EndpointsHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*v1.Endpoints))
	})
}

type endpointsFactory struct {
}

func (c endpointsFactory) Object() runtime.Object {
	return &v1.Endpoints{}
}

func (c endpointsFactory) List() runtime.Object {
	return &EndpointsList{}
}

func (s *endpointsClient) Controller() EndpointsController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.endpointsControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(EndpointsGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &endpointsController{
		GenericController: genericController,
	}

	s.client.endpointsControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type endpointsClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   EndpointsController
}

func (s *endpointsClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *endpointsClient) Create(o *v1.Endpoints) (*v1.Endpoints, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.Endpoints), err
}

func (s *endpointsClient) Get(name string, opts metav1.GetOptions) (*v1.Endpoints, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.Endpoints), err
}

func (s *endpointsClient) GetNamespace(name, namespace string, opts metav1.GetOptions) (*v1.Endpoints, error) {
	obj, err := s.objectClient.GetNamespace(name, namespace, opts)
	return obj.(*v1.Endpoints), err
}

func (s *endpointsClient) Update(o *v1.Endpoints) (*v1.Endpoints, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.Endpoints), err
}

func (s *endpointsClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *endpointsClient) DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespace(name, namespace, options)
}

func (s *endpointsClient) List(opts metav1.ListOptions) (*EndpointsList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*EndpointsList), err
}

func (s *endpointsClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *endpointsClient) Patch(o *v1.Endpoints, data []byte, subresources ...string) (*v1.Endpoints, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*v1.Endpoints), err
}

func (s *endpointsClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *endpointsClient) AddHandler(name string, sync EndpointsHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *endpointsClient) AddLifecycle(name string, lifecycle EndpointsLifecycle) {
	sync := NewEndpointsLifecycleAdapter(name, s, lifecycle)
	s.AddHandler(name, sync)
}
