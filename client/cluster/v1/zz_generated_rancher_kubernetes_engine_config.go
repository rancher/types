package client

const (
	RancherKubernetesEngineConfigType                = "rancherKubernetesEngineConfig"
	RancherKubernetesEngineConfigFieldAuthentication = "authentication"
	RancherKubernetesEngineConfigFieldHosts          = "hosts"
	RancherKubernetesEngineConfigFieldNetwork        = "network"
	RancherKubernetesEngineConfigFieldServices       = "services"
)

type RancherKubernetesEngineConfig struct {
	Authentication AuthConfig        `json:"authentication,omitempty"`
	Hosts          []RKEConfigHost   `json:"hosts,omitempty"`
	Network        NetworkConfig     `json:"network,omitempty"`
	Services       RKEConfigServices `json:"services,omitempty"`
}
