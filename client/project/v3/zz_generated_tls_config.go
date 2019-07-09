package client

const (
	TLSConfigType                    = "tlsConfig"
	TLSConfigFieldCAFile             = "caFile"
	TLSConfigFieldCertFile           = "certFile"
	TLSConfigFieldInsecureSkipVerify = "insecureSkipVerify"
	TLSConfigFieldKeyFile            = "keyFile"
	TLSConfigFieldServerName         = "serverName"
)

type TLSConfig struct {
	CAFile             string `json:"caFile,omitempty" yaml:"ca_file,omitempty"`
	CertFile           string `json:"certFile,omitempty" yaml:"cert_file,omitempty"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify,omitempty" yaml:"insecure_skip_verify,omitempty"`
	KeyFile            string `json:"keyFile,omitempty" yaml:"key_file,omitempty"`
	ServerName         string `json:"serverName,omitempty" yaml:"server_name,omitempty"`
}
