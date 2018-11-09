package client

const (
	GlobalDNSProviderSpecType                       = "globalDNSProviderSpec"
	GlobalDNSProviderSpecFieldProviderName          = "providerName"
	GlobalDNSProviderSpecFieldRoute53ProviderConfig = "route53ProviderConfig"
)

type GlobalDNSProviderSpec struct {
	ProviderName          string                 `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	Route53ProviderConfig *Route53ProviderConfig `json:"route53ProviderConfig,omitempty" yaml:"route53ProviderConfig,omitempty"`
}
