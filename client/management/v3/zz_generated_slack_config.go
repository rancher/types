package client

const (
	SlackConfigType                  = "slackConfig"
	SlackConfigFieldDefaultRecipient = "defaultRecipient"
	SlackConfigFieldProxyURL         = "proxyUrl"
	SlackConfigFieldURL              = "url"
)

type SlackConfig struct {
	DefaultRecipient string `json:"defaultRecipient,omitempty" yaml:"default_recipient,omitempty"`
	ProxyURL         string `json:"proxyUrl,omitempty" yaml:"proxy_url,omitempty"`
	URL              string `json:"url,omitempty" yaml:"url,omitempty"`
}
