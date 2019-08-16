package v1alpha3

import (
	"context"

	"github.com/knative/pkg/apis/istio/v1alpha3"
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
	VirtualServiceGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "VirtualService",
	}
	VirtualServiceResource = metav1.APIResource{
		Name:         "virtualservices",
		SingularName: "virtualservice",
		Namespaced:   true,

		Kind: VirtualServiceGroupVersionKind.Kind,
	}

	VirtualServiceGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "virtualservices",
	}
)

func init() {
	resource.Put(VirtualServiceGroupVersionResource)
}

func NewVirtualService(namespace, name string, obj v1alpha3.VirtualService) *v1alpha3.VirtualService {
	obj.APIVersion, obj.Kind = VirtualServiceGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1alpha3.VirtualService `json:"items"`
}

type VirtualServiceHandlerFunc func(key string, obj *v1alpha3.VirtualService) (runtime.Object, error)

type VirtualServiceChangeHandlerFunc func(obj *v1alpha3.VirtualService) (runtime.Object, error)

type VirtualServiceLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1alpha3.VirtualService, err error)
	Get(namespace, name string) (*v1alpha3.VirtualService, error)
}

type VirtualServiceController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() VirtualServiceLister
	AddHandler(ctx context.Context, name string, handler VirtualServiceHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync VirtualServiceHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler VirtualServiceHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler VirtualServiceHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type VirtualServiceInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1alpha3.VirtualService, error)
	Get(name string, opts metav1.GetOptions) (*v1alpha3.VirtualService, error)
	Update(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*VirtualServiceList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() VirtualServiceController
	AddHandler(ctx context.Context, name string, sync VirtualServiceHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync VirtualServiceHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle VirtualServiceLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle VirtualServiceLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync VirtualServiceHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync VirtualServiceHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle VirtualServiceLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle VirtualServiceLifecycle)
}

type virtualServiceLister struct {
	controller *virtualServiceController
}

func (l *virtualServiceLister) List(namespace string, selector labels.Selector) (ret []*v1alpha3.VirtualService, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1alpha3.VirtualService))
	})
	return
}

func (l *virtualServiceLister) Get(namespace, name string) (*v1alpha3.VirtualService, error) {
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
			Group:    VirtualServiceGroupVersionKind.Group,
			Resource: "virtualService",
		}, key)
	}
	return obj.(*v1alpha3.VirtualService), nil
}

type virtualServiceController struct {
	controller.GenericController
}

func (c *virtualServiceController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *virtualServiceController) Lister() VirtualServiceLister {
	return &virtualServiceLister{
		controller: c,
	}
}

func (c *virtualServiceController) AddHandler(ctx context.Context, name string, handler VirtualServiceHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha3.VirtualService); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *virtualServiceController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler VirtualServiceHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha3.VirtualService); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *virtualServiceController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler VirtualServiceHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha3.VirtualService); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *virtualServiceController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler VirtualServiceHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha3.VirtualService); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type virtualServiceFactory struct {
}

func (c virtualServiceFactory) Object() runtime.Object {
	return &v1alpha3.VirtualService{}
}

func (c virtualServiceFactory) List() runtime.Object {
	return &VirtualServiceList{}
}

func (s *virtualServiceClient) Controller() VirtualServiceController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.virtualServiceControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(VirtualServiceGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &virtualServiceController{
		GenericController: genericController,
	}

	s.client.virtualServiceControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type virtualServiceClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   VirtualServiceController
}

func (s *virtualServiceClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *virtualServiceClient) Create(o *v1alpha3.VirtualService) (*v1alpha3.VirtualService, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1alpha3.VirtualService), err
}

func (s *virtualServiceClient) Get(name string, opts metav1.GetOptions) (*v1alpha3.VirtualService, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1alpha3.VirtualService), err
}

func (s *virtualServiceClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1alpha3.VirtualService, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1alpha3.VirtualService), err
}

func (s *virtualServiceClient) Update(o *v1alpha3.VirtualService) (*v1alpha3.VirtualService, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1alpha3.VirtualService), err
}

func (s *virtualServiceClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *virtualServiceClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *virtualServiceClient) List(opts metav1.ListOptions) (*VirtualServiceList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*VirtualServiceList), err
}

