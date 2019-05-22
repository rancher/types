package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ExampleConfigGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ExampleConfig",
	}
	ExampleConfigResource = metav1.APIResource{
		Name:         "exampleconfigs",
		SingularName: "exampleconfig",
		Namespaced:   false,
		Kind:         ExampleConfigGroupVersionKind.Kind,
	}

	ExampleConfigGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "exampleconfigs",
	}
)

func init() {
	resource.Put(ExampleConfigGroupVersionResource)
}

func NewExampleConfig(namespace, name string, obj ExampleConfig) *ExampleConfig {
	obj.APIVersion, obj.Kind = ExampleConfigGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type ExampleConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExampleConfig `json:"items"`
}

type ExampleConfigHandlerFunc func(key string, obj *ExampleConfig) (runtime.Object, error)

type ExampleConfigChangeHandlerFunc func(obj *ExampleConfig) (runtime.Object, error)

type ExampleConfigLister interface {
	List(namespace string, selector labels.Selector) (ret []*ExampleConfig, err error)
	Get(namespace, name string) (*ExampleConfig, error)
}

type ExampleConfigController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ExampleConfigLister
	AddHandler(ctx context.Context, name string, handler ExampleConfigHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ExampleConfigHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ExampleConfigHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ExampleConfigInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ExampleConfig) (*ExampleConfig, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ExampleConfig, error)
	Get(name string, opts metav1.GetOptions) (*ExampleConfig, error)
	Update(*ExampleConfig) (*ExampleConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ExampleConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ExampleConfigController
	AddHandler(ctx context.Context, name string, sync ExampleConfigHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ExampleConfigHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ExampleConfigLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle ExampleConfigLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ExampleConfigHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ExampleConfigLifecycle)
}

type exampleConfigLister struct {
	controller *exampleConfigController
}

func (l *exampleConfigLister) List(namespace string, selector labels.Selector) (ret []*ExampleConfig, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ExampleConfig))
	})
	return
}

func (l *exampleConfigLister) Get(namespace, name string) (*ExampleConfig, error) {
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
			Group:    ExampleConfigGroupVersionKind.Group,
			Resource: "exampleConfig",
		}, key)
	}
	return obj.(*ExampleConfig), nil
}

type exampleConfigController struct {
	controller.GenericController
}

func (c *exampleConfigController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *exampleConfigController) Lister() ExampleConfigLister {
	return &exampleConfigLister{
		controller: c,
	}
}

func (c *exampleConfigController) AddHandler(ctx context.Context, name string, handler ExampleConfigHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ExampleConfig); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *exampleConfigController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler ExampleConfigHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ExampleConfig); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *exampleConfigController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ExampleConfigHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ExampleConfig); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type exampleConfigFactory struct {
}

func (c exampleConfigFactory) Object() runtime.Object {
	return &ExampleConfig{}
}

func (c exampleConfigFactory) List() runtime.Object {
	return &ExampleConfigList{}
}

func (s *exampleConfigClient) Controller() ExampleConfigController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.exampleConfigControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ExampleConfigGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &exampleConfigController{
		GenericController: genericController,
	}

	s.client.exampleConfigControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type exampleConfigClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ExampleConfigController
}

func (s *exampleConfigClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *exampleConfigClient) Create(o *ExampleConfig) (*ExampleConfig, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ExampleConfig), err
}

func (s *exampleConfigClient) Get(name string, opts metav1.GetOptions) (*ExampleConfig, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ExampleConfig), err
}

func (s *exampleConfigClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ExampleConfig, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ExampleConfig), err
}

func (s *exampleConfigClient) Update(o *ExampleConfig) (*ExampleConfig, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ExampleConfig), err
}

func (s *exampleConfigClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *exampleConfigClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *exampleConfigClient) List(opts metav1.ListOptions) (*ExampleConfigList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ExampleConfigList), err
}

func (s *exampleConfigClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *exampleConfigClient) Patch(o *ExampleConfig, patchType types.PatchType, data []byte, subresources ...string) (*ExampleConfig, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*ExampleConfig), err
}

