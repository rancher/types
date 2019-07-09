package client

const (
	KubeControllerServiceType                       = "kubeControllerService"
	KubeControllerServiceFieldClusterCIDR           = "clusterCidr"
	KubeControllerServiceFieldExtraArgs             = "extraArgs"
	KubeControllerServiceFieldExtraBinds            = "extraBinds"
	KubeControllerServiceFieldExtraEnv              = "extraEnv"
	KubeControllerServiceFieldImage                 = "image"
	KubeControllerServiceFieldServiceClusterIPRange = "serviceClusterIpRange"
)

type KubeControllerService struct {
	ClusterCIDR           string            `json:"clusterCidr,omitempty" yaml:"cluster_cidr,omitempty"`
	ExtraArgs             map[string]string `json:"extraArgs,omitempty" yaml:"extra_args,omitempty"`
	ExtraBinds            []string          `json:"extraBinds,omitempty" yaml:"extra_binds,omitempty"`
	ExtraEnv              []string          `json:"extraEnv,omitempty" yaml:"extra_env,omitempty"`
	Image                 string            `json:"image,omitempty" yaml:"image,omitempty"`
	ServiceClusterIPRange string            `json:"serviceClusterIpRange,omitempty" yaml:"service_cluster_ip_range,omitempty"`
}
