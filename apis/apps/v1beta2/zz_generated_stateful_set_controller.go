package v1beta2

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	"k8s.io/api/apps/v1beta2"
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
	StatefulSetGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "StatefulSet",
	}
	StatefulSetResource = metav1.APIResource{
		Name:         "statefulsets",
		SingularName: "statefulset",
		Namespaced:   true,

		Kind: StatefulSetGroupVersionKind.Kind,
	}

	StatefulSetGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "statefulsets",
	}
)

func init() {
	resource.Put(StatefulSetGroupVersionResource)
}

func NewStatefulSet(namespace, name string, obj v1beta2.StatefulSet) *v1beta2.StatefulSet {
	obj.APIVersion, obj.Kind = StatefulSetGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type StatefulSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1beta2.StatefulSet `json:"items"`
}

type StatefulSetHandlerFunc func(key string, obj *v1beta2.StatefulSet) (runtime.Object, error)

type StatefulSetChangeHandlerFunc func(obj *v1beta2.StatefulSet) (runtime.Object, error)

type StatefulSetLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1beta2.StatefulSet, err error)
	Get(namespace, name string) (*v1beta2.StatefulSet, error)
}

type StatefulSetController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() StatefulSetLister
	AddHandler(ctx context.Context, name string, handler StatefulSetHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync StatefulSetHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler StatefulSetHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type StatefulSetInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1beta2.StatefulSet) (*v1beta2.StatefulSet, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1beta2.StatefulSet, error)
	Get(name string, opts metav1.GetOptions) (*v1beta2.StatefulSet, error)
	Update(*v1beta2.StatefulSet) (*v1beta2.StatefulSet, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*StatefulSetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() StatefulSetController
	AddHandler(ctx context.Context, name string, sync StatefulSetHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync StatefulSetHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle StatefulSetLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle StatefulSetLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync StatefulSetHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle StatefulSetLifecycle)
}

type statefulSetLister struct {
	controller *statefulSetController
}

func (l *statefulSetLister) List(namespace string, selector labels.Selector) (ret []*v1beta2.StatefulSet, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1beta2.StatefulSet))
	})
	return
}

func (l *statefulSetLister) Get(namespace, name string) (*v1beta2.StatefulSet, error) {
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
			Group:    StatefulSetGroupVersionKind.Group,
			Resource: "statefulSet",
		}, key)
	}
	return obj.(*v1beta2.StatefulSet), nil
}

type statefulSetController struct {
	controller.GenericController
}

func (c *statefulSetController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *statefulSetController) Lister() StatefulSetLister {
	return &statefulSetLister{
		controller: c,
	}
}

func (c *statefulSetController) AddHandler(ctx context.Context, name string, handler StatefulSetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1beta2.StatefulSet); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *statefulSetController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler StatefulSetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1beta2.StatefulSet); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *statefulSetController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler StatefulSetHandlerFunc) {
	resource.PutClusterScoped(StatefulSetGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1beta2.StatefulSet); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type statefulSetFactory struct {
}

func (c statefulSetFactory) Object() runtime.Object {
	return &v1beta2.StatefulSet{}
}

func (c statefulSetFactory) List() runtime.Object {
	return &StatefulSetList{}
}

func (s *statefulSetClient) Controller() StatefulSetController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.statefulSetControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(StatefulSetGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &statefulSetController{
		GenericController: genericController,
	}

	s.client.statefulSetControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type statefulSetClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   StatefulSetController
}

func (s *statefulSetClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *statefulSetClient) Create(o *v1beta2.StatefulSet) (*v1beta2.StatefulSet, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1beta2.StatefulSet), err
}

func (s *statefulSetClient) Get(name string, opts metav1.GetOptions) (*v1beta2.StatefulSet, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1beta2.StatefulSet), err
}

func (s *statefulSetClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1beta2.StatefulSet, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1beta2.StatefulSet), err
}

func (s *statefulSetClient) Update(o *v1beta2.StatefulSet) (*v1beta2.StatefulSet, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1beta2.StatefulSet), err
}

func (s *statefulSetClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *statefulSetClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *statefulSetClient) List(opts metav1.ListOptions) (*StatefulSetList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*StatefulSetList), err
}

