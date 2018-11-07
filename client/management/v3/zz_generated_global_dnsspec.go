package client

const (
	GlobalDNSSpecType              = "globalDNSSpec"
	GlobalDNSSpecFieldClusterIDs   = "clusterIds"
	GlobalDNSSpecFieldProviderName = "providerName"
	GlobalDNSSpecFieldRootDomain   = "rootDomain"
)

type GlobalDNSSpec struct {
	ClusterIDs   []string `json:"clusterIds,omitempty" yaml:"clusterIds,omitempty"`
	ProviderName string   `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	RootDomain   string   `json:"rootDomain,omitempty" yaml:"rootDomain,omitempty"`
}
