package client

const (
	RemoteWriteSpecType                     = "remoteWriteSpec"
	RemoteWriteSpecFieldBasicAuth           = "basicAuth"
	RemoteWriteSpecFieldBearerToken         = "bearerToken"
	RemoteWriteSpecFieldBearerTokenFile     = "bearerTokenFile"
	RemoteWriteSpecFieldProxyURL            = "proxyUrl"
	RemoteWriteSpecFieldQueueConfig         = "queueConfig"
	RemoteWriteSpecFieldRemoteTimeout       = "remoteTimeout"
	RemoteWriteSpecFieldTLSConfig           = "tlsConfig"
	RemoteWriteSpecFieldURL                 = "url"
	RemoteWriteSpecFieldWriteRelabelConfigs = "writeRelabelConfigs"
)

type RemoteWriteSpec struct {
	BasicAuth           *BasicAuth      `json:"basicAuth,omitempty" yaml:"basic_auth,omitempty"`
	BearerToken         string          `json:"bearerToken,omitempty" yaml:"bearer_token,omitempty"`
	BearerTokenFile     string          `json:"bearerTokenFile,omitempty" yaml:"bearer_token_file,omitempty"`
	ProxyURL            string          `json:"proxyUrl,omitempty" yaml:"proxy_url,omitempty"`
	QueueConfig         *QueueConfig    `json:"queueConfig,omitempty" yaml:"queue_config,omitempty"`
	RemoteTimeout       string          `json:"remoteTimeout,omitempty" yaml:"remote_timeout,omitempty"`
	TLSConfig           *TLSConfig      `json:"tlsConfig,omitempty" yaml:"tls_config,omitempty"`
	URL                 string          `json:"url,omitempty" yaml:"url,omitempty"`
	WriteRelabelConfigs []RelabelConfig `json:"writeRelabelConfigs,omitempty" yaml:"write_relabel_configs,omitempty"`
}
