package v1

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	v1 "k8s.io/api/networking/v1"
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
	NetworkPolicyGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "NetworkPolicy",
	}
	NetworkPolicyResource = metav1.APIResource{
		Name:         "networkpolicies",
		SingularName: "networkpolicy",
		Namespaced:   true,

		Kind: NetworkPolicyGroupVersionKind.Kind,
	}

	NetworkPolicyGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "networkpolicies",
	}
)

func init() {
	resource.Put(NetworkPolicyGroupVersionResource)
}

func NewNetworkPolicy(namespace, name string, obj v1.NetworkPolicy) *v1.NetworkPolicy {
	obj.APIVersion, obj.Kind = NetworkPolicyGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type NetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.NetworkPolicy `json:"items"`
}

type NetworkPolicyHandlerFunc func(key string, obj *v1.NetworkPolicy) (runtime.Object, error)

type NetworkPolicyChangeHandlerFunc func(obj *v1.NetworkPolicy) (runtime.Object, error)

type NetworkPolicyLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.NetworkPolicy, err error)
	Get(namespace, name string) (*v1.NetworkPolicy, error)
}

type NetworkPolicyController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() NetworkPolicyLister
	AddHandler(ctx context.Context, name string, handler NetworkPolicyHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync NetworkPolicyHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler NetworkPolicyHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type NetworkPolicyInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.NetworkPolicy) (*v1.NetworkPolicy, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error)
	Get(name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error)
	Update(*v1.NetworkPolicy) (*v1.NetworkPolicy, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*NetworkPolicyList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() NetworkPolicyController
	AddHandler(ctx context.Context, name string, sync NetworkPolicyHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync NetworkPolicyHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle NetworkPolicyLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle NetworkPolicyLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync NetworkPolicyHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle NetworkPolicyLifecycle)
}

type networkPolicyLister struct {
	controller *networkPolicyController
}

func (l *networkPolicyLister) List(namespace string, selector labels.Selector) (ret []*v1.NetworkPolicy, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.NetworkPolicy))
	})
	return
}

func (l *networkPolicyLister) Get(namespace, name string) (*v1.NetworkPolicy, error) {
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
			Group:    NetworkPolicyGroupVersionKind.Group,
			Resource: "networkPolicy",
		}, key)
	}
	return obj.(*v1.NetworkPolicy), nil
}

type networkPolicyController struct {
	controller.GenericController
}

func (c *networkPolicyController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *networkPolicyController) Lister() NetworkPolicyLister {
	return &networkPolicyLister{
		controller: c,
	}
}

func (c *networkPolicyController) AddHandler(ctx context.Context, name string, handler NetworkPolicyHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.NetworkPolicy); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *networkPolicyController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler NetworkPolicyHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.NetworkPolicy); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *networkPolicyController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler NetworkPolicyHandlerFunc) {
	resource.PutClusterScoped(NetworkPolicyGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.NetworkPolicy); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type networkPolicyFactory struct {
}

func (c networkPolicyFactory) Object() runtime.Object {
	return &v1.NetworkPolicy{}
}

func (c networkPolicyFactory) List() runtime.Object {
	return &NetworkPolicyList{}
}

func (s *networkPolicyClient) Controller() NetworkPolicyController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.networkPolicyControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(NetworkPolicyGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &networkPolicyController{
		GenericController: genericController,
	}

	s.client.networkPolicyControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type networkPolicyClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   NetworkPolicyController
}

func (s *networkPolicyClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *networkPolicyClient) Create(o *v1.NetworkPolicy) (*v1.NetworkPolicy, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.NetworkPolicy), err
}

func (s *networkPolicyClient) Get(name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.NetworkPolicy), err
}

func (s *networkPolicyClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.NetworkPolicy), err
}

func (s *networkPolicyClient) Update(o *v1.NetworkPolicy) (*v1.NetworkPolicy, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.NetworkPolicy), err
}

func (s *networkPolicyClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *networkPolicyClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *networkPolicyClient) List(opts metav1.ListOptions) (*NetworkPolicyList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*NetworkPolicyList), err
}

func (s *networkPolicyClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *networkPolicyClient) Patch(o *v1.NetworkPolicy, patchType types.PatchType, data []byte, subresources ...string) (*v1.NetworkPolicy, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.NetworkPolicy), err
}

