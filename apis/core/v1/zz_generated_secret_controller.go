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
	SecretGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Secret",
	}
	SecretResource = metav1.APIResource{
		Name:         "secrets",
		SingularName: "secret",
		Namespaced:   true,

		Kind: SecretGroupVersionKind.Kind,
	}

	SecretGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "secrets",
	}
)

func init() {
	resource.Put(SecretGroupVersionResource)
}

func NewSecret(namespace, name string, obj v1.Secret) *v1.Secret {
	obj.APIVersion, obj.Kind = SecretGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.Secret `json:"items"`
}

type SecretHandlerFunc func(key string, obj *v1.Secret) (runtime.Object, error)

type SecretChangeHandlerFunc func(obj *v1.Secret) (runtime.Object, error)

type SecretLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.Secret, err error)
	Get(namespace, name string) (*v1.Secret, error)
}

type SecretController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() SecretLister
	AddHandler(ctx context.Context, name string, handler SecretHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync SecretHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler SecretHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type SecretInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.Secret) (*v1.Secret, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.Secret, error)
	Get(name string, opts metav1.GetOptions) (*v1.Secret, error)
	Update(*v1.Secret) (*v1.Secret, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*SecretList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() SecretController
	AddHandler(ctx context.Context, name string, sync SecretHandlerFunc)
	AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync SecretHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle SecretLifecycle)
	AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle SecretLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync SecretHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle SecretLifecycle)
}

type secretLister struct {
	controller *secretController
}

func (l *secretLister) List(namespace string, selector labels.Selector) (ret []*v1.Secret, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.Secret))
	})
	return
}

func (l *secretLister) Get(namespace, name string) (*v1.Secret, error) {
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
			Group:    SecretGroupVersionKind.Group,
			Resource: "secret",
		}, key)
	}
	return obj.(*v1.Secret), nil
}

type secretController struct {
	controller.GenericController
}

func (c *secretController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *secretController) Lister() SecretLister {
	return &secretLister{
		controller: c,
	}
}

func (c *secretController) AddHandler(ctx context.Context, name string, handler SecretHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Secret); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *secretController) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, handler SecretHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled(feat) {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Secret); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *secretController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler SecretHandlerFunc) {
	resource.PutClusterScoped(SecretGroupVersionResource)
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.Secret); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type secretFactory struct {
}

func (c secretFactory) Object() runtime.Object {
	return &v1.Secret{}
}

func (c secretFactory) List() runtime.Object {
	return &SecretList{}
}

func (s *secretClient) Controller() SecretController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.secretControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(SecretGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &secretController{
		GenericController: genericController,
	}

	s.client.secretControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type secretClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   SecretController
}

func (s *secretClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *secretClient) Create(o *v1.Secret) (*v1.Secret, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.Secret), err
}

func (s *secretClient) Get(name string, opts metav1.GetOptions) (*v1.Secret, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.Secret), err
}

func (s *secretClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.Secret, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.Secret), err
}

func (s *secretClient) Update(o *v1.Secret) (*v1.Secret, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.Secret), err
}

func (s *secretClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *secretClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *secretClient) List(opts metav1.ListOptions) (*SecretList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*SecretList), err
}

func (s *secretClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *secretClient) Patch(o *v1.Secret, patchType types.PatchType, data []byte, subresources ...string) (*v1.Secret, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.Secret), err
}

func (s *secretClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *secretClient) AddHandler(ctx context.Context, name string, sync SecretHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *secretClient) AddFeatureHandler(enabled func(string) bool, feat string, ctx context.Context, name string, sync SecretHandlerFunc) {
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *secretClient) AddLifecycle(ctx context.Context, name string, lifecycle SecretLifecycle) {
	sync := NewSecretLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *secretClient) AddFeatureLifecycle(enabled func(string) bool, feat string, ctx context.Context, name string, lifecycle SecretLifecycle) {
	sync := NewSecretLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(enabled, feat, ctx, name, sync)
}

func (s *secretClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync SecretHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *secretClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle SecretLifecycle) {
	sync := NewSecretLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type SecretIndexer func(obj *v1.Secret) ([]string, error)

type SecretClientCache interface {
	Get(namespace, name string) (*v1.Secret, error)
	List(namespace string, selector labels.Selector) ([]*v1.Secret, error)

	Index(name string, indexer SecretIndexer)
	GetIndexed(name, key string) ([]*v1.Secret, error)
}

type SecretClient interface {
	Create(*v1.Secret) (*v1.Secret, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1.Secret, error)
	Update(*v1.Secret) (*v1.Secret, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*SecretList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() SecretClientCache

	OnCreate(ctx context.Context, name string, sync SecretChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync SecretChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync SecretChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() SecretInterface
}

type secretClientCache struct {
	client *secretClient2
}

type secretClient2 struct {
	iface      SecretInterface
	controller SecretController
}

func (n *secretClient2) Interface() SecretInterface {
	return n.iface
}

func (n *secretClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *secretClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *secretClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *secretClient2) Create(obj *v1.Secret) (*v1.Secret, error) {
	return n.iface.Create(obj)
}

func (n *secretClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1.Secret, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *secretClient2) Update(obj *v1.Secret) (*v1.Secret, error) {
	return n.iface.Update(obj)
}

func (n *secretClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *secretClient2) List(namespace string, opts metav1.ListOptions) (*SecretList, error) {
	return n.iface.List(opts)
}

func (n *secretClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *secretClientCache) Get(namespace, name string) (*v1.Secret, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *secretClientCache) List(namespace string, selector labels.Selector) ([]*v1.Secret, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *secretClient2) Cache() SecretClientCache {
	n.loadController()
	return &secretClientCache{
		client: n,
	}
}

func (n *secretClient2) OnCreate(ctx context.Context, name string, sync SecretChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &secretLifecycleDelegate{create: sync})
}

func (n *secretClient2) OnChange(ctx context.Context, name string, sync SecretChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &secretLifecycleDelegate{update: sync})
}

func (n *secretClient2) OnRemove(ctx context.Context, name string, sync SecretChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &secretLifecycleDelegate{remove: sync})
}

func (n *secretClientCache) Index(name string, indexer SecretIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1.Secret); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *secretClientCache) GetIndexed(name, key string) ([]*v1.Secret, error) {
	var result []*v1.Secret
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1.Secret); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *secretClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type secretLifecycleDelegate struct {
	create SecretChangeHandlerFunc
	update SecretChangeHandlerFunc
	remove SecretChangeHandlerFunc
}

func (n *secretLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *secretLifecycleDelegate) Create(obj *v1.Secret) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *secretLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *secretLifecycleDelegate) Remove(obj *v1.Secret) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *secretLifecycleDelegate) Updated(obj *v1.Secret) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
