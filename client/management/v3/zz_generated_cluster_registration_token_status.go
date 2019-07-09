package client

const (
	ClusterRegistrationTokenStatusType                    = "clusterRegistrationTokenStatus"
	ClusterRegistrationTokenStatusFieldCommand            = "command"
	ClusterRegistrationTokenStatusFieldInsecureCommand    = "insecureCommand"
	ClusterRegistrationTokenStatusFieldManifestURL        = "manifestUrl"
	ClusterRegistrationTokenStatusFieldNodeCommand        = "nodeCommand"
	ClusterRegistrationTokenStatusFieldToken              = "token"
	ClusterRegistrationTokenStatusFieldWindowsNodeCommand = "windowsNodeCommand"
)

type ClusterRegistrationTokenStatus struct {
	Command            string `json:"command,omitempty" yaml:"command,omitempty"`
	InsecureCommand    string `json:"insecureCommand,omitempty" yaml:"insecure_command,omitempty"`
	ManifestURL        string `json:"manifestUrl,omitempty" yaml:"manifest_url,omitempty"`
	NodeCommand        string `json:"nodeCommand,omitempty" yaml:"node_command,omitempty"`
	Token              string `json:"token,omitempty" yaml:"token,omitempty"`
	WindowsNodeCommand string `json:"windowsNodeCommand,omitempty" yaml:"windows_node_command,omitempty"`
}
