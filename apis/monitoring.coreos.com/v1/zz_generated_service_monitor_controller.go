package v1

import (
	"context"

	"github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
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
	ServiceMonitorGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ServiceMonitor",
	}
	ServiceMonitorResource = metav1.APIResource{
		Name:         "servicemonitors",
		SingularName: "servicemonitor",
		Namespaced:   true,

		Kind: ServiceMonitorGroupVersionKind.Kind,
	}
)

type ServiceMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.ServiceMonitor
}

type ServiceMonitorHandlerFunc func(key string, obj *v1.ServiceMonitor) (runtime.Object, error)

type ServiceMonitorLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.ServiceMonitor, err error)
	Get(namespace, name string) (*v1.ServiceMonitor, error)
}

type ServiceMonitorController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ServiceMonitorLister
	AddHandler(ctx context.Context, name string, handler ServiceMonitorHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ServiceMonitorHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ServiceMonitorInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.ServiceMonitor) (*v1.ServiceMonitor, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ServiceMonitor, error)
	Get(name string, opts metav1.GetOptions) (*v1.ServiceMonitor, error)
	Update(*v1.ServiceMonitor) (*v1.ServiceMonitor, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ServiceMonitorList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ServiceMonitorController
	AddHandler(ctx context.Context, name string, sync ServiceMonitorHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ServiceMonitorLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ServiceMonitorHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ServiceMonitorLifecycle)
}

type serviceMonitorLister struct {
	controller *serviceMonitorController
}

func (l *serviceMonitorLister) List(namespace string, selector labels.Selector) (ret []*v1.ServiceMonitor, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.ServiceMonitor))
	})
	return
}

func (l *serviceMonitorLister) Get(namespace, name string) (*v1.ServiceMonitor, error) {
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
			Group:    ServiceMonitorGroupVersionKind.Group,
			Resource: "serviceMonitor",
		}, key)
	}
	return obj.(*v1.ServiceMonitor), nil
}

type serviceMonitorController struct {
	controller.GenericController
}

func (c *serviceMonitorController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *serviceMonitorController) Lister() ServiceMonitorLister {
	return &serviceMonitorLister{
		controller: c,
	}
}

func (c *serviceMonitorController) AddHandler(ctx context.Context, name string, handler ServiceMonitorHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ServiceMonitor); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *serviceMonitorController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ServiceMonitorHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ServiceMonitor); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type serviceMonitorFactory struct {
}

func (c serviceMonitorFactory) Object() runtime.Object {
	return &v1.ServiceMonitor{}
}

func (c serviceMonitorFactory) List() runtime.Object {
	return &ServiceMonitorList{}
}

func (s *serviceMonitorClient) Controller() ServiceMonitorController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.serviceMonitorControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ServiceMonitorGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &serviceMonitorController{
		GenericController: genericController,
	}

	s.client.serviceMonitorControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type serviceMonitorClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ServiceMonitorController
}

func (s *serviceMonitorClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *serviceMonitorClient) Create(o *v1.ServiceMonitor) (*v1.ServiceMonitor, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.ServiceMonitor), err
}

func (s *serviceMonitorClient) Get(name string, opts metav1.GetOptions) (*v1.ServiceMonitor, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.ServiceMonitor), err
}

func (s *serviceMonitorClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ServiceMonitor, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.ServiceMonitor), err
}

func (s *serviceMonitorClient) Update(o *v1.ServiceMonitor) (*v1.ServiceMonitor, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.ServiceMonitor), err
}

func (s *serviceMonitorClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *serviceMonitorClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *serviceMonitorClient) List(opts metav1.ListOptions) (*ServiceMonitorList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ServiceMonitorList), err
}

func (s *serviceMonitorClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *serviceMonitorClient) Patch(o *v1.ServiceMonitor, data []byte, subresources ...string) (*v1.ServiceMonitor, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*v1.ServiceMonitor), err
}

func (s *serviceMonitorClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *serviceMonitorClient) AddHandler(ctx context.Context, name string, sync ServiceMonitorHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *serviceMonitorClient) AddLifecycle(ctx context.Context, name string, lifecycle ServiceMonitorLifecycle) {
	sync := NewServiceMonitorLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *serviceMonitorClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ServiceMonitorHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *serviceMonitorClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ServiceMonitorLifecycle) {
	sync := NewServiceMonitorLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}
