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
	ComponentStatusGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ComponentStatus",
	}
	ComponentStatusResource = metav1.APIResource{
		Name:         "componentstatuses",
		SingularName: "componentstatus",
		Namespaced:   false,
		Kind:         ComponentStatusGroupVersionKind.Kind,
	}

	ComponentStatusGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "componentstatuses",
	}
)

func init() {
	resource.Put(ComponentStatusGroupVersionResource)
}

func NewComponentStatus(namespace, name string, obj v1.ComponentStatus) *v1.ComponentStatus {
	obj.APIVersion, obj.Kind = ComponentStatusGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type ComponentStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.ComponentStatus `json:"items"`
}

type ComponentStatusHandlerFunc func(key string, obj *v1.ComponentStatus) (runtime.Object, error)

type ComponentStatusChangeHandlerFunc func(obj *v1.ComponentStatus) (runtime.Object, error)

type ComponentStatusLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.ComponentStatus, err error)
	Get(namespace, name string) (*v1.ComponentStatus, error)
}

type ComponentStatusController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ComponentStatusLister
	AddHandler(ctx context.Context, name string, handler ComponentStatusHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ComponentStatusHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ComponentStatusHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ComponentStatusInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.ComponentStatus) (*v1.ComponentStatus, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ComponentStatus, error)
	Get(name string, opts metav1.GetOptions) (*v1.ComponentStatus, error)
	Update(*v1.ComponentStatus) (*v1.ComponentStatus, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ComponentStatusList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ComponentStatusController
	AddHandler(ctx context.Context, name string, sync ComponentStatusHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ComponentStatusHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ComponentStatusLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle ComponentStatusLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ComponentStatusHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ComponentStatusLifecycle)
}

type componentStatusLister struct {
	controller *componentStatusController
}

func (l *componentStatusLister) List(namespace string, selector labels.Selector) (ret []*v1.ComponentStatus, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.ComponentStatus))
	})
	return
}

func (l *componentStatusLister) Get(namespace, name string) (*v1.ComponentStatus, error) {
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
			Group:    ComponentStatusGroupVersionKind.Group,
			Resource: "componentStatus",
		}, key)
	}
	return obj.(*v1.ComponentStatus), nil
}

type componentStatusController struct {
	controller.GenericController
}

func (c *componentStatusController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *componentStatusController) Lister() ComponentStatusLister {
	return &componentStatusLister{
		controller: c,
	}
}

func (c *componentStatusController) AddHandler(ctx context.Context, name string, handler ComponentStatusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ComponentStatus); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *componentStatusController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler ComponentStatusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ComponentStatus); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *componentStatusController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ComponentStatusHandlerFunc) {
	resource.PutClusterScoped(ComponentStatusGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ComponentStatus); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type componentStatusFactory struct {
}

func (c componentStatusFactory) Object() runtime.Object {
	return &v1.ComponentStatus{}
}

func (c componentStatusFactory) List() runtime.Object {
	return &ComponentStatusList{}
}

func (s *componentStatusClient) Controller() ComponentStatusController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.componentStatusControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ComponentStatusGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &componentStatusController{
		GenericController: genericController,
	}

	s.client.componentStatusControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type componentStatusClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ComponentStatusController
}

func (s *componentStatusClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *componentStatusClient) Create(o *v1.ComponentStatus) (*v1.ComponentStatus, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.ComponentStatus), err
}

func (s *componentStatusClient) Get(name string, opts metav1.GetOptions) (*v1.ComponentStatus, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.ComponentStatus), err
}

func (s *componentStatusClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ComponentStatus, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.ComponentStatus), err
}

func (s *componentStatusClient) Update(o *v1.ComponentStatus) (*v1.ComponentStatus, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.ComponentStatus), err
}

func (s *componentStatusClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *componentStatusClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *componentStatusClient) List(opts metav1.ListOptions) (*ComponentStatusList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ComponentStatusList), err
}

func (s *componentStatusClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *componentStatusClient) Patch(o *v1.ComponentStatus, patchType types.PatchType, data []byte, subresources ...string) (*v1.ComponentStatus, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.ComponentStatus), err
}

