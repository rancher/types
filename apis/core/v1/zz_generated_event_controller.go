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
	EventGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Event",
	}
	EventResource = metav1.APIResource{
		Name:         "events",
		SingularName: "event",
		Namespaced:   false,
		Kind:         EventGroupVersionKind.Kind,
	}

	EventGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "events",
	}
)

func init() {
	resource.Put(EventGroupVersionResource)
}

func NewEvent(namespace, name string, obj v1.Event) *v1.Event {
	obj.APIVersion, obj.Kind = EventGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type EventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.Event `json:"items"`
}

type EventHandlerFunc func(key string, obj *v1.Event) (runtime.Object, error)

type EventChangeHandlerFunc func(obj *v1.Event) (runtime.Object, error)

type EventLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.Event, err error)
	Get(namespace, name string) (*v1.Event, error)
}

type EventController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() EventLister
	AddHandler(ctx context.Context, name string, handler EventHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync EventHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler EventHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type EventInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.Event) (*v1.Event, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.Event, error)
	Get(name string, opts metav1.GetOptions) (*v1.Event, error)
	Update(*v1.Event) (*v1.Event, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*EventList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() EventController
	AddHandler(ctx context.Context, name string, sync EventHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync EventHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle EventLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle EventLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync EventHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle EventLifecycle)
}

type eventLister struct {
	controller *eventController
}

func (l *eventLister) List(namespace string, selector labels.Selector) (ret []*v1.Event, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.Event))
	})
	return
}

func (l *eventLister) Get(namespace, name string) (*v1.Event, error) {
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
			Group:    EventGroupVersionKind.Group,
			Resource: "event",
		}, key)
	}
	return obj.(*v1.Event), nil
}

type eventController struct {
	controller.GenericController
}

func (c *eventController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *eventController) Lister() EventLister {
	return &eventLister{
		controller: c,
	}
}

func (c *eventController) AddHandler(ctx context.Context, name string, handler EventHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Event); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *eventController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler EventHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Event); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *eventController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler EventHandlerFunc) {
	resource.PutClusterScoped(EventGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Event); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type eventFactory struct {
}

func (c eventFactory) Object() runtime.Object {
	return &v1.Event{}
}

func (c eventFactory) List() runtime.Object {
	return &EventList{}
}

func (s *eventClient) Controller() EventController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.eventControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(EventGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &eventController{
		GenericController: genericController,
	}

	s.client.eventControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type eventClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   EventController
}

func (s *eventClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *eventClient) Create(o *v1.Event) (*v1.Event, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.Event), err
}

func (s *eventClient) Get(name string, opts metav1.GetOptions) (*v1.Event, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.Event), err
}

func (s *eventClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.Event, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.Event), err
}

func (s *eventClient) Update(o *v1.Event) (*v1.Event, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.Event), err
}

func (s *eventClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *eventClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *eventClient) List(opts metav1.ListOptions) (*EventList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*EventList), err
}

func (s *eventClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *eventClient) Patch(o *v1.Event, patchType types.PatchType, data []byte, subresources ...string) (*v1.Event, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.Event), err
}

func (s *eventClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *eventClient) AddHandler(ctx context.Context, name string, sync EventHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *eventClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync EventHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *eventClient) AddLifecycle(ctx context.Context, name string, lifecycle EventLifecycle) {
	sync := NewEventLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *eventClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle EventLifecycle) {
	sync := NewEventLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *eventClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync EventHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *eventClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle EventLifecycle) {
	sync := NewEventLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type EventIndexer func(obj *v1.Event) ([]string, error)

type EventClientCache interface {
	Get(namespace, name string) (*v1.Event, error)
	List(namespace string, selector labels.Selector) ([]*v1.Event, error)

	Index(name string, indexer EventIndexer)
	GetIndexed(name, key string) ([]*v1.Event, error)
}

type EventClient interface {
	Create(*v1.Event) (*v1.Event, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1.Event, error)
	Update(*v1.Event) (*v1.Event, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*EventList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() EventClientCache

	OnCreate(ctx context.Context, name string, sync EventChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync EventChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync EventChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() EventInterface
}

type eventClientCache struct {
	client *eventClient2
}

type eventClient2 struct {
	iface      EventInterface
	controller EventController
}

func (n *eventClient2) Interface() EventInterface {
	return n.iface
}

func (n *eventClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *eventClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *eventClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *eventClient2) Create(obj *v1.Event) (*v1.Event, error) {
	return n.iface.Create(obj)
}

func (n *eventClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1.Event, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *eventClient2) Update(obj *v1.Event) (*v1.Event, error) {
	return n.iface.Update(obj)
}

func (n *eventClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *eventClient2) List(namespace string, opts metav1.ListOptions) (*EventList, error) {
	return n.iface.List(opts)
}

func (n *eventClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *eventClientCache) Get(namespace, name string) (*v1.Event, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *eventClientCache) List(namespace string, selector labels.Selector) ([]*v1.Event, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *eventClient2) Cache() EventClientCache {
	n.loadController()
	return &eventClientCache{
		client: n,
	}
}

func (n *eventClient2) OnCreate(ctx context.Context, name string, sync EventChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &eventLifecycleDelegate{create: sync})
}

func (n *eventClient2) OnChange(ctx context.Context, name string, sync EventChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &eventLifecycleDelegate{update: sync})
}

func (n *eventClient2) OnRemove(ctx context.Context, name string, sync EventChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &eventLifecycleDelegate{remove: sync})
}

func (n *eventClientCache) Index(name string, indexer EventIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1.Event); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *eventClientCache) GetIndexed(name, key string) ([]*v1.Event, error) {
	var result []*v1.Event
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1.Event); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *eventClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type eventLifecycleDelegate struct {
	create EventChangeHandlerFunc
	update EventChangeHandlerFunc
	remove EventChangeHandlerFunc
}

func (n *eventLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *eventLifecycleDelegate) Create(obj *v1.Event) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *eventLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *eventLifecycleDelegate) Remove(obj *v1.Event) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *eventLifecycleDelegate) Updated(obj *v1.Event) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
