package v3public

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ActiveDirectoryProviderGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ActiveDirectoryProvider",
	}
	ActiveDirectoryProviderResource = metav1.APIResource{
		Name:         "activedirectoryproviders",
		SingularName: "activedirectoryprovider",
		Namespaced:   false,
		Kind:         ActiveDirectoryProviderGroupVersionKind.Kind,
	}
)

type ActiveDirectoryProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryProvider
}

type ActiveDirectoryProviderHandlerFunc func(key string, obj *ActiveDirectoryProvider) error

type ActiveDirectoryProviderLister interface {
	List(namespace string, selector labels.Selector) (ret []*ActiveDirectoryProvider, err error)
	Get(namespace, name string) (*ActiveDirectoryProvider, error)
}

type ActiveDirectoryProviderController interface {
	Informer() cache.SharedIndexInformer
	Lister() ActiveDirectoryProviderLister
	AddHandler(name string, handler ActiveDirectoryProviderHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ActiveDirectoryProviderHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ActiveDirectoryProviderInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error)
	Get(name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error)
	Update(*ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ActiveDirectoryProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ActiveDirectoryProviderController
	AddHandler(name string, sync ActiveDirectoryProviderHandlerFunc)
	AddLifecycle(name string, lifecycle ActiveDirectoryProviderLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ActiveDirectoryProviderHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ActiveDirectoryProviderLifecycle)
}

type activeDirectoryProviderLister struct {
	controller *activeDirectoryProviderController
}

func (l *activeDirectoryProviderLister) List(namespace string, selector labels.Selector) (ret []*ActiveDirectoryProvider, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ActiveDirectoryProvider))
	})
	return
}

func (l *activeDirectoryProviderLister) Get(namespace, name string) (*ActiveDirectoryProvider, error) {
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
			Group:    ActiveDirectoryProviderGroupVersionKind.Group,
			Resource: "activeDirectoryProvider",
		}, key)
	}
	return obj.(*ActiveDirectoryProvider), nil
}

type activeDirectoryProviderController struct {
	controller.GenericController
}

func (c *activeDirectoryProviderController) Lister() ActiveDirectoryProviderLister {
	return &activeDirectoryProviderLister{
		controller: c,
	}
}

func (c *activeDirectoryProviderController) AddHandler(name string, handler ActiveDirectoryProviderHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ActiveDirectoryProvider))
	})
}

func (c *activeDirectoryProviderController) AddClusterScopedHandler(name, cluster string, handler ActiveDirectoryProviderHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*ActiveDirectoryProvider))
	})
}

type activeDirectoryProviderFactory struct {
}

func (c activeDirectoryProviderFactory) Object() runtime.Object {
	return &ActiveDirectoryProvider{}
}

func (c activeDirectoryProviderFactory) List() runtime.Object {
	return &ActiveDirectoryProviderList{}
}

func (s *activeDirectoryProviderClient) Controller() ActiveDirectoryProviderController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.activeDirectoryProviderControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ActiveDirectoryProviderGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &activeDirectoryProviderController{
		GenericController: genericController,
	}

	s.client.activeDirectoryProviderControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type activeDirectoryProviderClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ActiveDirectoryProviderController
}

func (s *activeDirectoryProviderClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *activeDirectoryProviderClient) Create(o *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) Get(name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) Update(o *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *activeDirectoryProviderClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *activeDirectoryProviderClient) List(opts metav1.ListOptions) (*ActiveDirectoryProviderList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ActiveDirectoryProviderList), err
}

func (s *activeDirectoryProviderClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *activeDirectoryProviderClient) Patch(o *ActiveDirectoryProvider, data []byte, subresources ...string) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *activeDirectoryProviderClient) AddHandler(name string, sync ActiveDirectoryProviderHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *activeDirectoryProviderClient) AddLifecycle(name string, lifecycle ActiveDirectoryProviderLifecycle) {
	sync := NewActiveDirectoryProviderLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *activeDirectoryProviderClient) AddClusterScopedHandler(name, clusterName string, sync ActiveDirectoryProviderHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *activeDirectoryProviderClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ActiveDirectoryProviderLifecycle) {
	sync := NewActiveDirectoryProviderLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
