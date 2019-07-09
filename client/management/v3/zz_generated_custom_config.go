package client

const (
	CustomConfigType                 = "customConfig"
	CustomConfigFieldAddress         = "address"
	CustomConfigFieldDockerSocket    = "dockerSocket"
	CustomConfigFieldInternalAddress = "internalAddress"
	CustomConfigFieldLabel           = "label"
	CustomConfigFieldSSHCert         = "sshCert"
	CustomConfigFieldSSHKey          = "sshKey"
	CustomConfigFieldUser            = "user"
)

type CustomConfig struct {
	Address         string            `json:"address,omitempty" yaml:"address,omitempty"`
	DockerSocket    string            `json:"dockerSocket,omitempty" yaml:"docker_socket,omitempty"`
	InternalAddress string            `json:"internalAddress,omitempty" yaml:"internal_address,omitempty"`
	Label           map[string]string `json:"label,omitempty" yaml:"label,omitempty"`
	SSHCert         string            `json:"sshCert,omitempty" yaml:"ssh_cert,omitempty"`
	SSHKey          string            `json:"sshKey,omitempty" yaml:"ssh_key,omitempty"`
	User            string            `json:"user,omitempty" yaml:"user,omitempty"`
}
