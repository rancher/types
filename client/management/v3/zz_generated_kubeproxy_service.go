package client

const (
	KubeproxyServiceType            = "kubeproxyService"
	KubeproxyServiceFieldExtraArgs  = "extraArgs"
	KubeproxyServiceFieldExtraBinds = "extraBinds"
	KubeproxyServiceFieldExtraEnv   = "extraEnv"
	KubeproxyServiceFieldImage      = "image"
)

type KubeproxyService struct {
	ExtraArgs  map[string]string `json:"extraArgs,omitempty" yaml:"extra_args,omitempty"`
	ExtraBinds []string          `json:"extraBinds,omitempty" yaml:"extra_binds,omitempty"`
	ExtraEnv   []string          `json:"extraEnv,omitempty" yaml:"extra_env,omitempty"`
	Image      string            `json:"image,omitempty" yaml:"image,omitempty"`
}
