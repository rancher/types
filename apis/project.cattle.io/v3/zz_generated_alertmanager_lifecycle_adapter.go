package v3

import (
	
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
)

type AlertmanagerLifecycle interface {
	Create(obj *Alertmanager) (runtime.Object, error)
	Remove(obj *Alertmanager) (runtime.Object, error)
	Updated(obj *Alertmanager) (runtime.Object, error)
}

type alertmanagerLifecycleAdapter struct {
	lifecycle AlertmanagerLifecycle
}

func (w *alertmanagerLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *alertmanagerLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *alertmanagerLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*Alertmanager))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *alertmanagerLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*Alertmanager))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *alertmanagerLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*Alertmanager))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewAlertmanagerLifecycleAdapter(name string, clusterScoped bool, client AlertmanagerInterface, l AlertmanagerLifecycle) AlertmanagerHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(AlertmanagerGroupVersionResource)
	}
	adapter := &alertmanagerLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *Alertmanager) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
