package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type PhabricatorConfigLifecycle interface {
	Create(obj *PhabricatorConfig) (*PhabricatorConfig, error)
	Remove(obj *PhabricatorConfig) (*PhabricatorConfig, error)
	Updated(obj *PhabricatorConfig) (*PhabricatorConfig, error)
}

type phabricatorConfigLifecycleAdapter struct {
	lifecycle PhabricatorConfigLifecycle
}

func (w *phabricatorConfigLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*PhabricatorConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *phabricatorConfigLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*PhabricatorConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *phabricatorConfigLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*PhabricatorConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewPhabricatorConfigLifecycleAdapter(name string, clusterScoped bool, client PhabricatorConfigInterface, l PhabricatorConfigLifecycle) PhabricatorConfigHandlerFunc {
	adapter := &phabricatorConfigLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *PhabricatorConfig) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
