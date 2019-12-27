package client

const (
	NetworkConfigType                        = "networkConfig"
	NetworkConfigFieldAciNetworkProvider     = "aci_network_provider"
	NetworkConfigFieldCalicoNetworkProvider  = "calicoNetworkProvider"
	NetworkConfigFieldCanalNetworkProvider   = "canalNetworkProvider"
	NetworkConfigFieldFlannelNetworkProvider = "flannelNetworkProvider"
	NetworkConfigFieldMTU                    = "mtu"
	NetworkConfigFieldNodeSelector           = "nodeSelector"
	NetworkConfigFieldOptions                = "options"
	NetworkConfigFieldPlugin                 = "plugin"
	NetworkConfigFieldWeaveNetworkProvider   = "weaveNetworkProvider"
)

type NetworkConfig struct {
	AciNetworkProvider     *AciNetworkProvider     `json:"aci_network_provider,omitempty" yaml:"aci_network_provider,omitempty"`
	CalicoNetworkProvider  *CalicoNetworkProvider  `json:"calicoNetworkProvider,omitempty" yaml:"calicoNetworkProvider,omitempty"`
	CanalNetworkProvider   *CanalNetworkProvider   `json:"canalNetworkProvider,omitempty" yaml:"canalNetworkProvider,omitempty"`
	FlannelNetworkProvider *FlannelNetworkProvider `json:"flannelNetworkProvider,omitempty" yaml:"flannelNetworkProvider,omitempty"`
	MTU                    int64                   `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	NodeSelector           map[string]string       `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	Options                map[string]string       `json:"options,omitempty" yaml:"options,omitempty"`
	Plugin                 string                  `json:"plugin,omitempty" yaml:"plugin,omitempty"`
	WeaveNetworkProvider   *WeaveNetworkProvider   `json:"weaveNetworkProvider,omitempty" yaml:"weaveNetworkProvider,omitempty"`
}
