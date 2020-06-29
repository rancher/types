package client

const (
	LoadBalancerClassOpenstackOptsType                   = "loadBalancerClassOpenstackOpts"
	LoadBalancerClassOpenstackOptsFieldFloatingNetworkID = "floating-network-id"
	LoadBalancerClassOpenstackOptsFieldFloatingSubnetID  = "floating-subnet-id"
	LoadBalancerClassOpenstackOptsFieldSubnetID          = "subnet-id"
)

type LoadBalancerClassOpenstackOpts struct {
	FloatingNetworkID string `json:"floating-network-id,omitempty" yaml:"floating-network-id,omitempty"`
	FloatingSubnetID  string `json:"floating-subnet-id,omitempty" yaml:"floating-subnet-id,omitempty"`
	SubnetID          string `json:"subnet-id,omitempty" yaml:"subnet-id,omitempty"`
}
