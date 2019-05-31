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
	ReplicationControllerGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ReplicationController",
	}
	ReplicationControllerResource = metav1.APIResource{
		Name:         "replicationcontrollers",
		SingularName: "replicationcontroller",
		Namespaced:   true,

		Kind: ReplicationControllerGroupVersionKind.Kind,
	}

	ReplicationControllerGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "replicationcontrollers",
	}
)

func init() {
	resource.Put(ReplicationControllerGroupVersionResource)
}

func NewReplicationController(namespace, name string, obj v1.ReplicationController) *v1.ReplicationController {
	obj.APIVersion, obj.Kind = ReplicationControllerGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type ReplicationControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.ReplicationController `json:"items"`
}

type ReplicationControllerHandlerFunc func(key string, obj *v1.ReplicationController) (runtime.Object, error)

type ReplicationControllerChangeHandlerFunc func(obj *v1.ReplicationController) (runtime.Object, error)

type ReplicationControllerLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.ReplicationController, err error)
	Get(namespace, name string) (*v1.ReplicationController, error)
}

type ReplicationControllerController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ReplicationControllerLister
	AddHandler(ctx context.Context, name string, handler ReplicationControllerHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ReplicationControllerHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ReplicationControllerHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ReplicationControllerInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.ReplicationController) (*v1.ReplicationController, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ReplicationController, error)
	Get(name string, opts metav1.GetOptions) (*v1.ReplicationController, error)
	Update(*v1.ReplicationController) (*v1.ReplicationController, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ReplicationControllerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ReplicationControllerController
	AddHandler(ctx context.Context, name string, sync ReplicationControllerHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ReplicationControllerHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ReplicationControllerLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle ReplicationControllerLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ReplicationControllerHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ReplicationControllerLifecycle)
}

type replicationControllerLister struct {
	controller *replicationControllerController
}

func (l *replicationControllerLister) List(namespace string, selector labels.Selector) (ret []*v1.ReplicationController, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.ReplicationController))
	})
	return
}

func (l *replicationControllerLister) Get(namespace, name string) (*v1.ReplicationController, error) {
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
			Group:    ReplicationControllerGroupVersionKind.Group,
			Resource: "replicationController",
		}, key)
	}
	return obj.(*v1.ReplicationController), nil
}

type replicationControllerController struct {
	controller.GenericController
}

func (c *replicationControllerController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *replicationControllerController) Lister() ReplicationControllerLister {
	return &replicationControllerLister{
		controller: c,
	}
}

func (c *replicationControllerController) AddHandler(ctx context.Context, name string, handler ReplicationControllerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ReplicationController); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *replicationControllerController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler ReplicationControllerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ReplicationController); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *replicationControllerController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ReplicationControllerHandlerFunc) {
	resource.PutClusterScoped(ReplicationControllerGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ReplicationController); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type replicationControllerFactory struct {
}

func (c replicationControllerFactory) Object() runtime.Object {
	return &v1.ReplicationController{}
}

func (c replicationControllerFactory) List() runtime.Object {
	return &ReplicationControllerList{}
}

func (s *replicationControllerClient) Controller() ReplicationControllerController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.replicationControllerControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ReplicationControllerGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &replicationControllerController{
		GenericController: genericController,
	}

	s.client.replicationControllerControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type replicationControllerClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ReplicationControllerController
}

func (s *replicationControllerClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *replicationControllerClient) Create(o *v1.ReplicationController) (*v1.ReplicationController, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.ReplicationController), err
}

func (s *replicationControllerClient) Get(name string, opts metav1.GetOptions) (*v1.ReplicationController, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.ReplicationController), err
}

func (s *replicationControllerClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ReplicationController, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.ReplicationController), err
}

func (s *replicationControllerClient) Update(o *v1.ReplicationController) (*v1.ReplicationController, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.ReplicationController), err
}

func (s *replicationControllerClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *replicationControllerClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *replicationControllerClient) List(opts metav1.ListOptions) (*ReplicationControllerList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ReplicationControllerList), err
}

func (s *replicationControllerClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *replicationControllerClient) Patch(o *v1.ReplicationController, patchType types.PatchType, data []byte, subresources ...string) (*v1.ReplicationController, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.ReplicationController), err
}

