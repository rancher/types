package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type PrometheusRuleLifecycle interface {
	Create(obj *v1.PrometheusRule) (runtime.Object, error)
	Remove(obj *v1.PrometheusRule) (runtime.Object, error)
	Updated(obj *v1.PrometheusRule) (runtime.Object, error)
}

type prometheusRuleLifecycleAdapter struct {
	lifecycle PrometheusRuleLifecycle
}

func (w *prometheusRuleLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *prometheusRuleLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *prometheusRuleLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.PrometheusRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *prometheusRuleLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.PrometheusRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *prometheusRuleLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.PrometheusRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewPrometheusRuleLifecycleAdapter(name string, clusterScoped bool, client PrometheusRuleInterface, l PrometheusRuleLifecycle) PrometheusRuleHandlerFunc {
	adapter := &prometheusRuleLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.PrometheusRule) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