func (s *networkPolicyClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *networkPolicyClient) AddHandler(ctx context.Context, name string, sync NetworkPolicyHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *networkPolicyClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync NetworkPolicyHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *networkPolicyClient) AddLifecycle(ctx context.Context, name string, lifecycle NetworkPolicyLifecycle) {
	sync := NewNetworkPolicyLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *networkPolicyClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle NetworkPolicyLifecycle) {
	sync := NewNetworkPolicyLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *networkPolicyClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync NetworkPolicyHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *networkPolicyClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle NetworkPolicyLifecycle) {
	sync := NewNetworkPolicyLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type NetworkPolicyIndexer func(obj *v1.NetworkPolicy) ([]string, error)

type NetworkPolicyClientCache interface {
	Get(namespace, name string) (*v1.NetworkPolicy, error)
	List(namespace string, selector labels.Selector) ([]*v1.NetworkPolicy, error)

	Index(name string, indexer NetworkPolicyIndexer)
	GetIndexed(name, key string) ([]*v1.NetworkPolicy, error)
}

type NetworkPolicyClient interface {
	Create(*v1.NetworkPolicy) (*v1.NetworkPolicy, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error)
	Update(*v1.NetworkPolicy) (*v1.NetworkPolicy, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*NetworkPolicyList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() NetworkPolicyClientCache

	OnCreate(ctx context.Context, name string, sync NetworkPolicyChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync NetworkPolicyChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync NetworkPolicyChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() NetworkPolicyInterface
}

type networkPolicyClientCache struct {
	client *networkPolicyClient2
}

type networkPolicyClient2 struct {
	iface      NetworkPolicyInterface
	controller NetworkPolicyController
}

func (n *networkPolicyClient2) Interface() NetworkPolicyInterface {
	return n.iface
}

func (n *networkPolicyClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *networkPolicyClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *networkPolicyClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *networkPolicyClient2) Create(obj *v1.NetworkPolicy) (*v1.NetworkPolicy, error) {
	return n.iface.Create(obj)
}

func (n *networkPolicyClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *networkPolicyClient2) Update(obj *v1.NetworkPolicy) (*v1.NetworkPolicy, error) {
	return n.iface.Update(obj)
}

func (n *networkPolicyClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *networkPolicyClient2) List(namespace string, opts metav1.ListOptions) (*NetworkPolicyList, error) {
	return n.iface.List(opts)
}

func (n *networkPolicyClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *networkPolicyClientCache) Get(namespace, name string) (*v1.NetworkPolicy, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *networkPolicyClientCache) List(namespace string, selector labels.Selector) ([]*v1.NetworkPolicy, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *networkPolicyClient2) Cache() NetworkPolicyClientCache {
	n.loadController()
	return &networkPolicyClientCache{
		client: n,
	}
}

func (n *networkPolicyClient2) OnCreate(ctx context.Context, name string, sync NetworkPolicyChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &networkPolicyLifecycleDelegate{create: sync})
}

func (n *networkPolicyClient2) OnChange(ctx context.Context, name string, sync NetworkPolicyChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &networkPolicyLifecycleDelegate{update: sync})
}

func (n *networkPolicyClient2) OnRemove(ctx context.Context, name string, sync NetworkPolicyChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &networkPolicyLifecycleDelegate{remove: sync})
}

func (n *networkPolicyClientCache) Index(name string, indexer NetworkPolicyIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1.NetworkPolicy); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *networkPolicyClientCache) GetIndexed(name, key string) ([]*v1.NetworkPolicy, error) {
	var result []*v1.NetworkPolicy
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1.NetworkPolicy); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *networkPolicyClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type networkPolicyLifecycleDelegate struct {
	create NetworkPolicyChangeHandlerFunc
	update NetworkPolicyChangeHandlerFunc
	remove NetworkPolicyChangeHandlerFunc
}

func (n *networkPolicyLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *networkPolicyLifecycleDelegate) Create(obj *v1.NetworkPolicy) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *networkPolicyLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *networkPolicyLifecycleDelegate) Remove(obj *v1.NetworkPolicy) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *networkPolicyLifecycleDelegate) Updated(obj *v1.NetworkPolicy) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
