package client

const (
	RemoteReadSpecType                     = "remoteReadSpec"
	RemoteReadSpecFieldBearerToken         = "bearerToken"
	RemoteReadSpecFieldBearerTokenFile     = "bearerTokenFile"
	RemoteReadSpecFieldPrometheusBasicAuth = "basicAuth"
	RemoteReadSpecFieldProxyURL            = "proxyUrl"
	RemoteReadSpecFieldReadRecent          = "readRecent"
	RemoteReadSpecFieldRemoteTimeout       = "remoteTimeout"
	RemoteReadSpecFieldRequiredMatchers    = "requiredMatchers"
	RemoteReadSpecFieldTLSConfig           = "tlsConfig"
	RemoteReadSpecFieldURL                 = "url"
)

type RemoteReadSpec struct {
	BearerToken         string               `json:"bearerToken,omitempty" yaml:"bearerToken,omitempty"`
	BearerTokenFile     string               `json:"bearerTokenFile,omitempty" yaml:"bearerTokenFile,omitempty"`
	PrometheusBasicAuth *PrometheusBasicAuth `json:"basicAuth,omitempty" yaml:"basicAuth,omitempty"`
	ProxyURL            string               `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
	ReadRecent          bool                 `json:"readRecent,omitempty" yaml:"readRecent,omitempty"`
	RemoteTimeout       string               `json:"remoteTimeout,omitempty" yaml:"remoteTimeout,omitempty"`
	RequiredMatchers    map[string]string    `json:"requiredMatchers,omitempty" yaml:"requiredMatchers,omitempty"`
	TLSConfig           *TLSConfig           `json:"tlsConfig,omitempty" yaml:"tlsConfig,omitempty"`
	URL                 string               `json:"url,omitempty" yaml:"url,omitempty"`
}