func (s *virtualServiceClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *virtualServiceClient) Patch(o *v1alpha3.VirtualService, patchType types.PatchType, data []byte, subresources ...string) (*v1alpha3.VirtualService, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1alpha3.VirtualService), err
}

func (s *virtualServiceClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *virtualServiceClient) AddHandler(ctx context.Context, name string, sync VirtualServiceHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *virtualServiceClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync VirtualServiceHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *virtualServiceClient) AddLifecycle(ctx context.Context, name string, lifecycle VirtualServiceLifecycle) {
	sync := NewVirtualServiceLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *virtualServiceClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle VirtualServiceLifecycle) {
	sync := NewVirtualServiceLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *virtualServiceClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync VirtualServiceHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *virtualServiceClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync VirtualServiceHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *virtualServiceClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle VirtualServiceLifecycle) {
	sync := NewVirtualServiceLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *virtualServiceClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle VirtualServiceLifecycle) {
	sync := NewVirtualServiceLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type VirtualServiceIndexer func(obj *v1alpha3.VirtualService) ([]string, error)

type VirtualServiceClientCache interface {
	Get(namespace, name string) (*v1alpha3.VirtualService, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha3.VirtualService, error)

	Index(name string, indexer VirtualServiceIndexer)
	GetIndexed(name, key string) ([]*v1alpha3.VirtualService, error)
}

type VirtualServiceClient interface {
	Create(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1alpha3.VirtualService, error)
	Update(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*VirtualServiceList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() VirtualServiceClientCache

	OnCreate(ctx context.Context, name string, sync VirtualServiceChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync VirtualServiceChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync VirtualServiceChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() VirtualServiceInterface
}

type virtualServiceClientCache struct {
	client *virtualServiceClient2
}

type virtualServiceClient2 struct {
	iface      VirtualServiceInterface
	controller VirtualServiceController
}

func (n *virtualServiceClient2) Interface() VirtualServiceInterface {
	return n.iface
}

func (n *virtualServiceClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *virtualServiceClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *virtualServiceClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *virtualServiceClient2) Create(obj *v1alpha3.VirtualService) (*v1alpha3.VirtualService, error) {
	return n.iface.Create(obj)
}

func (n *virtualServiceClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1alpha3.VirtualService, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *virtualServiceClient2) Update(obj *v1alpha3.VirtualService) (*v1alpha3.VirtualService, error) {
	return n.iface.Update(obj)
}

func (n *virtualServiceClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *virtualServiceClient2) List(namespace string, opts metav1.ListOptions) (*VirtualServiceList, error) {
	return n.iface.List(opts)
}

func (n *virtualServiceClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *virtualServiceClientCache) Get(namespace, name string) (*v1alpha3.VirtualService, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *virtualServiceClientCache) List(namespace string, selector labels.Selector) ([]*v1alpha3.VirtualService, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *virtualServiceClient2) Cache() VirtualServiceClientCache {
	n.loadController()
	return &virtualServiceClientCache{
		client: n,
	}
}

func (n *virtualServiceClient2) OnCreate(ctx context.Context, name string, sync VirtualServiceChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &virtualServiceLifecycleDelegate{create: sync})
}

func (n *virtualServiceClient2) OnChange(ctx context.Context, name string, sync VirtualServiceChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &virtualServiceLifecycleDelegate{update: sync})
}

func (n *virtualServiceClient2) OnRemove(ctx context.Context, name string, sync VirtualServiceChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &virtualServiceLifecycleDelegate{remove: sync})
}

func (n *virtualServiceClientCache) Index(name string, indexer VirtualServiceIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1alpha3.VirtualService); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *virtualServiceClientCache) GetIndexed(name, key string) ([]*v1alpha3.VirtualService, error) {
	var result []*v1alpha3.VirtualService
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1alpha3.VirtualService); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *virtualServiceClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type virtualServiceLifecycleDelegate struct {
	create VirtualServiceChangeHandlerFunc
	update VirtualServiceChangeHandlerFunc
	remove VirtualServiceChangeHandlerFunc
}

func (n *virtualServiceLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *virtualServiceLifecycleDelegate) Create(obj *v1alpha3.VirtualService) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *virtualServiceLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *virtualServiceLifecycleDelegate) Remove(obj *v1alpha3.VirtualService) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *virtualServiceLifecycleDelegate) Updated(obj *v1alpha3.VirtualService) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
