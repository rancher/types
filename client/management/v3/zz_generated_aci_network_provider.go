package client

const (
	AciNetworkProviderType              = "aciNetworkProvider"
	AciNetworkProviderFieldAciCNIConfig = "aci_cni_config"
)

type AciNetworkProvider struct {
	AciCNIConfig string `json:"aci_cni_config,omitempty" yaml:"aci_cni_config,omitempty"`
}
