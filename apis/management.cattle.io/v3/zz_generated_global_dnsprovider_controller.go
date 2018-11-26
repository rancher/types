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
	GlobalDNSProviderGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "GlobalDNSProvider",
	}
	GlobalDNSProviderResource = metav1.APIResource{
		Name:         "globaldnsproviders",
		SingularName: "globaldnsprovider",
		Namespaced:   false,
		Kind:         GlobalDNSProviderGroupVersionKind.Kind,
	}
)

type GlobalDNSProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalDNSProvider
}

type GlobalDNSProviderHandlerFunc func(key string, obj *GlobalDNSProvider) (runtime.Object, error)

type GlobalDNSProviderChangeHandlerFunc func(obj *GlobalDNSProvider) (runtime.Object, error)

type GlobalDNSProviderLister interface {
	List(namespace string, selector labels.Selector) (ret []*GlobalDNSProvider, err error)
	Get(namespace, name string) (*GlobalDNSProvider, error)
}

type GlobalDNSProviderController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() GlobalDNSProviderLister
	AddHandler(ctx context.Context, name string, handler GlobalDNSProviderHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler GlobalDNSProviderHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type GlobalDNSProviderInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*GlobalDNSProvider) (*GlobalDNSProvider, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalDNSProvider, error)
	Get(name string, opts metav1.GetOptions) (*GlobalDNSProvider, error)
	Update(*GlobalDNSProvider) (*GlobalDNSProvider, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*GlobalDNSProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() GlobalDNSProviderController
	AddHandler(ctx context.Context, name string, sync GlobalDNSProviderHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle GlobalDNSProviderLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync GlobalDNSProviderHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle GlobalDNSProviderLifecycle)
}

type globalDNSProviderLister struct {
	controller *globalDNSProviderController
}

func (l *globalDNSProviderLister) List(namespace string, selector labels.Selector) (ret []*GlobalDNSProvider, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*GlobalDNSProvider))
	})
	return
}

func (l *globalDNSProviderLister) Get(namespace, name string) (*GlobalDNSProvider, error) {
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
			Group:    GlobalDNSProviderGroupVersionKind.Group,
			Resource: "globalDNSProvider",
		}, key)
	}
	return obj.(*GlobalDNSProvider), nil
}

type globalDNSProviderController struct {
	controller.GenericController
}

func (c *globalDNSProviderController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *globalDNSProviderController) Lister() GlobalDNSProviderLister {
	return &globalDNSProviderLister{
		controller: c,
	}
}

func (c *globalDNSProviderController) AddHandler(ctx context.Context, name string, handler GlobalDNSProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalDNSProvider); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *globalDNSProviderController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler GlobalDNSProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalDNSProvider); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type globalDNSProviderFactory struct {
}

func (c globalDNSProviderFactory) Object() runtime.Object {
	return &GlobalDNSProvider{}
}

func (c globalDNSProviderFactory) List() runtime.Object {
	return &GlobalDNSProviderList{}
}

func (s *globalDNSProviderClient) Controller() GlobalDNSProviderController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.globalDNSProviderControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(GlobalDNSProviderGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &globalDNSProviderController{
		GenericController: genericController,
	}

	s.client.globalDNSProviderControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type globalDNSProviderClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   GlobalDNSProviderController
}

func (s *globalDNSProviderClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *globalDNSProviderClient) Create(o *GlobalDNSProvider) (*GlobalDNSProvider, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*GlobalDNSProvider), err
}

func (s *globalDNSProviderClient) Get(name string, opts metav1.GetOptions) (*GlobalDNSProvider, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*GlobalDNSProvider), err
}

func (s *globalDNSProviderClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalDNSProvider, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*GlobalDNSProvider), err
}

func (s *globalDNSProviderClient) Update(o *GlobalDNSProvider) (*GlobalDNSProvider, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*GlobalDNSProvider), err
}

func (s *globalDNSProviderClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *globalDNSProviderClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *globalDNSProviderClient) List(opts metav1.ListOptions) (*GlobalDNSProviderList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*GlobalDNSProviderList), err
}

func (s *globalDNSProviderClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *globalDNSProviderClient) Patch(o *GlobalDNSProvider, data []byte, subresources ...string) (*GlobalDNSProvider, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*GlobalDNSProvider), err
}

