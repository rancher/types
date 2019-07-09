package client

const (
	KubeAPIServiceType                       = "kubeAPIService"
	KubeAPIServiceFieldAlwaysPullImages      = "alwaysPullImages"
	KubeAPIServiceFieldExtraArgs             = "extraArgs"
	KubeAPIServiceFieldExtraBinds            = "extraBinds"
	KubeAPIServiceFieldExtraEnv              = "extraEnv"
	KubeAPIServiceFieldImage                 = "image"
	KubeAPIServiceFieldPodSecurityPolicy     = "podSecurityPolicy"
	KubeAPIServiceFieldServiceClusterIPRange = "serviceClusterIpRange"
	KubeAPIServiceFieldServiceNodePortRange  = "serviceNodePortRange"
)

type KubeAPIService struct {
	AlwaysPullImages      bool              `json:"alwaysPullImages,omitempty" yaml:"always_pull_images,omitempty"`
	ExtraArgs             map[string]string `json:"extraArgs,omitempty" yaml:"extra_args,omitempty"`
	ExtraBinds            []string          `json:"extraBinds,omitempty" yaml:"extra_binds,omitempty"`
	ExtraEnv              []string          `json:"extraEnv,omitempty" yaml:"extra_env,omitempty"`
	Image                 string            `json:"image,omitempty" yaml:"image,omitempty"`
	PodSecurityPolicy     bool              `json:"podSecurityPolicy,omitempty" yaml:"pod_security_policy,omitempty"`
	ServiceClusterIPRange string            `json:"serviceClusterIpRange,omitempty" yaml:"service_cluster_ip_range,omitempty"`
	ServiceNodePortRange  string            `json:"serviceNodePortRange,omitempty" yaml:"service_node_port_range,omitempty"`
}
