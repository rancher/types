package client

const (
	NetworkConfigType                        = "networkConfig"
	NetworkConfigFieldCalicoNetworkProvider  = "calicoNetworkProvider"
	NetworkConfigFieldCanalNetworkProvider   = "canalNetworkProvider"
	NetworkConfigFieldFlannelNetworkProvider = "flannelNetworkProvider"
	NetworkConfigFieldOptions                = "options"
	NetworkConfigFieldPlugin                 = "plugin"
	NetworkConfigFieldWeaveNetworkProvider   = "weaveNetworkProvider"
)

type NetworkConfig struct {
	CalicoNetworkProvider  *CalicoNetworkProvider  `json:"calicoNetworkProvider,omitempty" yaml:"calico_network_provider,omitempty"`
	CanalNetworkProvider   *CanalNetworkProvider   `json:"canalNetworkProvider,omitempty" yaml:"canal_network_provider,omitempty"`
	FlannelNetworkProvider *FlannelNetworkProvider `json:"flannelNetworkProvider,omitempty" yaml:"flannel_network_provider,omitempty"`
	Options                map[string]string       `json:"options,omitempty" yaml:"options,omitempty"`
	Plugin                 string                  `json:"plugin,omitempty" yaml:"plugin,omitempty"`
	WeaveNetworkProvider   *WeaveNetworkProvider   `json:"weaveNetworkProvider,omitempty" yaml:"weave_network_provider,omitempty"`
}
