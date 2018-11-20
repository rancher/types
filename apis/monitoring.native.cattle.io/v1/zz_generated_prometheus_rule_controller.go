package v1

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
	PrometheusRuleGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "PrometheusRule",
	}
	PrometheusRuleResource = metav1.APIResource{
		Name:         "prometheusrules",
		SingularName: "prometheusrule",
		Namespaced:   true,

		Kind: PrometheusRuleGroupVersionKind.Kind,
	}
)

type PrometheusRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrometheusRule
}

type PrometheusRuleHandlerFunc func(key string, obj *PrometheusRule) (runtime.Object, error)

type PrometheusRuleLister interface {
	List(namespace string, selector labels.Selector) (ret []*PrometheusRule, err error)
	Get(namespace, name string) (*PrometheusRule, error)
}

type PrometheusRuleController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() PrometheusRuleLister
	AddHandler(ctx context.Context, name string, handler PrometheusRuleHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler PrometheusRuleHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type PrometheusRuleInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*PrometheusRule) (*PrometheusRule, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PrometheusRule, error)
	Get(name string, opts metav1.GetOptions) (*PrometheusRule, error)
	Update(*PrometheusRule) (*PrometheusRule, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*PrometheusRuleList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() PrometheusRuleController
	AddHandler(ctx context.Context, name string, sync PrometheusRuleHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle PrometheusRuleLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PrometheusRuleHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PrometheusRuleLifecycle)
}

type prometheusRuleLister struct {
	controller *prometheusRuleController
}

func (l *prometheusRuleLister) List(namespace string, selector labels.Selector) (ret []*PrometheusRule, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*PrometheusRule))
	})
	return
}

func (l *prometheusRuleLister) Get(namespace, name string) (*PrometheusRule, error) {
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
			Group:    PrometheusRuleGroupVersionKind.Group,
			Resource: "prometheusRule",
		}, key)
	}
	return obj.(*PrometheusRule), nil
}

type prometheusRuleController struct {
	controller.GenericController
}

func (c *prometheusRuleController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *prometheusRuleController) Lister() PrometheusRuleLister {
	return &prometheusRuleLister{
		controller: c,
	}
}

func (c *prometheusRuleController) AddHandler(ctx context.Context, name string, handler PrometheusRuleHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*PrometheusRule); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *prometheusRuleController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler PrometheusRuleHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*PrometheusRule); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type prometheusRuleFactory struct {
}

func (c prometheusRuleFactory) Object() runtime.Object {
	return &PrometheusRule{}
}

func (c prometheusRuleFactory) List() runtime.Object {
	return &PrometheusRuleList{}
}

func (s *prometheusRuleClient) Controller() PrometheusRuleController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.prometheusRuleControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(PrometheusRuleGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &prometheusRuleController{
		GenericController: genericController,
	}

	s.client.prometheusRuleControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type prometheusRuleClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   PrometheusRuleController
}

func (s *prometheusRuleClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *prometheusRuleClient) Create(o *PrometheusRule) (*PrometheusRule, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*PrometheusRule), err
}

func (s *prometheusRuleClient) Get(name string, opts metav1.GetOptions) (*PrometheusRule, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*PrometheusRule), err
}

func (s *prometheusRuleClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PrometheusRule, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*PrometheusRule), err
}

func (s *prometheusRuleClient) Update(o *PrometheusRule) (*PrometheusRule, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*PrometheusRule), err
}

func (s *prometheusRuleClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *prometheusRuleClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *prometheusRuleClient) List(opts metav1.ListOptions) (*PrometheusRuleList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*PrometheusRuleList), err
}

func (s *prometheusRuleClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *prometheusRuleClient) Patch(o *PrometheusRule, data []byte, subresources ...string) (*PrometheusRule, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*PrometheusRule), err
}

func (s *prometheusRuleClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *prometheusRuleClient) AddHandler(ctx context.Context, name string, sync PrometheusRuleHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *prometheusRuleClient) AddLifecycle(ctx context.Context, name string, lifecycle PrometheusRuleLifecycle) {
	sync := NewPrometheusRuleLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *prometheusRuleClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PrometheusRuleHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *prometheusRuleClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PrometheusRuleLifecycle) {
	sync := NewPrometheusRuleLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}
