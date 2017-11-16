package v1

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	TokenGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "authentication.cattle.io",
		Kind:    "Token",
	}
	TokenResource = metav1.APIResource{
		Name:         "tokens",
		SingularName: "token",
		Namespaced:   false,
		Kind:         TokenGroupVersionKind.Kind,
	}
)

type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token
}

type TokenHandlerFunc func(key string, obj *Token) error

type TokenController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler TokenHandlerFunc)
	Enqueue(namespace, name string)
	Start(threadiness int, ctx context.Context) error
}

type TokenInterface interface {
	Create(*Token) (*Token, error)
	Get(name string, opts metav1.GetOptions) (*Token, error)
	Update(*Token) (*Token, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*TokenList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() TokenController
}

type tokenController struct {
	controller.GenericController
}

func (c *tokenController) AddHandler(handler TokenHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*Token))
	})
}

type tokenFactory struct {
}

func (c tokenFactory) Object() runtime.Object {
	return &Token{}
}

func (c tokenFactory) List() runtime.Object {
	return &TokenList{}
}

func (s *tokenClient) Controller() TokenController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.tokenControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(TokenGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &tokenController{
		GenericController: genericController,
	}

	s.client.tokenControllers[s.ns] = c

	return c
}

type tokenClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   TokenController
}

func (s *tokenClient) Create(o *Token) (*Token, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*Token), err
}

func (s *tokenClient) Get(name string, opts metav1.GetOptions) (*Token, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*Token), err
}

func (s *tokenClient) Update(o *Token) (*Token, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*Token), err
}

func (s *tokenClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *tokenClient) List(opts metav1.ListOptions) (*TokenList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*TokenList), err
}

func (s *tokenClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *tokenClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
