package client

const (
	APIServerConfigType                     = "apiServerConfig"
	APIServerConfigFieldBearerToken         = "bearerToken"
	APIServerConfigFieldBearerTokenFile     = "bearerTokenFile"
	APIServerConfigFieldHost                = "host"
	APIServerConfigFieldPrometheusBasicAuth = "basicAuth"
	APIServerConfigFieldTLSConfig           = "tlsConfig"
)

type APIServerConfig struct {
	BearerToken         string               `json:"bearerToken,omitempty" yaml:"bearerToken,omitempty"`
	BearerTokenFile     string               `json:"bearerTokenFile,omitempty" yaml:"bearerTokenFile,omitempty"`
	Host                string               `json:"host,omitempty" yaml:"host,omitempty"`
	PrometheusBasicAuth *PrometheusBasicAuth `json:"basicAuth,omitempty" yaml:"basicAuth,omitempty"`
	TLSConfig           *TLSConfig           `json:"tlsConfig,omitempty" yaml:"tlsConfig,omitempty"`
}
