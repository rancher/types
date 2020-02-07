package v3

import (
	"context"
	"time"

	
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/controller"
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

		Kind:         VirtualServiceGroupVersionKind.Kind,
	}

	VirtualServiceGroupVersionResource = schema.GroupVersionResource{
		Group:     GroupName,
		Version:   Version,
		Resource:  "virtualservices",
	}
)

func init() {
	resource.Put(VirtualServiceGroupVersionResource)
}

func NewVirtualService(namespace, name string, obj VirtualService) *VirtualService {
	obj.APIVersion, obj.Kind = VirtualServiceGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualServiceList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta   `json:"metadata,omitempty"`
	Items             []VirtualService `json:"items"`
}

type VirtualServiceHandlerFunc func(key string, obj *VirtualService) (runtime.Object, error)

type VirtualServiceChangeHandlerFunc func(obj *VirtualService) (runtime.Object, error)

type VirtualServiceLister interface {
	List(namespace string, selector labels.Selector) (ret []*VirtualService, err error)
	Get(namespace, name string) (*VirtualService, error)
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
	EnqueueAfter(namespace, name string, after time.Duration)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type VirtualServiceInterface interface {
    ObjectClient() *objectclient.ObjectClient
	Create(*VirtualService) (*VirtualService, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*VirtualService, error)
	Get(name string, opts metav1.GetOptions) (*VirtualService, error)
	Update(*VirtualService) (*VirtualService, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*VirtualServiceList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*VirtualServiceList, error)
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

func (l *virtualServiceLister) List(namespace string, selector labels.Selector) (ret []*VirtualService, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*VirtualService))
	})
	return
}

func (l *virtualServiceLister) Get(namespace, name string) (*VirtualService, error) {
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
			Group: VirtualServiceGroupVersionKind.Group,
			Resource: "virtualService",
		}, key)
	}
	return obj.(*VirtualService), nil
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
		} else if v, ok := obj.(*VirtualService); ok {
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
		} else if v, ok := obj.(*VirtualService); ok {
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
		} else if v, ok := obj.(*VirtualService); ok && controller.ObjectInCluster(cluster, obj) {
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
		} else if v, ok := obj.(*VirtualService); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type virtualServiceFactory struct {
}

func (c virtualServiceFactory) Object() runtime.Object {
	return &VirtualService{}
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
	client *Client
	ns string
	objectClient *objectclient.ObjectClient
	controller   VirtualServiceController
}

func (s *virtualServiceClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *virtualServiceClient) Create(o *VirtualService) (*VirtualService, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*VirtualService), err
}

func (s *virtualServiceClient) Get(name string, opts metav1.GetOptions) (*VirtualService, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*VirtualService), err
}

func (s *virtualServiceClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*VirtualService, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*VirtualService), err
}

func (s *virtualServiceClient) Update(o *VirtualService) (*VirtualService, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*VirtualService), err
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

func (s *virtualServiceClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*VirtualServiceList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*VirtualServiceList), err
}

func (s *virtualServiceClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *virtualServiceClient) Patch(o *VirtualService, patchType types.PatchType, data []byte, subresources ...string) (*VirtualService, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*VirtualService), err
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
