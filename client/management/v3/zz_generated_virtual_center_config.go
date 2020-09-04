package client

const (
	VirtualCenterConfigType                   = "virtualCenterConfig"
	VirtualCenterConfigFieldDatacenters       = "datacenters"
	VirtualCenterConfigFieldPassword          = "password"
	VirtualCenterConfigFieldRoundTripperCount = "soap-roundtrip-count"
	VirtualCenterConfigFieldSecretName        = "secret-name"
	VirtualCenterConfigFieldSecretNamespace   = "secret-namespace"
	VirtualCenterConfigFieldUser              = "user"
	VirtualCenterConfigFieldVCenterPort       = "port"
)

type VirtualCenterConfig struct {
	Datacenters       string `json:"datacenters,omitempty" yaml:"datacenters,omitempty"`
	Password          string `json:"password,omitempty" yaml:"password,omitempty"`
	RoundTripperCount int64  `json:"soap-roundtrip-count,omitempty" yaml:"soap-roundtrip-count,omitempty"`
	SecretName        string `json:"secret-name,omitempty" yaml:"secret-name,omitempty"`
	SecretNamespace   string `json:"secret-namespace,omitempty" yaml:"secret-namespace,omitempty"`
	User              string `json:"user,omitempty" yaml:"user,omitempty"`
	VCenterPort       string `json:"port,omitempty" yaml:"port,omitempty"`
}
