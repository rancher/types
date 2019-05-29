package v1

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	v1 "k8s.io/api/core/v1"
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
	LimitRangeGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "LimitRange",
	}
	LimitRangeResource = metav1.APIResource{
		Name:         "limitranges",
		SingularName: "limitrange",
		Namespaced:   true,

		Kind: LimitRangeGroupVersionKind.Kind,
	}

	LimitRangeGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "limitranges",
	}
)

func init() {
	resource.Put(LimitRangeGroupVersionResource)
}

func NewLimitRange(namespace, name string, obj v1.LimitRange) *v1.LimitRange {
	obj.APIVersion, obj.Kind = LimitRangeGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type LimitRangeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.LimitRange `json:"items"`
}

type LimitRangeHandlerFunc func(key string, obj *v1.LimitRange) (runtime.Object, error)

type LimitRangeChangeHandlerFunc func(obj *v1.LimitRange) (runtime.Object, error)

type LimitRangeLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.LimitRange, err error)
	Get(namespace, name string) (*v1.LimitRange, error)
}

type LimitRangeController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() LimitRangeLister
	AddHandler(ctx context.Context, name string, handler LimitRangeHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync LimitRangeHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler LimitRangeHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type LimitRangeInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.LimitRange) (*v1.LimitRange, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.LimitRange, error)
	Get(name string, opts metav1.GetOptions) (*v1.LimitRange, error)
	Update(*v1.LimitRange) (*v1.LimitRange, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*LimitRangeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() LimitRangeController
	AddHandler(ctx context.Context, name string, sync LimitRangeHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync LimitRangeHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle LimitRangeLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle LimitRangeLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync LimitRangeHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle LimitRangeLifecycle)
}

type limitRangeLister struct {
	controller *limitRangeController
}

func (l *limitRangeLister) List(namespace string, selector labels.Selector) (ret []*v1.LimitRange, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.LimitRange))
	})
	return
}

func (l *limitRangeLister) Get(namespace, name string) (*v1.LimitRange, error) {
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
			Group:    LimitRangeGroupVersionKind.Group,
			Resource: "limitRange",
		}, key)
	}
	return obj.(*v1.LimitRange), nil
}

type limitRangeController struct {
	controller.GenericController
}

func (c *limitRangeController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *limitRangeController) Lister() LimitRangeLister {
	return &limitRangeLister{
		controller: c,
	}
}

func (c *limitRangeController) AddHandler(ctx context.Context, name string, handler LimitRangeHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.LimitRange); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *limitRangeController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler LimitRangeHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.LimitRange); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *limitRangeController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler LimitRangeHandlerFunc) {
	resource.PutClusterScoped(LimitRangeGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.LimitRange); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type limitRangeFactory struct {
}

func (c limitRangeFactory) Object() runtime.Object {
	return &v1.LimitRange{}
}

func (c limitRangeFactory) List() runtime.Object {
	return &LimitRangeList{}
}

func (s *limitRangeClient) Controller() LimitRangeController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.limitRangeControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(LimitRangeGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &limitRangeController{
		GenericController: genericController,
	}

	s.client.limitRangeControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type limitRangeClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   LimitRangeController
}

func (s *limitRangeClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *limitRangeClient) Create(o *v1.LimitRange) (*v1.LimitRange, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.LimitRange), err
}

func (s *limitRangeClient) Get(name string, opts metav1.GetOptions) (*v1.LimitRange, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.LimitRange), err
}

func (s *limitRangeClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.LimitRange, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.LimitRange), err
}

func (s *limitRangeClient) Update(o *v1.LimitRange) (*v1.LimitRange, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.LimitRange), err
}

func (s *limitRangeClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *limitRangeClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *limitRangeClient) List(opts metav1.ListOptions) (*LimitRangeList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*LimitRangeList), err
}

func (s *limitRangeClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *limitRangeClient) Patch(o *v1.LimitRange, patchType types.PatchType, data []byte, subresources ...string) (*v1.LimitRange, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.LimitRange), err
}

