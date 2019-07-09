package client

const (
	AuthWebhookConfigType              = "authWebhookConfig"
	AuthWebhookConfigFieldCacheTimeout = "cacheTimeout"
	AuthWebhookConfigFieldConfigFile   = "configFile"
)

type AuthWebhookConfig struct {
	CacheTimeout string `json:"cacheTimeout,omitempty" yaml:"cache_timeout,omitempty"`
	ConfigFile   string `json:"configFile,omitempty" yaml:"config_file,omitempty"`
}
