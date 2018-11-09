package client

const (
	GlobalDNSSpecType                   = "globalDNSSpec"
	GlobalDNSSpecFieldFQDN              = "fqdn"
	GlobalDNSSpecFieldMultiClusterAppID = "multiClusterAppId"
	GlobalDNSSpecFieldProjectIDs        = "projectIds"
	GlobalDNSSpecFieldProviderName      = "providerName"
)

type GlobalDNSSpec struct {
	FQDN              string   `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	MultiClusterAppID string   `json:"multiClusterAppId,omitempty" yaml:"multiClusterAppId,omitempty"`
	ProjectIDs        []string `json:"projectIds,omitempty" yaml:"projectIds,omitempty"`
	ProviderName      string   `json:"providerName,omitempty" yaml:"providerName,omitempty"`
}