func (s *componentStatusClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *componentStatusClient) AddHandler(ctx context.Context, name string, sync ComponentStatusHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *componentStatusClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync ComponentStatusHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *componentStatusClient) AddLifecycle(ctx context.Context, name string, lifecycle ComponentStatusLifecycle) {
	sync := NewComponentStatusLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *componentStatusClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle ComponentStatusLifecycle) {
	sync := NewComponentStatusLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *componentStatusClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ComponentStatusHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *componentStatusClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ComponentStatusLifecycle) {
	sync := NewComponentStatusLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type ComponentStatusIndexer func(obj *v1.ComponentStatus) ([]string, error)

type ComponentStatusClientCache interface {
	Get(namespace, name string) (*v1.ComponentStatus, error)
	List(namespace string, selector labels.Selector) ([]*v1.ComponentStatus, error)

	Index(name string, indexer ComponentStatusIndexer)
	GetIndexed(name, key string) ([]*v1.ComponentStatus, error)
}

type ComponentStatusClient interface {
	Create(*v1.ComponentStatus) (*v1.ComponentStatus, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1.ComponentStatus, error)
	Update(*v1.ComponentStatus) (*v1.ComponentStatus, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*ComponentStatusList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() ComponentStatusClientCache

	OnCreate(ctx context.Context, name string, sync ComponentStatusChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync ComponentStatusChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync ComponentStatusChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() ComponentStatusInterface
}

type componentStatusClientCache struct {
	client *componentStatusClient2
}

type componentStatusClient2 struct {
	iface      ComponentStatusInterface
	controller ComponentStatusController
}

func (n *componentStatusClient2) Interface() ComponentStatusInterface {
	return n.iface
}

func (n *componentStatusClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *componentStatusClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *componentStatusClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *componentStatusClient2) Create(obj *v1.ComponentStatus) (*v1.ComponentStatus, error) {
	return n.iface.Create(obj)
}

func (n *componentStatusClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1.ComponentStatus, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *componentStatusClient2) Update(obj *v1.ComponentStatus) (*v1.ComponentStatus, error) {
	return n.iface.Update(obj)
}

func (n *componentStatusClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *componentStatusClient2) List(namespace string, opts metav1.ListOptions) (*ComponentStatusList, error) {
	return n.iface.List(opts)
}

func (n *componentStatusClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *componentStatusClientCache) Get(namespace, name string) (*v1.ComponentStatus, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *componentStatusClientCache) List(namespace string, selector labels.Selector) ([]*v1.ComponentStatus, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *componentStatusClient2) Cache() ComponentStatusClientCache {
	n.loadController()
	return &componentStatusClientCache{
		client: n,
	}
}

func (n *componentStatusClient2) OnCreate(ctx context.Context, name string, sync ComponentStatusChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &componentStatusLifecycleDelegate{create: sync})
}

func (n *componentStatusClient2) OnChange(ctx context.Context, name string, sync ComponentStatusChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &componentStatusLifecycleDelegate{update: sync})
}

func (n *componentStatusClient2) OnRemove(ctx context.Context, name string, sync ComponentStatusChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &componentStatusLifecycleDelegate{remove: sync})
}

func (n *componentStatusClientCache) Index(name string, indexer ComponentStatusIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1.ComponentStatus); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *componentStatusClientCache) GetIndexed(name, key string) ([]*v1.ComponentStatus, error) {
	var result []*v1.ComponentStatus
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1.ComponentStatus); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *componentStatusClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type componentStatusLifecycleDelegate struct {
	create ComponentStatusChangeHandlerFunc
	update ComponentStatusChangeHandlerFunc
	remove ComponentStatusChangeHandlerFunc
}

func (n *componentStatusLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *componentStatusLifecycleDelegate) Create(obj *v1.ComponentStatus) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *componentStatusLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *componentStatusLifecycleDelegate) Remove(obj *v1.ComponentStatus) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *componentStatusLifecycleDelegate) Updated(obj *v1.ComponentStatus) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
