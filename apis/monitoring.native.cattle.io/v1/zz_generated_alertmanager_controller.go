package v1

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
	AlertmanagerGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Alertmanager",
	}
	AlertmanagerResource = metav1.APIResource{
		Name:         "alertmanagers",
		SingularName: "alertmanager",
		Namespaced:   true,

		Kind: AlertmanagerGroupVersionKind.Kind,
	}
)

type AlertmanagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alertmanager
}

type AlertmanagerHandlerFunc func(key string, obj *Alertmanager) (runtime.Object, error)

type AlertmanagerLister interface {
	List(namespace string, selector labels.Selector) (ret []*Alertmanager, err error)
	Get(namespace, name string) (*Alertmanager, error)
}

type AlertmanagerController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() AlertmanagerLister
	AddHandler(ctx context.Context, name string, handler AlertmanagerHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler AlertmanagerHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type AlertmanagerInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*Alertmanager) (*Alertmanager, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*Alertmanager, error)
	Get(name string, opts metav1.GetOptions) (*Alertmanager, error)
	Update(*Alertmanager) (*Alertmanager, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*AlertmanagerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() AlertmanagerController
	AddHandler(ctx context.Context, name string, sync AlertmanagerHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle AlertmanagerLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync AlertmanagerHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle AlertmanagerLifecycle)
}

type alertmanagerLister struct {
	controller *alertmanagerController
}

func (l *alertmanagerLister) List(namespace string, selector labels.Selector) (ret []*Alertmanager, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*Alertmanager))
	})
	return
}

func (l *alertmanagerLister) Get(namespace, name string) (*Alertmanager, error) {
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
			Group:    AlertmanagerGroupVersionKind.Group,
			Resource: "alertmanager",
		}, key)
	}
	return obj.(*Alertmanager), nil
}

type alertmanagerController struct {
	controller.GenericController
}

func (c *alertmanagerController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *alertmanagerController) Lister() AlertmanagerLister {
	return &alertmanagerLister{
		controller: c,
	}
}

func (c *alertmanagerController) AddHandler(ctx context.Context, name string, handler AlertmanagerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*Alertmanager); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *alertmanagerController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler AlertmanagerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*Alertmanager); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type alertmanagerFactory struct {
}

func (c alertmanagerFactory) Object() runtime.Object {
	return &Alertmanager{}
}

func (c alertmanagerFactory) List() runtime.Object {
	return &AlertmanagerList{}
}

func (s *alertmanagerClient) Controller() AlertmanagerController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.alertmanagerControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(AlertmanagerGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &alertmanagerController{
		GenericController: genericController,
	}

	s.client.alertmanagerControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type alertmanagerClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   AlertmanagerController
}

func (s *alertmanagerClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *alertmanagerClient) Create(o *Alertmanager) (*Alertmanager, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*Alertmanager), err
}

func (s *alertmanagerClient) Get(name string, opts metav1.GetOptions) (*Alertmanager, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*Alertmanager), err
}

func (s *alertmanagerClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*Alertmanager, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*Alertmanager), err
}

func (s *alertmanagerClient) Update(o *Alertmanager) (*Alertmanager, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*Alertmanager), err
}

func (s *alertmanagerClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *alertmanagerClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *alertmanagerClient) List(opts metav1.ListOptions) (*AlertmanagerList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*AlertmanagerList), err
}

func (s *alertmanagerClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *alertmanagerClient) Patch(o *Alertmanager, data []byte, subresources ...string) (*Alertmanager, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*Alertmanager), err
}

func (s *alertmanagerClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *alertmanagerClient) AddHandler(ctx context.Context, name string, sync AlertmanagerHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *alertmanagerClient) AddLifecycle(ctx context.Context, name string, lifecycle AlertmanagerLifecycle) {
	sync := NewAlertmanagerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *alertmanagerClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync AlertmanagerHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *alertmanagerClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle AlertmanagerLifecycle) {
	sync := NewAlertmanagerLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}
