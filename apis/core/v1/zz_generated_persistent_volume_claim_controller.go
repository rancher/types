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
	PersistentVolumeClaimGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "PersistentVolumeClaim",
	}
	PersistentVolumeClaimResource = metav1.APIResource{
		Name:         "persistentvolumeclaims",
		SingularName: "persistentvolumeclaim",
		Namespaced:   true,

		Kind: PersistentVolumeClaimGroupVersionKind.Kind,
	}

	PersistentVolumeClaimGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "persistentvolumeclaims",
	}
)

func init() {
	resource.Put(PersistentVolumeClaimGroupVersionResource)
}

func NewPersistentVolumeClaim(namespace, name string, obj v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim {
	obj.APIVersion, obj.Kind = PersistentVolumeClaimGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type PersistentVolumeClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.PersistentVolumeClaim `json:"items"`
}

type PersistentVolumeClaimHandlerFunc func(key string, obj *v1.PersistentVolumeClaim) (runtime.Object, error)

type PersistentVolumeClaimChangeHandlerFunc func(obj *v1.PersistentVolumeClaim) (runtime.Object, error)

type PersistentVolumeClaimLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error)
	Get(namespace, name string) (*v1.PersistentVolumeClaim, error)
}

type PersistentVolumeClaimController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() PersistentVolumeClaimLister
	AddHandler(ctx context.Context, name string, handler PersistentVolumeClaimHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync PersistentVolumeClaimHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler PersistentVolumeClaimHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type PersistentVolumeClaimInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.PersistentVolumeClaim, error)
	Get(name string, opts metav1.GetOptions) (*v1.PersistentVolumeClaim, error)
	Update(*v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*PersistentVolumeClaimList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() PersistentVolumeClaimController
	AddHandler(ctx context.Context, name string, sync PersistentVolumeClaimHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync PersistentVolumeClaimHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle PersistentVolumeClaimLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle PersistentVolumeClaimLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PersistentVolumeClaimHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PersistentVolumeClaimLifecycle)
}

type persistentVolumeClaimLister struct {
	controller *persistentVolumeClaimController
}

func (l *persistentVolumeClaimLister) List(namespace string, selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.PersistentVolumeClaim))
	})
	return
}

func (l *persistentVolumeClaimLister) Get(namespace, name string) (*v1.PersistentVolumeClaim, error) {
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
			Group:    PersistentVolumeClaimGroupVersionKind.Group,
			Resource: "persistentVolumeClaim",
		}, key)
	}
	return obj.(*v1.PersistentVolumeClaim), nil
}

type persistentVolumeClaimController struct {
	controller.GenericController
}

func (c *persistentVolumeClaimController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *persistentVolumeClaimController) Lister() PersistentVolumeClaimLister {
	return &persistentVolumeClaimLister{
		controller: c,
	}
}

func (c *persistentVolumeClaimController) AddHandler(ctx context.Context, name string, handler PersistentVolumeClaimHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.PersistentVolumeClaim); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *persistentVolumeClaimController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler PersistentVolumeClaimHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.PersistentVolumeClaim); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *persistentVolumeClaimController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler PersistentVolumeClaimHandlerFunc) {
	resource.PutClusterScoped(PersistentVolumeClaimGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.PersistentVolumeClaim); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type persistentVolumeClaimFactory struct {
}

func (c persistentVolumeClaimFactory) Object() runtime.Object {
	return &v1.PersistentVolumeClaim{}
}

func (c persistentVolumeClaimFactory) List() runtime.Object {
	return &PersistentVolumeClaimList{}
}

func (s *persistentVolumeClaimClient) Controller() PersistentVolumeClaimController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.persistentVolumeClaimControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(PersistentVolumeClaimGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &persistentVolumeClaimController{
		GenericController: genericController,
	}

	s.client.persistentVolumeClaimControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type persistentVolumeClaimClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   PersistentVolumeClaimController
}

func (s *persistentVolumeClaimClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *persistentVolumeClaimClient) Create(o *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.PersistentVolumeClaim), err
}

func (s *persistentVolumeClaimClient) Get(name string, opts metav1.GetOptions) (*v1.PersistentVolumeClaim, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.PersistentVolumeClaim), err
}

func (s *persistentVolumeClaimClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.PersistentVolumeClaim, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.PersistentVolumeClaim), err
}

func (s *persistentVolumeClaimClient) Update(o *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.PersistentVolumeClaim), err
}

func (s *persistentVolumeClaimClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *persistentVolumeClaimClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *persistentVolumeClaimClient) List(opts metav1.ListOptions) (*PersistentVolumeClaimList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*PersistentVolumeClaimList), err
}

func (s *persistentVolumeClaimClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *persistentVolumeClaimClient) Patch(o *v1.PersistentVolumeClaim, patchType types.PatchType, data []byte, subresources ...string) (*v1.PersistentVolumeClaim, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.PersistentVolumeClaim), err
}

