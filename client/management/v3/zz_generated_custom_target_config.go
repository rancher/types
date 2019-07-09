package client

const (
	CustomTargetConfigType             = "customTargetConfig"
	CustomTargetConfigFieldCertificate = "certificate"
	CustomTargetConfigFieldClientCert  = "clientCert"
	CustomTargetConfigFieldClientKey   = "clientKey"
	CustomTargetConfigFieldContent     = "content"
)

type CustomTargetConfig struct {
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	ClientCert  string `json:"clientCert,omitempty" yaml:"client_cert,omitempty"`
	ClientKey   string `json:"clientKey,omitempty" yaml:"client_key,omitempty"`
	Content     string `json:"content,omitempty" yaml:"content,omitempty"`
}
