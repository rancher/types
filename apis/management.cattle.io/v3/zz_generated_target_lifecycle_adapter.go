package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type TargetLifecycle interface {
	Create(obj *Target) (*Target, error)
	Remove(obj *Target) (*Target, error)
	Updated(obj *Target) (*Target, error)
}

type targetLifecycleAdapter struct {
	lifecycle TargetLifecycle
}

func (w *targetLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*Target))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *targetLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*Target))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *targetLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*Target))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewTargetLifecycleAdapter(name string, clusterScoped bool, client TargetInterface, l TargetLifecycle) TargetHandlerFunc {
	adapter := &targetLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *Target) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
