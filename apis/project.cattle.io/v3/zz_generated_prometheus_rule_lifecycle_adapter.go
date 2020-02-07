package v3

import (
	
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
)

type PrometheusRuleLifecycle interface {
	Create(obj *PrometheusRule) (runtime.Object, error)
	Remove(obj *PrometheusRule) (runtime.Object, error)
	Updated(obj *PrometheusRule) (runtime.Object, error)
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
	o, err := w.lifecycle.Create(obj.(*PrometheusRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *prometheusRuleLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*PrometheusRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *prometheusRuleLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*PrometheusRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewPrometheusRuleLifecycleAdapter(name string, clusterScoped bool, client PrometheusRuleInterface, l PrometheusRuleLifecycle) PrometheusRuleHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(PrometheusRuleGroupVersionResource)
	}
	adapter := &prometheusRuleLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *PrometheusRule) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