func (s *replicationControllerClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *replicationControllerClient) AddHandler(ctx context.Context, name string, sync ReplicationControllerHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *replicationControllerClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ReplicationControllerHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *replicationControllerClient) AddLifecycle(ctx context.Context, name string, lifecycle ReplicationControllerLifecycle) {
	sync := NewReplicationControllerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *replicationControllerClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle ReplicationControllerLifecycle) {
	sync := NewReplicationControllerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *replicationControllerClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ReplicationControllerHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *replicationControllerClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ReplicationControllerLifecycle) {
	sync := NewReplicationControllerLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type ReplicationControllerIndexer func(obj *v1.ReplicationController) ([]string, error)

type ReplicationControllerClientCache interface {
	Get(namespace, name string) (*v1.ReplicationController, error)
	List(namespace string, selector labels.Selector) ([]*v1.ReplicationController, error)

	Index(name string, indexer ReplicationControllerIndexer)
	GetIndexed(name, key string) ([]*v1.ReplicationController, error)
}

type ReplicationControllerClient interface {
	Create(*v1.ReplicationController) (*v1.ReplicationController, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1.ReplicationController, error)
	Update(*v1.ReplicationController) (*v1.ReplicationController, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*ReplicationControllerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() ReplicationControllerClientCache

	OnCreate(ctx context.Context, name string, sync ReplicationControllerChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync ReplicationControllerChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync ReplicationControllerChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() ReplicationControllerInterface
}

type replicationControllerClientCache struct {
	client *replicationControllerClient2
}

type replicationControllerClient2 struct {
	iface      ReplicationControllerInterface
	controller ReplicationControllerController
}

func (n *replicationControllerClient2) Interface() ReplicationControllerInterface {
	return n.iface
}

func (n *replicationControllerClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *replicationControllerClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *replicationControllerClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *replicationControllerClient2) Create(obj *v1.ReplicationController) (*v1.ReplicationController, error) {
	return n.iface.Create(obj)
}

func (n *replicationControllerClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1.ReplicationController, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *replicationControllerClient2) Update(obj *v1.ReplicationController) (*v1.ReplicationController, error) {
	return n.iface.Update(obj)
}

func (n *replicationControllerClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *replicationControllerClient2) List(namespace string, opts metav1.ListOptions) (*ReplicationControllerList, error) {
	return n.iface.List(opts)
}

func (n *replicationControllerClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *replicationControllerClientCache) Get(namespace, name string) (*v1.ReplicationController, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *replicationControllerClientCache) List(namespace string, selector labels.Selector) ([]*v1.ReplicationController, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *replicationControllerClient2) Cache() ReplicationControllerClientCache {
	n.loadController()
	return &replicationControllerClientCache{
		client: n,
	}
}

func (n *replicationControllerClient2) OnCreate(ctx context.Context, name string, sync ReplicationControllerChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &replicationControllerLifecycleDelegate{create: sync})
}

func (n *replicationControllerClient2) OnChange(ctx context.Context, name string, sync ReplicationControllerChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &replicationControllerLifecycleDelegate{update: sync})
}

func (n *replicationControllerClient2) OnRemove(ctx context.Context, name string, sync ReplicationControllerChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &replicationControllerLifecycleDelegate{remove: sync})
}

func (n *replicationControllerClientCache) Index(name string, indexer ReplicationControllerIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1.ReplicationController); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *replicationControllerClientCache) GetIndexed(name, key string) ([]*v1.ReplicationController, error) {
	var result []*v1.ReplicationController
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1.ReplicationController); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *replicationControllerClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type replicationControllerLifecycleDelegate struct {
	create ReplicationControllerChangeHandlerFunc
	update ReplicationControllerChangeHandlerFunc
	remove ReplicationControllerChangeHandlerFunc
}

func (n *replicationControllerLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *replicationControllerLifecycleDelegate) Create(obj *v1.ReplicationController) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *replicationControllerLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *replicationControllerLifecycleDelegate) Remove(obj *v1.ReplicationController) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *replicationControllerLifecycleDelegate) Updated(obj *v1.ReplicationController) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