func (s *statefulSetClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *statefulSetClient) Patch(o *v1beta2.StatefulSet, patchType types.PatchType, data []byte, subresources ...string) (*v1beta2.StatefulSet, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1beta2.StatefulSet), err
}

func (s *statefulSetClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *statefulSetClient) AddHandler(ctx context.Context, name string, sync StatefulSetHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *statefulSetClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync StatefulSetHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *statefulSetClient) AddLifecycle(ctx context.Context, name string, lifecycle StatefulSetLifecycle) {
	sync := NewStatefulSetLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *statefulSetClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle StatefulSetLifecycle) {
	sync := NewStatefulSetLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *statefulSetClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync StatefulSetHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *statefulSetClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle StatefulSetLifecycle) {
	sync := NewStatefulSetLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type StatefulSetIndexer func(obj *v1beta2.StatefulSet) ([]string, error)

type StatefulSetClientCache interface {
	Get(namespace, name string) (*v1beta2.StatefulSet, error)
	List(namespace string, selector labels.Selector) ([]*v1beta2.StatefulSet, error)

	Index(name string, indexer StatefulSetIndexer)
	GetIndexed(name, key string) ([]*v1beta2.StatefulSet, error)
}

type StatefulSetClient interface {
	Create(*v1beta2.StatefulSet) (*v1beta2.StatefulSet, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1beta2.StatefulSet, error)
	Update(*v1beta2.StatefulSet) (*v1beta2.StatefulSet, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*StatefulSetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() StatefulSetClientCache

	OnCreate(ctx context.Context, name string, sync StatefulSetChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync StatefulSetChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync StatefulSetChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() StatefulSetInterface
}

type statefulSetClientCache struct {
	client *statefulSetClient2
}

type statefulSetClient2 struct {
	iface      StatefulSetInterface
	controller StatefulSetController
}

func (n *statefulSetClient2) Interface() StatefulSetInterface {
	return n.iface
}

func (n *statefulSetClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *statefulSetClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *statefulSetClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *statefulSetClient2) Create(obj *v1beta2.StatefulSet) (*v1beta2.StatefulSet, error) {
	return n.iface.Create(obj)
}

func (n *statefulSetClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1beta2.StatefulSet, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *statefulSetClient2) Update(obj *v1beta2.StatefulSet) (*v1beta2.StatefulSet, error) {
	return n.iface.Update(obj)
}

func (n *statefulSetClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *statefulSetClient2) List(namespace string, opts metav1.ListOptions) (*StatefulSetList, error) {
	return n.iface.List(opts)
}

func (n *statefulSetClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *statefulSetClientCache) Get(namespace, name string) (*v1beta2.StatefulSet, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *statefulSetClientCache) List(namespace string, selector labels.Selector) ([]*v1beta2.StatefulSet, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *statefulSetClient2) Cache() StatefulSetClientCache {
	n.loadController()
	return &statefulSetClientCache{
		client: n,
	}
}

func (n *statefulSetClient2) OnCreate(ctx context.Context, name string, sync StatefulSetChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &statefulSetLifecycleDelegate{create: sync})
}

func (n *statefulSetClient2) OnChange(ctx context.Context, name string, sync StatefulSetChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &statefulSetLifecycleDelegate{update: sync})
}

func (n *statefulSetClient2) OnRemove(ctx context.Context, name string, sync StatefulSetChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &statefulSetLifecycleDelegate{remove: sync})
}

func (n *statefulSetClientCache) Index(name string, indexer StatefulSetIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1beta2.StatefulSet); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *statefulSetClientCache) GetIndexed(name, key string) ([]*v1beta2.StatefulSet, error) {
	var result []*v1beta2.StatefulSet
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1beta2.StatefulSet); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *statefulSetClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type statefulSetLifecycleDelegate struct {
	create StatefulSetChangeHandlerFunc
	update StatefulSetChangeHandlerFunc
	remove StatefulSetChangeHandlerFunc
}

func (n *statefulSetLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *statefulSetLifecycleDelegate) Create(obj *v1beta2.StatefulSet) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *statefulSetLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *statefulSetLifecycleDelegate) Remove(obj *v1beta2.StatefulSet) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *statefulSetLifecycleDelegate) Updated(obj *v1beta2.StatefulSet) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
