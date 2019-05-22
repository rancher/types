package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ExampleConfigLifecycle interface {
	Create(obj *ExampleConfig) (runtime.Object, error)
	Remove(obj *ExampleConfig) (runtime.Object, error)
	Updated(obj *ExampleConfig) (runtime.Object, error)
}

type exampleConfigLifecycleAdapter struct {
	lifecycle ExampleConfigLifecycle
}

func (w *exampleConfigLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *exampleConfigLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *exampleConfigLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ExampleConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *exampleConfigLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ExampleConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *exampleConfigLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ExampleConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewExampleConfigLifecycleAdapter(name string, clusterScoped bool, client ExampleConfigInterface, l ExampleConfigLifecycle) ExampleConfigHandlerFunc {
	adapter := &exampleConfigLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ExampleConfig) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
