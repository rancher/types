package client

const (
	WebhookConfigType          = "webhookConfig"
	WebhookConfigFieldProxyURL = "proxyUrl"
	WebhookConfigFieldURL      = "url"
)

type WebhookConfig struct {
	ProxyURL string `json:"proxyUrl,omitempty" yaml:"proxy_url,omitempty"`
	URL      string `json:"url,omitempty" yaml:"url,omitempty"`
}
