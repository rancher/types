package client

const (
	FluentForwarderConfigType               = "fluentForwarderConfig"
	FluentForwarderConfigFieldCertificate   = "certificate"
	FluentForwarderConfigFieldCompress      = "compress"
	FluentForwarderConfigFieldEnableTLS     = "enableTls"
	FluentForwarderConfigFieldFluentServers = "fluentServers"
)

type FluentForwarderConfig struct {
	Certificate   string         `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	Compress      bool           `json:"compress,omitempty" yaml:"compress,omitempty"`
	EnableTLS     bool           `json:"enableTls,omitempty" yaml:"enable_tls,omitempty"`
	FluentServers []FluentServer `json:"fluentServers,omitempty" yaml:"fluent_servers,omitempty"`
}
