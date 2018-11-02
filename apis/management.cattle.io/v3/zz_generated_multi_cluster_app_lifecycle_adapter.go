package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type MultiClusterAppLifecycle interface {
	Create(obj *MultiClusterApp) (*MultiClusterApp, error)
	Remove(obj *MultiClusterApp) (*MultiClusterApp, error)
	Updated(obj *MultiClusterApp) (*MultiClusterApp, error)
}

type multiClusterAppLifecycleAdapter struct {
	lifecycle MultiClusterAppLifecycle
}

func (w *multiClusterAppLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*MultiClusterApp))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *multiClusterAppLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*MultiClusterApp))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *multiClusterAppLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*MultiClusterApp))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewMultiClusterAppLifecycleAdapter(name string, clusterScoped bool, client MultiClusterAppInterface, l MultiClusterAppLifecycle) MultiClusterAppHandlerFunc {
	adapter := &multiClusterAppLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *MultiClusterApp) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
