package client

const (
	KubeletServiceType                     = "kubeletService"
	KubeletServiceFieldClusterDNSServer    = "clusterDnsServer"
	KubeletServiceFieldClusterDomain       = "clusterDomain"
	KubeletServiceFieldExtraArgs           = "extraArgs"
	KubeletServiceFieldExtraBinds          = "extraBinds"
	KubeletServiceFieldExtraEnv            = "extraEnv"
	KubeletServiceFieldFailSwapOn          = "failSwapOn"
	KubeletServiceFieldImage               = "image"
	KubeletServiceFieldInfraContainerImage = "infraContainerImage"
)

type KubeletService struct {
	ClusterDNSServer    string            `json:"clusterDnsServer,omitempty" yaml:"cluster_dns_server,omitempty"`
	ClusterDomain       string            `json:"clusterDomain,omitempty" yaml:"cluster_domain,omitempty"`
	ExtraArgs           map[string]string `json:"extraArgs,omitempty" yaml:"extra_args,omitempty"`
	ExtraBinds          []string          `json:"extraBinds,omitempty" yaml:"extra_binds,omitempty"`
	ExtraEnv            []string          `json:"extraEnv,omitempty" yaml:"extra_env,omitempty"`
	FailSwapOn          bool              `json:"failSwapOn,omitempty" yaml:"fail_swap_on,omitempty"`
	Image               string            `json:"image,omitempty" yaml:"image,omitempty"`
	InfraContainerImage string            `json:"infraContainerImage,omitempty" yaml:"infra_container_image,omitempty"`
}
