package client

const (
	RancherKubernetesEngineConfigType                = "rancherKubernetesEngineConfig"
	RancherKubernetesEngineConfigFieldAddons         = "addons"
	RancherKubernetesEngineConfigFieldAuthentication = "auth"
	RancherKubernetesEngineConfigFieldHosts          = "hosts"
	RancherKubernetesEngineConfigFieldNetwork        = "network"
	RancherKubernetesEngineConfigFieldSSHKeyPath     = "sshKeyPath"
	RancherKubernetesEngineConfigFieldServices       = "services"
)

type RancherKubernetesEngineConfig struct {
	Addons         string             `json:"addons,omitempty"`
	Authentication *AuthConfig        `json:"auth,omitempty"`
	Hosts          []RKEConfigHost    `json:"hosts,omitempty"`
	Network        *NetworkConfig     `json:"network,omitempty"`
	SSHKeyPath     string             `json:"sshKeyPath,omitempty"`
	Services       *RKEConfigServices `json:"services,omitempty"`
}
