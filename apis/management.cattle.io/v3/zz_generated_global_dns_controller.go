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
	GlobalDNSGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "GlobalDNS",
	}
	GlobalDNSResource = metav1.APIResource{
		Name:         "globaldnss",
		SingularName: "globaldns",
		Namespaced:   false,
		Kind:         GlobalDNSGroupVersionKind.Kind,
	}
)

type GlobalDNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalDNS
}

type GlobalDNSHandlerFunc func(key string, obj *GlobalDNS) error

type GlobalDNSLister interface {
	List(namespace string, selector labels.Selector) (ret []*GlobalDNS, err error)
	Get(namespace, name string) (*GlobalDNS, error)
}

type GlobalDNSController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() GlobalDNSLister
	AddHandler(name string, handler GlobalDNSHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler GlobalDNSHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type GlobalDNSInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*GlobalDNS) (*GlobalDNS, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalDNS, error)
	Get(name string, opts metav1.GetOptions) (*GlobalDNS, error)
	Update(*GlobalDNS) (*GlobalDNS, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*GlobalDNSList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() GlobalDNSController
	AddHandler(name string, sync GlobalDNSHandlerFunc)
	AddLifecycle(name string, lifecycle GlobalDNSLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync GlobalDNSHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle GlobalDNSLifecycle)
}

type globalDNSLister struct {
	controller *globalDNSController
}

func (l *globalDNSLister) List(namespace string, selector labels.Selector) (ret []*GlobalDNS, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*GlobalDNS))
	})
	return
}

func (l *globalDNSLister) Get(namespace, name string) (*GlobalDNS, error) {
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
			Group:    GlobalDNSGroupVersionKind.Group,
			Resource: "globalDNS",
		}, key)
	}
	return obj.(*GlobalDNS), nil
}

type globalDNSController struct {
	controller.GenericController
}

func (c *globalDNSController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *globalDNSController) Lister() GlobalDNSLister {
	return &globalDNSLister{
		controller: c,
	}
}

func (c *globalDNSController) AddHandler(name string, handler GlobalDNSHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*GlobalDNS))
	})
}

func (c *globalDNSController) AddClusterScopedHandler(name, cluster string, handler GlobalDNSHandlerFunc) {
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

		return handler(key, obj.(*GlobalDNS))
	})
}

type globalDNSFactory struct {
}

func (c globalDNSFactory) Object() runtime.Object {
	return &GlobalDNS{}
}

func (c globalDNSFactory) List() runtime.Object {
	return &GlobalDNSList{}
}

func (s *globalDNSClient) Controller() GlobalDNSController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.globalDNSControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(GlobalDNSGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &globalDNSController{
		GenericController: genericController,
	}

	s.client.globalDNSControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type globalDNSClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   GlobalDNSController
}

func (s *globalDNSClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *globalDNSClient) Create(o *GlobalDNS) (*GlobalDNS, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*GlobalDNS), err
}

func (s *globalDNSClient) Get(name string, opts metav1.GetOptions) (*GlobalDNS, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*GlobalDNS), err
}

func (s *globalDNSClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalDNS, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*GlobalDNS), err
}

func (s *globalDNSClient) Update(o *GlobalDNS) (*GlobalDNS, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*GlobalDNS), err
}

func (s *globalDNSClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *globalDNSClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *globalDNSClient) List(opts metav1.ListOptions) (*GlobalDNSList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*GlobalDNSList), err
}

func (s *globalDNSClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *globalDNSClient) Patch(o *GlobalDNS, data []byte, subresources ...string) (*GlobalDNS, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*GlobalDNS), err
}

func (s *globalDNSClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *globalDNSClient) AddHandler(name string, sync GlobalDNSHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *globalDNSClient) AddLifecycle(name string, lifecycle GlobalDNSLifecycle) {
	sync := NewGlobalDNSLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *globalDNSClient) AddClusterScopedHandler(name, clusterName string, sync GlobalDNSHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *globalDNSClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle GlobalDNSLifecycle) {
	sync := NewGlobalDNSLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
