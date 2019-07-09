package client

const (
	CapabilitiesType                          = "capabilities"
	CapabilitiesFieldIngressCapabilities      = "ingressCapabilities"
	CapabilitiesFieldLoadBalancerCapabilities = "loadBalancerCapabilities"
	CapabilitiesFieldNodePoolScalingSupported = "nodePoolScalingSupported"
	CapabilitiesFieldNodePortRange            = "nodePortRange"
)

type Capabilities struct {
	IngressCapabilities      []IngressCapabilities     `json:"ingressCapabilities,omitempty" yaml:"ingress_capabilities,omitempty"`
	LoadBalancerCapabilities *LoadBalancerCapabilities `json:"loadBalancerCapabilities,omitempty" yaml:"load_balancer_capabilities,omitempty"`
	NodePoolScalingSupported bool                      `json:"nodePoolScalingSupported,omitempty" yaml:"node_pool_scaling_supported,omitempty"`
	NodePortRange            string                    `json:"nodePortRange,omitempty" yaml:"node_port_range,omitempty"`
}
