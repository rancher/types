package client

const (
	RKEConfigNodeType                  = "rkeConfigNode"
	RKEConfigNodeFieldAddress          = "address"
	RKEConfigNodeFieldDockerSocket     = "dockerSocket"
	RKEConfigNodeFieldHostnameOverride = "hostnameOverride"
	RKEConfigNodeFieldInternalAddress  = "internalAddress"
	RKEConfigNodeFieldLabels           = "labels"
	RKEConfigNodeFieldNodeID           = "nodeId"
	RKEConfigNodeFieldPort             = "port"
	RKEConfigNodeFieldRole             = "role"
	RKEConfigNodeFieldSSHAgentAuth     = "sshAgentAuth"
	RKEConfigNodeFieldSSHCert          = "sshCert"
	RKEConfigNodeFieldSSHCertPath      = "sshCertPath"
	RKEConfigNodeFieldSSHKey           = "sshKey"
	RKEConfigNodeFieldSSHKeyPath       = "sshKeyPath"
	RKEConfigNodeFieldUser             = "user"
)

type RKEConfigNode struct {
	Address          string            `json:"address,omitempty" yaml:"address,omitempty"`
	DockerSocket     string            `json:"dockerSocket,omitempty" yaml:"docker_socket,omitempty"`
	HostnameOverride string            `json:"hostnameOverride,omitempty" yaml:"hostname_override,omitempty"`
	InternalAddress  string            `json:"internalAddress,omitempty" yaml:"internal_address,omitempty"`
	Labels           map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	NodeID           string            `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	Port             string            `json:"port,omitempty" yaml:"port,omitempty"`
	Role             []string          `json:"role,omitempty" yaml:"role,omitempty"`
	SSHAgentAuth     bool              `json:"sshAgentAuth,omitempty" yaml:"ssh_agent_auth,omitempty"`
	SSHCert          string            `json:"sshCert,omitempty" yaml:"ssh_cert,omitempty"`
	SSHCertPath      string            `json:"sshCertPath,omitempty" yaml:"ssh_cert_path,omitempty"`
	SSHKey           string            `json:"sshKey,omitempty" yaml:"ssh_key,omitempty"`
	SSHKeyPath       string            `json:"sshKeyPath,omitempty" yaml:"ssh_key_path,omitempty"`
	User             string            `json:"user,omitempty" yaml:"user,omitempty"`
}
