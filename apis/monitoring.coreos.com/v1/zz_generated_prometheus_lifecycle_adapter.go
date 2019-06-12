package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type PrometheusLifecycle interface {
	Create(obj *v1.Prometheus) (runtime.Object, error)
	Remove(obj *v1.Prometheus) (runtime.Object, error)
	Updated(obj *v1.Prometheus) (runtime.Object, error)
}

type prometheusLifecycleAdapter struct {
	lifecycle PrometheusLifecycle
}

func (w *prometheusLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *prometheusLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *prometheusLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.Prometheus))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *prometheusLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.Prometheus))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *prometheusLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.Prometheus))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewPrometheusLifecycleAdapter(name string, clusterScoped bool, client PrometheusInterface, l PrometheusLifecycle) PrometheusHandlerFunc {
	adapter := &prometheusLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.Prometheus) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
