package v3public

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ActiveDirectoryProviderLifecycle interface {
	Create(obj *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	Remove(obj *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	Updated(obj *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
}

type activeDirectoryProviderLifecycleAdapter struct {
	lifecycle ActiveDirectoryProviderLifecycle
}

func (w *activeDirectoryProviderLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ActiveDirectoryProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activeDirectoryProviderLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ActiveDirectoryProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activeDirectoryProviderLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ActiveDirectoryProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewActiveDirectoryProviderLifecycleAdapter(name string, clusterScoped bool, client ActiveDirectoryProviderInterface, l ActiveDirectoryProviderLifecycle) ActiveDirectoryProviderHandlerFunc {
	adapter := &activeDirectoryProviderLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ActiveDirectoryProvider) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
