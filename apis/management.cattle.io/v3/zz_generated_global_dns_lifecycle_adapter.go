package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type GlobalDNSLifecycle interface {
	Create(obj *GlobalDNS) (*GlobalDNS, error)
	Remove(obj *GlobalDNS) (*GlobalDNS, error)
	Updated(obj *GlobalDNS) (*GlobalDNS, error)
}

type globalDNSLifecycleAdapter struct {
	lifecycle GlobalDNSLifecycle
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
	return func(key string, obj *GlobalDNS) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
