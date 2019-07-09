package client

const (
	CloudflareProviderConfigType              = "cloudflareProviderConfig"
	CloudflareProviderConfigFieldAPIEmail     = "apiEmail"
	CloudflareProviderConfigFieldAPIKey       = "apiKey"
	CloudflareProviderConfigFieldProxySetting = "proxySetting"
)

type CloudflareProviderConfig struct {
	APIEmail     string `json:"apiEmail,omitempty" yaml:"api_email,omitempty"`
	APIKey       string `json:"apiKey,omitempty" yaml:"api_key,omitempty"`
	ProxySetting *bool  `json:"proxySetting,omitempty" yaml:"proxy_setting,omitempty"`
}
