package v1

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type NodeLifecycle interface {
	Initialize(obj *v1.Node) error
	Remove(obj *v1.Node) error
	Updated(obj *v1.Node) error
}

type nodeLifecycleAdapter struct {
	lifecycle NodeLifecycle
}

func (w *nodeLifecycleAdapter) Initialize(obj runtime.Object) error {
	return w.lifecycle.Initialize(obj.(*v1.Node))
}

func (w *nodeLifecycleAdapter) Finalize(obj runtime.Object) error {
	return w.lifecycle.Remove(obj.(*v1.Node))
}

func (w *nodeLifecycleAdapter) Updated(obj runtime.Object) error {
	return w.lifecycle.Updated(obj.(*v1.Node))
}

func NewNodeLifecycleAdapter(name string, client NodeInterface, l NodeLifecycle) NodeHandlerFunc {
	adapter := &nodeLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, adapter, client.ObjectClient())
	return func(key string, obj *v1.Node) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
