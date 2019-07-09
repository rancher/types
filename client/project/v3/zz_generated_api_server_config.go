package client

const (
	APIServerConfigType                 = "apiServerConfig"
	APIServerConfigFieldBasicAuth       = "basicAuth"
	APIServerConfigFieldBearerToken     = "bearerToken"
	APIServerConfigFieldBearerTokenFile = "bearerTokenFile"
	APIServerConfigFieldHost            = "host"
	APIServerConfigFieldTLSConfig       = "tlsConfig"
)

type APIServerConfig struct {
	BasicAuth       *BasicAuth `json:"basicAuth,omitempty" yaml:"basic_auth,omitempty"`
	BearerToken     string     `json:"bearerToken,omitempty" yaml:"bearer_token,omitempty"`
	BearerTokenFile string     `json:"bearerTokenFile,omitempty" yaml:"bearer_token_file,omitempty"`
	Host            string     `json:"host,omitempty" yaml:"host,omitempty"`
	TLSConfig       *TLSConfig `json:"tlsConfig,omitempty" yaml:"tls_config,omitempty"`
}
