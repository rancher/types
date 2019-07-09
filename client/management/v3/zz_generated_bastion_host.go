package client

const (
	BastionHostType              = "bastionHost"
	BastionHostFieldAddress      = "address"
	BastionHostFieldPort         = "port"
	BastionHostFieldSSHAgentAuth = "sshAgentAuth"
	BastionHostFieldSSHCert      = "sshCert"
	BastionHostFieldSSHCertPath  = "sshCertPath"
	BastionHostFieldSSHKey       = "sshKey"
	BastionHostFieldSSHKeyPath   = "sshKeyPath"
	BastionHostFieldUser         = "user"
)

type BastionHost struct {
	Address      string `json:"address,omitempty" yaml:"address,omitempty"`
	Port         string `json:"port,omitempty" yaml:"port,omitempty"`
	SSHAgentAuth bool   `json:"sshAgentAuth,omitempty" yaml:"ssh_agent_auth,omitempty"`
	SSHCert      string `json:"sshCert,omitempty" yaml:"ssh_cert,omitempty"`
	SSHCertPath  string `json:"sshCertPath,omitempty" yaml:"ssh_cert_path,omitempty"`
	SSHKey       string `json:"sshKey,omitempty" yaml:"ssh_key,omitempty"`
	SSHKeyPath   string `json:"sshKeyPath,omitempty" yaml:"ssh_key_path,omitempty"`
	User         string `json:"user,omitempty" yaml:"user,omitempty"`
}
