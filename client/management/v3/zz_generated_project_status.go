package client

const (
	ProjectStatusType                               = "projectStatus"
	ProjectStatusFieldConditions                    = "conditions"
	ProjectStatusFieldMonitoringStatus              = "monitoringStatus"
	ProjectStatusFieldPodSecurityPolicyTemplateName = "podSecurityPolicyTemplateId"
)

type ProjectStatus struct {
	Conditions                    []ProjectCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	MonitoringStatus              *MonitoringStatus  `json:"monitoringStatus,omitempty" yaml:"monitoring_status,omitempty"`
	PodSecurityPolicyTemplateName string             `json:"podSecurityPolicyTemplateId,omitempty" yaml:"pod_security_policy_template_id,omitempty"`
}
