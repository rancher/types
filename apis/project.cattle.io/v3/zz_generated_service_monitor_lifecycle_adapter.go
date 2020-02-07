package v3

import (
	
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
)

type ServiceMonitorLifecycle interface {
	Create(obj *ServiceMonitor) (runtime.Object, error)
	Remove(obj *ServiceMonitor) (runtime.Object, error)
	Updated(obj *ServiceMonitor) (runtime.Object, error)
}

type serviceMonitorLifecycleAdapter struct {
	lifecycle ServiceMonitorLifecycle
}

func (w *serviceMonitorLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *serviceMonitorLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *serviceMonitorLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ServiceMonitor))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *serviceMonitorLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ServiceMonitor))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *serviceMonitorLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ServiceMonitor))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewServiceMonitorLifecycleAdapter(name string, clusterScoped bool, client ServiceMonitorInterface, l ServiceMonitorLifecycle) ServiceMonitorHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(ServiceMonitorGroupVersionResource)
	}
	adapter := &serviceMonitorLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ServiceMonitor) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