func (s *limitRangeClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *limitRangeClient) AddHandler(ctx context.Context, name string, sync LimitRangeHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *limitRangeClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync LimitRangeHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *limitRangeClient) AddLifecycle(ctx context.Context, name string, lifecycle LimitRangeLifecycle) {
	sync := NewLimitRangeLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *limitRangeClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle LimitRangeLifecycle) {
	sync := NewLimitRangeLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *limitRangeClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync LimitRangeHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *limitRangeClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle LimitRangeLifecycle) {
	sync := NewLimitRangeLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type LimitRangeIndexer func(obj *v1.LimitRange) ([]string, error)

type LimitRangeClientCache interface {
	Get(namespace, name string) (*v1.LimitRange, error)
	List(namespace string, selector labels.Selector) ([]*v1.LimitRange, error)

	Index(name string, indexer LimitRangeIndexer)
	GetIndexed(name, key string) ([]*v1.LimitRange, error)
}

type LimitRangeClient interface {
	Create(*v1.LimitRange) (*v1.LimitRange, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1.LimitRange, error)
	Update(*v1.LimitRange) (*v1.LimitRange, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*LimitRangeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() LimitRangeClientCache

	OnCreate(ctx context.Context, name string, sync LimitRangeChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync LimitRangeChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync LimitRangeChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() LimitRangeInterface
}

type limitRangeClientCache struct {
	client *limitRangeClient2
}

type limitRangeClient2 struct {
	iface      LimitRangeInterface
	controller LimitRangeController
}

func (n *limitRangeClient2) Interface() LimitRangeInterface {
	return n.iface
}

func (n *limitRangeClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *limitRangeClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *limitRangeClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *limitRangeClient2) Create(obj *v1.LimitRange) (*v1.LimitRange, error) {
	return n.iface.Create(obj)
}

func (n *limitRangeClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1.LimitRange, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *limitRangeClient2) Update(obj *v1.LimitRange) (*v1.LimitRange, error) {
	return n.iface.Update(obj)
}

func (n *limitRangeClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *limitRangeClient2) List(namespace string, opts metav1.ListOptions) (*LimitRangeList, error) {
	return n.iface.List(opts)
}

func (n *limitRangeClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *limitRangeClientCache) Get(namespace, name string) (*v1.LimitRange, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *limitRangeClientCache) List(namespace string, selector labels.Selector) ([]*v1.LimitRange, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *limitRangeClient2) Cache() LimitRangeClientCache {
	n.loadController()
	return &limitRangeClientCache{
		client: n,
	}
}

func (n *limitRangeClient2) OnCreate(ctx context.Context, name string, sync LimitRangeChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &limitRangeLifecycleDelegate{create: sync})
}

func (n *limitRangeClient2) OnChange(ctx context.Context, name string, sync LimitRangeChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &limitRangeLifecycleDelegate{update: sync})
}

func (n *limitRangeClient2) OnRemove(ctx context.Context, name string, sync LimitRangeChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &limitRangeLifecycleDelegate{remove: sync})
}

func (n *limitRangeClientCache) Index(name string, indexer LimitRangeIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1.LimitRange); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *limitRangeClientCache) GetIndexed(name, key string) ([]*v1.LimitRange, error) {
	var result []*v1.LimitRange
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1.LimitRange); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *limitRangeClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type limitRangeLifecycleDelegate struct {
	create LimitRangeChangeHandlerFunc
	update LimitRangeChangeHandlerFunc
	remove LimitRangeChangeHandlerFunc
}

func (n *limitRangeLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *limitRangeLifecycleDelegate) Create(obj *v1.LimitRange) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *limitRangeLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *limitRangeLifecycleDelegate) Remove(obj *v1.LimitRange) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *limitRangeLifecycleDelegate) Updated(obj *v1.LimitRange) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