func (s *exampleConfigClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *exampleConfigClient) AddHandler(ctx context.Context, name string, sync ExampleConfigHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *exampleConfigClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ExampleConfigHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *exampleConfigClient) AddLifecycle(ctx context.Context, name string, lifecycle ExampleConfigLifecycle) {
	sync := NewExampleConfigLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *exampleConfigClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle ExampleConfigLifecycle) {
	sync := NewExampleConfigLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *exampleConfigClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ExampleConfigHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *exampleConfigClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ExampleConfigLifecycle) {
	sync := NewExampleConfigLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type ExampleConfigIndexer func(obj *ExampleConfig) ([]string, error)

type ExampleConfigClientCache interface {
	Get(namespace, name string) (*ExampleConfig, error)
	List(namespace string, selector labels.Selector) ([]*ExampleConfig, error)

	Index(name string, indexer ExampleConfigIndexer)
	GetIndexed(name, key string) ([]*ExampleConfig, error)
}

type ExampleConfigClient interface {
	Create(*ExampleConfig) (*ExampleConfig, error)
	Get(namespace, name string, opts metav1.GetOptions) (*ExampleConfig, error)
	Update(*ExampleConfig) (*ExampleConfig, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*ExampleConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() ExampleConfigClientCache

	OnCreate(ctx context.Context, name string, sync ExampleConfigChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync ExampleConfigChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync ExampleConfigChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() ExampleConfigInterface
}

type exampleConfigClientCache struct {
	client *exampleConfigClient2
}

type exampleConfigClient2 struct {
	iface      ExampleConfigInterface
	controller ExampleConfigController
}

func (n *exampleConfigClient2) Interface() ExampleConfigInterface {
	return n.iface
}

func (n *exampleConfigClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *exampleConfigClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *exampleConfigClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *exampleConfigClient2) Create(obj *ExampleConfig) (*ExampleConfig, error) {
	return n.iface.Create(obj)
}

func (n *exampleConfigClient2) Get(namespace, name string, opts metav1.GetOptions) (*ExampleConfig, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *exampleConfigClient2) Update(obj *ExampleConfig) (*ExampleConfig, error) {
	return n.iface.Update(obj)
}

func (n *exampleConfigClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *exampleConfigClient2) List(namespace string, opts metav1.ListOptions) (*ExampleConfigList, error) {
	return n.iface.List(opts)
}

func (n *exampleConfigClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *exampleConfigClientCache) Get(namespace, name string) (*ExampleConfig, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *exampleConfigClientCache) List(namespace string, selector labels.Selector) ([]*ExampleConfig, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *exampleConfigClient2) Cache() ExampleConfigClientCache {
	n.loadController()
	return &exampleConfigClientCache{
		client: n,
	}
}

func (n *exampleConfigClient2) OnCreate(ctx context.Context, name string, sync ExampleConfigChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &exampleConfigLifecycleDelegate{create: sync})
}

func (n *exampleConfigClient2) OnChange(ctx context.Context, name string, sync ExampleConfigChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &exampleConfigLifecycleDelegate{update: sync})
}

func (n *exampleConfigClient2) OnRemove(ctx context.Context, name string, sync ExampleConfigChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &exampleConfigLifecycleDelegate{remove: sync})
}

func (n *exampleConfigClientCache) Index(name string, indexer ExampleConfigIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*ExampleConfig); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *exampleConfigClientCache) GetIndexed(name, key string) ([]*ExampleConfig, error) {
	var result []*ExampleConfig
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*ExampleConfig); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *exampleConfigClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type exampleConfigLifecycleDelegate struct {
	create ExampleConfigChangeHandlerFunc
	update ExampleConfigChangeHandlerFunc
	remove ExampleConfigChangeHandlerFunc
}

func (n *exampleConfigLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *exampleConfigLifecycleDelegate) Create(obj *ExampleConfig) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *exampleConfigLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *exampleConfigLifecycleDelegate) Remove(obj *ExampleConfig) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *exampleConfigLifecycleDelegate) Updated(obj *ExampleConfig) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
