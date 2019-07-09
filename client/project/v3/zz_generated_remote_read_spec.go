package client

const (
	RemoteReadSpecType                  = "remoteReadSpec"
	RemoteReadSpecFieldBasicAuth        = "basicAuth"
	RemoteReadSpecFieldBearerToken      = "bearerToken"
	RemoteReadSpecFieldBearerTokenFile  = "bearerTokenFile"
	RemoteReadSpecFieldProxyURL         = "proxyUrl"
	RemoteReadSpecFieldReadRecent       = "readRecent"
	RemoteReadSpecFieldRemoteTimeout    = "remoteTimeout"
	RemoteReadSpecFieldRequiredMatchers = "requiredMatchers"
	RemoteReadSpecFieldTLSConfig        = "tlsConfig"
	RemoteReadSpecFieldURL              = "url"
)

type RemoteReadSpec struct {
	BasicAuth        *BasicAuth        `json:"basicAuth,omitempty" yaml:"basic_auth,omitempty"`
	BearerToken      string            `json:"bearerToken,omitempty" yaml:"bearer_token,omitempty"`
	BearerTokenFile  string            `json:"bearerTokenFile,omitempty" yaml:"bearer_token_file,omitempty"`
	ProxyURL         string            `json:"proxyUrl,omitempty" yaml:"proxy_url,omitempty"`
	ReadRecent       bool              `json:"readRecent,omitempty" yaml:"read_recent,omitempty"`
	RemoteTimeout    string            `json:"remoteTimeout,omitempty" yaml:"remote_timeout,omitempty"`
	RequiredMatchers map[string]string `json:"requiredMatchers,omitempty" yaml:"required_matchers,omitempty"`
	TLSConfig        *TLSConfig        `json:"tlsConfig,omitempty" yaml:"tls_config,omitempty"`
	URL              string            `json:"url,omitempty" yaml:"url,omitempty"`
}
