package client

const (
	NetworkVshpereOptsType               = "networkVshpereOpts"
	NetworkVshpereOptsFieldPublicNetwork = "public-network"
)

type NetworkVshpereOpts struct {
	PublicNetwork string `json:"public-network,omitempty" yaml:"public_network,omitempty"`
}
