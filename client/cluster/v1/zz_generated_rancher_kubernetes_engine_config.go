package client

const (
	RancherKubernetesEngineConfigType                = "rancherKubernetesEngineConfig"
	RancherKubernetesEngineConfigFieldAddons         = "addons"
	RancherKubernetesEngineConfigFieldAuthentication = "auth"
	RancherKubernetesEngineConfigFieldNetwork        = "network"
	RancherKubernetesEngineConfigFieldNodes          = "nodes"
	RancherKubernetesEngineConfigFieldSSHKeyPath     = "sshKeyPath"
	RancherKubernetesEngineConfigFieldServices       = "services"
)

type RancherKubernetesEngineConfig struct {
	Addons         string             `json:"addons,omitempty"`
	Authentication *AuthConfig        `json:"auth,omitempty"`
	Network        *NetworkConfig     `json:"network,omitempty"`
	Nodes          []RKEConfigNode    `json:"nodes,omitempty"`
	SSHKeyPath     string             `json:"sshKeyPath,omitempty"`
	Services       *RKEConfigServices `json:"services,omitempty"`
}
