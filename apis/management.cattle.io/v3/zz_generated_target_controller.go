package v3

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
	TargetGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Target",
	}
	TargetResource = metav1.APIResource{
		Name:         "targets",
		SingularName: "target",
		Namespaced:   false,
		Kind:         TargetGroupVersionKind.Kind,
	}
)

type TargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Target
}

type TargetHandlerFunc func(key string, obj *Target) error

type TargetLister interface {
	List(namespace string, selector labels.Selector) (ret []*Target, err error)
	Get(namespace, name string) (*Target, error)
}

type TargetController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() TargetLister
	AddHandler(name string, handler TargetHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler TargetHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type TargetInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*Target) (*Target, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*Target, error)
	Get(name string, opts metav1.GetOptions) (*Target, error)
	Update(*Target) (*Target, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*TargetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() TargetController
	AddHandler(name string, sync TargetHandlerFunc)
	AddLifecycle(name string, lifecycle TargetLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync TargetHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle TargetLifecycle)
}

type targetLister struct {
	controller *targetController
}

func (l *targetLister) List(namespace string, selector labels.Selector) (ret []*Target, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*Target))
	})
	return
}

func (l *targetLister) Get(namespace, name string) (*Target, error) {
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
			Group:    TargetGroupVersionKind.Group,
			Resource: "target",
		}, key)
	}
	return obj.(*Target), nil
}

type targetController struct {
	controller.GenericController
}

func (c *targetController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *targetController) Lister() TargetLister {
	return &targetLister{
		controller: c,
	}
}

func (c *targetController) AddHandler(name string, handler TargetHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*Target))
	})
}

func (c *targetController) AddClusterScopedHandler(name, cluster string, handler TargetHandlerFunc) {
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

		return handler(key, obj.(*Target))
	})
}

type targetFactory struct {
}

func (c targetFactory) Object() runtime.Object {
	return &Target{}
}

func (c targetFactory) List() runtime.Object {
	return &TargetList{}
}

func (s *targetClient) Controller() TargetController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.targetControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(TargetGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &targetController{
		GenericController: genericController,
	}

	s.client.targetControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type targetClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   TargetController
}

func (s *targetClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *targetClient) Create(o *Target) (*Target, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*Target), err
}

func (s *targetClient) Get(name string, opts metav1.GetOptions) (*Target, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*Target), err
}

func (s *targetClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*Target, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*Target), err
}

func (s *targetClient) Update(o *Target) (*Target, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*Target), err
}

func (s *targetClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *targetClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *targetClient) List(opts metav1.ListOptions) (*TargetList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*TargetList), err
}

func (s *targetClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *targetClient) Patch(o *Target, data []byte, subresources ...string) (*Target, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*Target), err
}

func (s *targetClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *targetClient) AddHandler(name string, sync TargetHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *targetClient) AddLifecycle(name string, lifecycle TargetLifecycle) {
	sync := NewTargetLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *targetClient) AddClusterScopedHandler(name, clusterName string, sync TargetHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *targetClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle TargetLifecycle) {
	sync := NewTargetLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
