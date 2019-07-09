package client

const (
	ClusterStatusType                                      = "clusterStatus"
	ClusterStatusFieldAPIEndpoint                          = "apiEndpoint"
	ClusterStatusFieldAgentImage                           = "agentImage"
	ClusterStatusFieldAllocatable                          = "allocatable"
	ClusterStatusFieldAppliedEnableNetworkPolicy           = "appliedEnableNetworkPolicy"
	ClusterStatusFieldAppliedPodSecurityPolicyTemplateName = "appliedPodSecurityPolicyTemplateId"
	ClusterStatusFieldAppliedSpec                          = "appliedSpec"
	ClusterStatusFieldAuthImage                            = "authImage"
	ClusterStatusFieldCACert                               = "caCert"
	ClusterStatusFieldCapabilities                         = "capabilities"
	ClusterStatusFieldCapacity                             = "capacity"
	ClusterStatusFieldCertificatesExpiration               = "certificatesExpiration"
	ClusterStatusFieldComponentStatuses                    = "componentStatuses"
	ClusterStatusFieldConditions                           = "conditions"
	ClusterStatusFieldDriver                               = "driver"
	ClusterStatusFieldFailedSpec                           = "failedSpec"
	ClusterStatusFieldIstioEnabled                         = "istioEnabled"
	ClusterStatusFieldLimits                               = "limits"
	ClusterStatusFieldMonitoringStatus                     = "monitoringStatus"
	ClusterStatusFieldRequested                            = "requested"
	ClusterStatusFieldVersion                              = "version"
)

type ClusterStatus struct {
	APIEndpoint                          string                    `json:"apiEndpoint,omitempty" yaml:"api_endpoint,omitempty"`
	AgentImage                           string                    `json:"agentImage,omitempty" yaml:"agent_image,omitempty"`
	Allocatable                          map[string]string         `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	AppliedEnableNetworkPolicy           bool                      `json:"appliedEnableNetworkPolicy,omitempty" yaml:"applied_enable_network_policy,omitempty"`
	AppliedPodSecurityPolicyTemplateName string                    `json:"appliedPodSecurityPolicyTemplateId,omitempty" yaml:"applied_pod_security_policy_template_id,omitempty"`
	AppliedSpec                          *ClusterSpec              `json:"appliedSpec,omitempty" yaml:"applied_spec,omitempty"`
	AuthImage                            string                    `json:"authImage,omitempty" yaml:"auth_image,omitempty"`
	CACert                               string                    `json:"caCert,omitempty" yaml:"ca_cert,omitempty"`
	Capabilities                         *Capabilities             `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Capacity                             map[string]string         `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	CertificatesExpiration               map[string]CertExpiration `json:"certificatesExpiration,omitempty" yaml:"certificates_expiration,omitempty"`
	ComponentStatuses                    []ClusterComponentStatus  `json:"componentStatuses,omitempty" yaml:"component_statuses,omitempty"`
	Conditions                           []ClusterCondition        `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Driver                               string                    `json:"driver,omitempty" yaml:"driver,omitempty"`
	FailedSpec                           *ClusterSpec              `json:"failedSpec,omitempty" yaml:"failed_spec,omitempty"`
	IstioEnabled                         bool                      `json:"istioEnabled,omitempty" yaml:"istio_enabled,omitempty"`
	Limits                               map[string]string         `json:"limits,omitempty" yaml:"limits,omitempty"`
	MonitoringStatus                     *MonitoringStatus         `json:"monitoringStatus,omitempty" yaml:"monitoring_status,omitempty"`
	Requested                            map[string]string         `json:"requested,omitempty" yaml:"requested,omitempty"`
	Version                              *Info                     `json:"version,omitempty" yaml:"version,omitempty"`
}
