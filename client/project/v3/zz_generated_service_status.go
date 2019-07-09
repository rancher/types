package client

const (
	ServiceStatusType              = "serviceStatus"
	ServiceStatusFieldLoadBalancer = "loadBalancer"
)

type ServiceStatus struct {
	LoadBalancer *LoadBalancerStatus `json:"loadBalancer,omitempty" yaml:"load_balancer,omitempty"`
}
