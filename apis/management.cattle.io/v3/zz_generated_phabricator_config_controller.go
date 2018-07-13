package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	PhabricatorConfigGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "PhabricatorConfig",
	}
	PhabricatorConfigResource = metav1.APIResource{
		Name:         "phabricatorconfigs",
		SingularName: "phabricatorconfig",
		Namespaced:   false,
		Kind:         PhabricatorConfigGroupVersionKind.Kind,
	}
)

type PhabricatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PhabricatorConfig
}

type PhabricatorConfigHandlerFunc func(key string, obj *PhabricatorConfig) error

type PhabricatorConfigLister interface {
	List(namespace string, selector labels.Selector) (ret []*PhabricatorConfig, err error)
	Get(namespace, name string) (*PhabricatorConfig, error)
}

type PhabricatorConfigController interface {
	Informer() cache.SharedIndexInformer
	Lister() PhabricatorConfigLister
	AddHandler(name string, handler PhabricatorConfigHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler PhabricatorConfigHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type PhabricatorConfigInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*PhabricatorConfig) (*PhabricatorConfig, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PhabricatorConfig, error)
	Get(name string, opts metav1.GetOptions) (*PhabricatorConfig, error)
	Update(*PhabricatorConfig) (*PhabricatorConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*PhabricatorConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() PhabricatorConfigController
	AddHandler(name string, sync PhabricatorConfigHandlerFunc)
	AddLifecycle(name string, lifecycle PhabricatorConfigLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync PhabricatorConfigHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle PhabricatorConfigLifecycle)
}

type phabricatorConfigLister struct {
	controller *phabricatorConfigController
}

func (l *phabricatorConfigLister) List(namespace string, selector labels.Selector) (ret []*PhabricatorConfig, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*PhabricatorConfig))
	})
	return
}

func (l *phabricatorConfigLister) Get(namespace, name string) (*PhabricatorConfig, error) {
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
			Group:    PhabricatorConfigGroupVersionKind.Group,
			Resource: "phabricatorConfig",
		}, key)
	}
	return obj.(*PhabricatorConfig), nil
}

type phabricatorConfigController struct {
	controller.GenericController
}

func (c *phabricatorConfigController) Lister() PhabricatorConfigLister {
	return &phabricatorConfigLister{
		controller: c,
	}
}

func (c *phabricatorConfigController) AddHandler(name string, handler PhabricatorConfigHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*PhabricatorConfig))
	})
}

func (c *phabricatorConfigController) AddClusterScopedHandler(name, cluster string, handler PhabricatorConfigHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*PhabricatorConfig))
	})
}

type phabricatorConfigFactory struct {
}

func (c phabricatorConfigFactory) Object() runtime.Object {
	return &PhabricatorConfig{}
}

func (c phabricatorConfigFactory) List() runtime.Object {
	return &PhabricatorConfigList{}
}

func (s *phabricatorConfigClient) Controller() PhabricatorConfigController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.phabricatorConfigControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(PhabricatorConfigGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &phabricatorConfigController{
		GenericController: genericController,
	}

	s.client.phabricatorConfigControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type phabricatorConfigClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   PhabricatorConfigController
}

func (s *phabricatorConfigClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *phabricatorConfigClient) Create(o *PhabricatorConfig) (*PhabricatorConfig, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*PhabricatorConfig), err
}

func (s *phabricatorConfigClient) Get(name string, opts metav1.GetOptions) (*PhabricatorConfig, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*PhabricatorConfig), err
}

func (s *phabricatorConfigClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PhabricatorConfig, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*PhabricatorConfig), err
}

func (s *phabricatorConfigClient) Update(o *PhabricatorConfig) (*PhabricatorConfig, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*PhabricatorConfig), err
}

func (s *phabricatorConfigClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *phabricatorConfigClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *phabricatorConfigClient) List(opts metav1.ListOptions) (*PhabricatorConfigList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*PhabricatorConfigList), err
}

func (s *phabricatorConfigClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *phabricatorConfigClient) Patch(o *PhabricatorConfig, data []byte, subresources ...string) (*PhabricatorConfig, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*PhabricatorConfig), err
}

func (s *phabricatorConfigClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *phabricatorConfigClient) AddHandler(name string, sync PhabricatorConfigHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *phabricatorConfigClient) AddLifecycle(name string, lifecycle PhabricatorConfigLifecycle) {
	sync := NewPhabricatorConfigLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *phabricatorConfigClient) AddClusterScopedHandler(name, clusterName string, sync PhabricatorConfigHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *phabricatorConfigClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle PhabricatorConfigLifecycle) {
	sync := NewPhabricatorConfigLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
