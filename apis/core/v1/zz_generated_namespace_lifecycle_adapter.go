package v1

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type NamespaceLifecycle interface {
	Create(obj *v1.Namespace) (*v1.Namespace, error)
	Remove(obj *v1.Namespace) (*v1.Namespace, error)
	Updated(obj *v1.Namespace) (*v1.Namespace, error)
}

type namespaceLifecycleAdapter struct {
	lifecycle NamespaceLifecycle
}

func (w *namespaceLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.Namespace))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *namespaceLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.Namespace))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *namespaceLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.Namespace))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewNamespaceLifecycleAdapter(name string, clusterScoped bool, client NamespaceInterface, l NamespaceLifecycle) NamespaceHandlerFunc {
	adapter := &namespaceLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.Namespace) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