func (s *persistentVolumeClaimClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *persistentVolumeClaimClient) AddHandler(ctx context.Context, name string, sync PersistentVolumeClaimHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *persistentVolumeClaimClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync PersistentVolumeClaimHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *persistentVolumeClaimClient) AddLifecycle(ctx context.Context, name string, lifecycle PersistentVolumeClaimLifecycle) {
	sync := NewPersistentVolumeClaimLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *persistentVolumeClaimClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle PersistentVolumeClaimLifecycle) {
	sync := NewPersistentVolumeClaimLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *persistentVolumeClaimClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PersistentVolumeClaimHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *persistentVolumeClaimClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PersistentVolumeClaimLifecycle) {
	sync := NewPersistentVolumeClaimLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type PersistentVolumeClaimIndexer func(obj *v1.PersistentVolumeClaim) ([]string, error)

type PersistentVolumeClaimClientCache interface {
	Get(namespace, name string) (*v1.PersistentVolumeClaim, error)
	List(namespace string, selector labels.Selector) ([]*v1.PersistentVolumeClaim, error)

	Index(name string, indexer PersistentVolumeClaimIndexer)
	GetIndexed(name, key string) ([]*v1.PersistentVolumeClaim, error)
}

type PersistentVolumeClaimClient interface {
	Create(*v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1.PersistentVolumeClaim, error)
	Update(*v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*PersistentVolumeClaimList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() PersistentVolumeClaimClientCache

	OnCreate(ctx context.Context, name string, sync PersistentVolumeClaimChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync PersistentVolumeClaimChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync PersistentVolumeClaimChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() PersistentVolumeClaimInterface
}

type persistentVolumeClaimClientCache struct {
	client *persistentVolumeClaimClient2
}

type persistentVolumeClaimClient2 struct {
	iface      PersistentVolumeClaimInterface
	controller PersistentVolumeClaimController
}

func (n *persistentVolumeClaimClient2) Interface() PersistentVolumeClaimInterface {
	return n.iface
}

func (n *persistentVolumeClaimClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *persistentVolumeClaimClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *persistentVolumeClaimClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *persistentVolumeClaimClient2) Create(obj *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	return n.iface.Create(obj)
}

func (n *persistentVolumeClaimClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1.PersistentVolumeClaim, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *persistentVolumeClaimClient2) Update(obj *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	return n.iface.Update(obj)
}

func (n *persistentVolumeClaimClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *persistentVolumeClaimClient2) List(namespace string, opts metav1.ListOptions) (*PersistentVolumeClaimList, error) {
	return n.iface.List(opts)
}

func (n *persistentVolumeClaimClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *persistentVolumeClaimClientCache) Get(namespace, name string) (*v1.PersistentVolumeClaim, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *persistentVolumeClaimClientCache) List(namespace string, selector labels.Selector) ([]*v1.PersistentVolumeClaim, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *persistentVolumeClaimClient2) Cache() PersistentVolumeClaimClientCache {
	n.loadController()
	return &persistentVolumeClaimClientCache{
		client: n,
	}
}

func (n *persistentVolumeClaimClient2) OnCreate(ctx context.Context, name string, sync PersistentVolumeClaimChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &persistentVolumeClaimLifecycleDelegate{create: sync})
}

func (n *persistentVolumeClaimClient2) OnChange(ctx context.Context, name string, sync PersistentVolumeClaimChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &persistentVolumeClaimLifecycleDelegate{update: sync})
}

func (n *persistentVolumeClaimClient2) OnRemove(ctx context.Context, name string, sync PersistentVolumeClaimChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &persistentVolumeClaimLifecycleDelegate{remove: sync})
}

func (n *persistentVolumeClaimClientCache) Index(name string, indexer PersistentVolumeClaimIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1.PersistentVolumeClaim); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *persistentVolumeClaimClientCache) GetIndexed(name, key string) ([]*v1.PersistentVolumeClaim, error) {
	var result []*v1.PersistentVolumeClaim
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1.PersistentVolumeClaim); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *persistentVolumeClaimClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type persistentVolumeClaimLifecycleDelegate struct {
	create PersistentVolumeClaimChangeHandlerFunc
	update PersistentVolumeClaimChangeHandlerFunc
	remove PersistentVolumeClaimChangeHandlerFunc
}

func (n *persistentVolumeClaimLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *persistentVolumeClaimLifecycleDelegate) Create(obj *v1.PersistentVolumeClaim) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *persistentVolumeClaimLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *persistentVolumeClaimLifecycleDelegate) Remove(obj *v1.PersistentVolumeClaim) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *persistentVolumeClaimLifecycleDelegate) Updated(obj *v1.PersistentVolumeClaim) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