func (s *globalDNSProviderClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *globalDNSProviderClient) AddHandler(ctx context.Context, name string, sync GlobalDNSProviderHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *globalDNSProviderClient) AddLifecycle(ctx context.Context, name string, lifecycle GlobalDNSProviderLifecycle) {
	sync := NewGlobalDNSProviderLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *globalDNSProviderClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync GlobalDNSProviderHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *globalDNSProviderClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle GlobalDNSProviderLifecycle) {
	sync := NewGlobalDNSProviderLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type GlobalDNSProviderIndexer func(obj *GlobalDNSProvider) ([]string, error)

type GlobalDNSProviderClientCache interface {
	Get(namespace, name string) (*GlobalDNSProvider, error)
	List(namespace string, selector labels.Selector) ([]*GlobalDNSProvider, error)

	Index(name string, indexer GlobalDNSProviderIndexer)
	GetIndexed(name, key string) ([]*GlobalDNSProvider, error)
}

type GlobalDNSProviderClient interface {
	Create(*GlobalDNSProvider) (*GlobalDNSProvider, error)
	Get(namespace, name string, opts metav1.GetOptions) (*GlobalDNSProvider, error)
	Update(*GlobalDNSProvider) (*GlobalDNSProvider, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*GlobalDNSProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() GlobalDNSProviderClientCache

	OnCreate(ctx context.Context, name string, sync GlobalDNSProviderChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync GlobalDNSProviderChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync GlobalDNSProviderChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	Interface() GlobalDNSProviderInterface
}

type globalDNSProviderClientCache struct {
	client *globalDNSProviderClient2
}

type globalDNSProviderClient2 struct {
	iface      GlobalDNSProviderInterface
	controller GlobalDNSProviderController
}

func (n *globalDNSProviderClient2) Interface() GlobalDNSProviderInterface {
	return n.iface
}

func (n *globalDNSProviderClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *globalDNSProviderClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *globalDNSProviderClient2) Create(obj *GlobalDNSProvider) (*GlobalDNSProvider, error) {
	return n.iface.Create(obj)
}

func (n *globalDNSProviderClient2) Get(namespace, name string, opts metav1.GetOptions) (*GlobalDNSProvider, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *globalDNSProviderClient2) Update(obj *GlobalDNSProvider) (*GlobalDNSProvider, error) {
	return n.iface.Update(obj)
}

func (n *globalDNSProviderClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *globalDNSProviderClient2) List(namespace string, opts metav1.ListOptions) (*GlobalDNSProviderList, error) {
	return n.iface.List(opts)
}

func (n *globalDNSProviderClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *globalDNSProviderClientCache) Get(namespace, name string) (*GlobalDNSProvider, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *globalDNSProviderClientCache) List(namespace string, selector labels.Selector) ([]*GlobalDNSProvider, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *globalDNSProviderClient2) Cache() GlobalDNSProviderClientCache {
	n.loadController()
	return &globalDNSProviderClientCache{
		client: n,
	}
}

func (n *globalDNSProviderClient2) OnCreate(ctx context.Context, name string, sync GlobalDNSProviderChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &globalDNSProviderLifecycleDelegate{create: sync})
}

func (n *globalDNSProviderClient2) OnChange(ctx context.Context, name string, sync GlobalDNSProviderChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &globalDNSProviderLifecycleDelegate{update: sync})
}

func (n *globalDNSProviderClient2) OnRemove(ctx context.Context, name string, sync GlobalDNSProviderChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &globalDNSProviderLifecycleDelegate{remove: sync})
}

func (n *globalDNSProviderClientCache) Index(name string, indexer GlobalDNSProviderIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*GlobalDNSProvider); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *globalDNSProviderClientCache) GetIndexed(name, key string) ([]*GlobalDNSProvider, error) {
	var result []*GlobalDNSProvider
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*GlobalDNSProvider); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *globalDNSProviderClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type globalDNSProviderLifecycleDelegate struct {
	create GlobalDNSProviderChangeHandlerFunc
	update GlobalDNSProviderChangeHandlerFunc
	remove GlobalDNSProviderChangeHandlerFunc
}

func (n *globalDNSProviderLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *globalDNSProviderLifecycleDelegate) Create(obj *GlobalDNSProvider) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *globalDNSProviderLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *globalDNSProviderLifecycleDelegate) Remove(obj *GlobalDNSProvider) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *globalDNSProviderLifecycleDelegate) Updated(obj *GlobalDNSProvider) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
