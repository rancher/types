package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type GlobalDNSProviderLifecycle interface {
	Create(obj *GlobalDNSProvider) (runtime.Object, error)
	Remove(obj *GlobalDNSProvider) (runtime.Object, error)
	Updated(obj *GlobalDNSProvider) (runtime.Object, error)
}

type globalDNSProviderLifecycleAdapter struct {
	lifecycle GlobalDNSProviderLifecycle
}

func (w *globalDNSProviderLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *globalDNSProviderLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *globalDNSProviderLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*GlobalDNSProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalDNSProviderLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*GlobalDNSProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalDNSProviderLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*GlobalDNSProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewGlobalDNSProviderLifecycleAdapter(name string, clusterScoped bool, client GlobalDNSProviderInterface, l GlobalDNSProviderLifecycle) GlobalDNSProviderHandlerFunc {
	adapter := &globalDNSProviderLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *GlobalDNSProvider) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
