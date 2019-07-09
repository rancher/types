package client

const (
	LifecycleType           = "lifecycle"
	LifecycleFieldPostStart = "postStart"
	LifecycleFieldPreStop   = "preStop"
)

type Lifecycle struct {
	PostStart *Handler `json:"postStart,omitempty" yaml:"post_start,omitempty"`
	PreStop   *Handler `json:"preStop,omitempty" yaml:"pre_stop,omitempty"`
}
