package v3

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	UserGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "User",
	}
	UserResource = metav1.APIResource{
		Name:         "users",
		SingularName: "user",
		Namespaced:   false,
		Kind:         UserGroupVersionKind.Kind,
	}
)

type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User
}

type UserHandlerFunc func(key string, obj *User) error

type UserLister interface {
	List(namespace string, selector labels.Selector) (ret []*User, err error)
	Get(namespace, name string) (*User, error)
}

type UserController interface {
	Informer() cache.SharedIndexInformer
	Lister() UserLister
	AddHandler(handler UserHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type UserInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*User) (*User, error)
	GetNamespace(name, namespace string, opts metav1.GetOptions) (*User, error)
	Get(name string, opts metav1.GetOptions) (*User, error)
	Update(*User) (*User, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*UserList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() UserController
	AddSyncHandler(sync UserHandlerFunc)
	AddLifecycle(name string, lifecycle UserLifecycle)
}

type userLister struct {
	controller *userController
}

func (l *userLister) List(namespace string, selector labels.Selector) (ret []*User, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*User))
	})
	return
}

func (l *userLister) Get(namespace, name string) (*User, error) {
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
			Group:    UserGroupVersionKind.Group,
			Resource: "user",
		}, name)
	}
	return obj.(*User), nil
}

type userController struct {
	controller.GenericController
}

func (c *userController) Lister() UserLister {
	return &userLister{
		controller: c,
	}
}

func (c *userController) AddHandler(handler UserHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*User))
	})
}

type userFactory struct {
}

func (c userFactory) Object() runtime.Object {
	return &User{}
}

func (c userFactory) List() runtime.Object {
	return &UserList{}
}

func (s *userClient) Controller() UserController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.userControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(UserGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &userController{
		GenericController: genericController,
	}

	s.client.userControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type userClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   UserController
}

func (s *userClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *userClient) Create(o *User) (*User, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*User), err
}

func (s *userClient) Get(name string, opts metav1.GetOptions) (*User, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*User), err
}

func (s *userClient) GetNamespace(name, namespace string, opts metav1.GetOptions) (*User, error) {
	obj, err := s.objectClient.GetNamespace(name, namespace, opts)
	return obj.(*User), err
}

func (s *userClient) Update(o *User) (*User, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*User), err
}

func (s *userClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *userClient) DeleteNamespace(name, namespace string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespace(name, namespace, options)
}

func (s *userClient) List(opts metav1.ListOptions) (*UserList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*UserList), err
}

func (s *userClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *userClient) Patch(o *User, data []byte, subresources ...string) (*User, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*User), err
}

func (s *userClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *userClient) AddSyncHandler(sync UserHandlerFunc) {
	s.Controller().AddHandler(sync)
}

func (s *userClient) AddLifecycle(name string, lifecycle UserLifecycle) {
	sync := NewUserLifecycleAdapter(name, s, lifecycle)
	s.AddSyncHandler(sync)
}
