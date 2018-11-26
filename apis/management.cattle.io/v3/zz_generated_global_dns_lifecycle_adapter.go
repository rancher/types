package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type GlobalDNSLifecycle interface {
	Create(obj *GlobalDNS) (runtime.Object, error)
	Remove(obj *GlobalDNS) (runtime.Object, error)
	Updated(obj *GlobalDNS) (runtime.Object, error)
}

type globalDNSLifecycleAdapter struct {
	lifecycle GlobalDNSLifecycle
}

func (w *globalDNSLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *globalDNSLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *globalDNSLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*GlobalDNS))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalDNSLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*GlobalDNS))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalDNSLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*GlobalDNS))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewGlobalDNSLifecycleAdapter(name string, clusterScoped bool, client GlobalDNSInterface, l GlobalDNSLifecycle) GlobalDNSHandlerFunc {
	adapter := &globalDNSLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *GlobalDNS) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
