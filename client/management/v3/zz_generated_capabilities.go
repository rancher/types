package client

const (
	CapabilitiesType                          = "capabilities"
	CapabilitiesFieldIngressController        = "ingressController"
	CapabilitiesFieldL4LoadBalancer           = "l4loadBalancer"
	CapabilitiesFieldNodePoolScalingSupported = "nodePoolScalingSupported"
	CapabilitiesFieldNodePortRange            = "nodePortRange"
)

type Capabilities struct {
	IngressController        *IngressController `json:"ingressController,omitempty" yaml:"ingressController,omitempty"`
	L4LoadBalancer           *L4LoadBalancer    `json:"l4loadBalancer,omitempty" yaml:"l4loadBalancer,omitempty"`
	NodePoolScalingSupported bool               `json:"nodePoolScalingSupported,omitempty" yaml:"nodePoolScalingSupported,omitempty"`
	NodePortRange            string             `json:"nodePortRange,omitempty" yaml:"nodePortRange,omitempty"`
}
