package client

const (
	ServiceSpecType                          = "serviceSpec"
	ServiceSpecFieldClusterIp                = "clusterIp"
	ServiceSpecFieldExternalIPs              = "externalIPs"
	ServiceSpecFieldExternalTrafficPolicy    = "externalTrafficPolicy"
	ServiceSpecFieldHealthCheckNodePort      = "healthCheckNodePort"
	ServiceSpecFieldHostname                 = "hostname"
	ServiceSpecFieldLoadBalancerIP           = "loadBalancerIP"
	ServiceSpecFieldLoadBalancerSourceRanges = "loadBalancerSourceRanges"
	ServiceSpecFieldPorts                    = "ports"
	ServiceSpecFieldPublishNotReadyAddresses = "publishNotReadyAddresses"
	ServiceSpecFieldSelector                 = "selector"
	ServiceSpecFieldServiceKind              = "serviceKind"
	ServiceSpecFieldSessionAffinity          = "sessionAffinity"
	ServiceSpecFieldSessionAffinityConfig    = "sessionAffinityConfig"
)

type ServiceSpec struct {
	ClusterIp                string                 `json:"clusterIp,omitempty" yaml:"cluster_ip,omitempty"`
	ExternalIPs              []string               `json:"externalIPs,omitempty" yaml:"external_i_ps,omitempty"`
	ExternalTrafficPolicy    string                 `json:"externalTrafficPolicy,omitempty" yaml:"external_traffic_policy,omitempty"`
	HealthCheckNodePort      int64                  `json:"healthCheckNodePort,omitempty" yaml:"health_check_node_port,omitempty"`
	Hostname                 string                 `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	LoadBalancerIP           string                 `json:"loadBalancerIP,omitempty" yaml:"load_balancer_ip,omitempty"`
	LoadBalancerSourceRanges []string               `json:"loadBalancerSourceRanges,omitempty" yaml:"load_balancer_source_ranges,omitempty"`
	Ports                    []ServicePort          `json:"ports,omitempty" yaml:"ports,omitempty"`
	PublishNotReadyAddresses bool                   `json:"publishNotReadyAddresses,omitempty" yaml:"publish_not_ready_addresses,omitempty"`
	Selector                 map[string]string      `json:"selector,omitempty" yaml:"selector,omitempty"`
	ServiceKind              string                 `json:"serviceKind,omitempty" yaml:"service_kind,omitempty"`
	SessionAffinity          string                 `json:"sessionAffinity,omitempty" yaml:"session_affinity,omitempty"`
	SessionAffinityConfig    *SessionAffinityConfig `json:"sessionAffinityConfig,omitempty" yaml:"session_affinity_config,omitempty"`
}
