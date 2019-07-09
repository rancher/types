package client

const (
	PagerdutyConfigType            = "pagerdutyConfig"
	PagerdutyConfigFieldProxyURL   = "proxyUrl"
	PagerdutyConfigFieldServiceKey = "serviceKey"
)

type PagerdutyConfig struct {
	ProxyURL   string `json:"proxyUrl,omitempty" yaml:"proxy_url,omitempty"`
	ServiceKey string `json:"serviceKey,omitempty" yaml:"service_key,omitempty"`
}
