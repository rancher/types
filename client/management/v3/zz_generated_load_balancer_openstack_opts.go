package client

const (
	LoadBalancerOpenstackOptsType                      = "loadBalancerOpenstackOpts"
	LoadBalancerOpenstackOptsFieldCreateMonitor        = "create-monitor"
	LoadBalancerOpenstackOptsFieldFloatingNetworkID    = "floating-network-id"
	LoadBalancerOpenstackOptsFieldLBMethod             = "lb-method"
	LoadBalancerOpenstackOptsFieldLBProvider           = "lb-provider"
	LoadBalancerOpenstackOptsFieldLBVersion            = "lb-version"
	LoadBalancerOpenstackOptsFieldManageSecurityGroups = "manage-security-groups"
	LoadBalancerOpenstackOptsFieldMonitorDelay         = "monitor-delay"
	LoadBalancerOpenstackOptsFieldMonitorMaxRetries    = "monitor-max-retries"
	LoadBalancerOpenstackOptsFieldMonitorTimeout       = "monitor-timeout"
	LoadBalancerOpenstackOptsFieldSubnetID             = "subnet-id"
	LoadBalancerOpenstackOptsFieldUseOctavia           = "use-octavia"
)

type LoadBalancerOpenstackOpts struct {
	CreateMonitor        bool   `json:"create-monitor,omitempty" yaml:"create_monitor,omitempty"`
	FloatingNetworkID    string `json:"floating-network-id,omitempty" yaml:"floating_network_id,omitempty"`
	LBMethod             string `json:"lb-method,omitempty" yaml:"lb_method,omitempty"`
	LBProvider           string `json:"lb-provider,omitempty" yaml:"lb_provider,omitempty"`
	LBVersion            string `json:"lb-version,omitempty" yaml:"lb_version,omitempty"`
	ManageSecurityGroups bool   `json:"manage-security-groups,omitempty" yaml:"manage_security_groups,omitempty"`
	MonitorDelay         string `json:"monitor-delay,omitempty" yaml:"monitor_delay,omitempty"`
	MonitorMaxRetries    int64  `json:"monitor-max-retries,omitempty" yaml:"monitor_max_retries,omitempty"`
	MonitorTimeout       string `json:"monitor-timeout,omitempty" yaml:"monitor_timeout,omitempty"`
	SubnetID             string `json:"subnet-id,omitempty" yaml:"subnet_id,omitempty"`
	UseOctavia           bool   `json:"use-octavia,omitempty" yaml:"use_octavia,omitempty"`
}
