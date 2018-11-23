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
	PrometheusGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Prometheus",
	}
	PrometheusResource = metav1.APIResource{
		Name:         "prometheuses",
		SingularName: "prometheus",
		Namespaced:   true,

		Kind: PrometheusGroupVersionKind.Kind,
	}
)

type PrometheusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.Prometheus
}

type PrometheusHandlerFunc func(key string, obj *v1.Prometheus) (runtime.Object, error)

type PrometheusLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.Prometheus, err error)
	Get(namespace, name string) (*v1.Prometheus, error)
}

type PrometheusController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() PrometheusLister
	AddHandler(ctx context.Context, name string, handler PrometheusHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler PrometheusHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type PrometheusInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.Prometheus) (*v1.Prometheus, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.Prometheus, error)
	Get(name string, opts metav1.GetOptions) (*v1.Prometheus, error)
	Update(*v1.Prometheus) (*v1.Prometheus, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*PrometheusList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() PrometheusController
	AddHandler(ctx context.Context, name string, sync PrometheusHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle PrometheusLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PrometheusHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PrometheusLifecycle)
}

type prometheusLister struct {
	controller *prometheusController
}

func (l *prometheusLister) List(namespace string, selector labels.Selector) (ret []*v1.Prometheus, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.Prometheus))
	})
	return
}

func (l *prometheusLister) Get(namespace, name string) (*v1.Prometheus, error) {
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
			Group:    PrometheusGroupVersionKind.Group,
			Resource: "prometheus",
		}, key)
	}
	return obj.(*v1.Prometheus), nil
}

type prometheusController struct {
	controller.GenericController
}

func (c *prometheusController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *prometheusController) Lister() PrometheusLister {
	return &prometheusLister{
		controller: c,
	}
}

func (c *prometheusController) AddHandler(ctx context.Context, name string, handler PrometheusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Prometheus); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *prometheusController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler PrometheusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Prometheus); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type prometheusFactory struct {
}

func (c prometheusFactory) Object() runtime.Object {
	return &v1.Prometheus{}
}

func (c prometheusFactory) List() runtime.Object {
	return &PrometheusList{}
}

func (s *prometheusClient) Controller() PrometheusController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.prometheusControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(PrometheusGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &prometheusController{
		GenericController: genericController,
	}

	s.client.prometheusControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type prometheusClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   PrometheusController
}

func (s *prometheusClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *prometheusClient) Create(o *v1.Prometheus) (*v1.Prometheus, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.Prometheus), err
}

func (s *prometheusClient) Get(name string, opts metav1.GetOptions) (*v1.Prometheus, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.Prometheus), err
}

func (s *prometheusClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.Prometheus, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.Prometheus), err
}

func (s *prometheusClient) Update(o *v1.Prometheus) (*v1.Prometheus, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.Prometheus), err
}

func (s *prometheusClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *prometheusClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *prometheusClient) List(opts metav1.ListOptions) (*PrometheusList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*PrometheusList), err
}

func (s *prometheusClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *prometheusClient) Patch(o *v1.Prometheus, data []byte, subresources ...string) (*v1.Prometheus, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*v1.Prometheus), err
}

func (s *prometheusClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *prometheusClient) AddHandler(ctx context.Context, name string, sync PrometheusHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *prometheusClient) AddLifecycle(ctx context.Context, name string, lifecycle PrometheusLifecycle) {
	sync := NewPrometheusLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *prometheusClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PrometheusHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *prometheusClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PrometheusLifecycle) {
	sync := NewPrometheusLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}
