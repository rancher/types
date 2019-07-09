package client

const (
	ProjectAlertRuleSpecType                       = "projectAlertRuleSpec"
	ProjectAlertRuleSpecFieldDisplayName           = "displayName"
	ProjectAlertRuleSpecFieldGroupID               = "groupId"
	ProjectAlertRuleSpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ProjectAlertRuleSpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	ProjectAlertRuleSpecFieldInherited             = "inherited"
	ProjectAlertRuleSpecFieldMetricRule            = "metricRule"
	ProjectAlertRuleSpecFieldPodRule               = "podRule"
	ProjectAlertRuleSpecFieldProjectID             = "projectId"
	ProjectAlertRuleSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ProjectAlertRuleSpecFieldSeverity              = "severity"
	ProjectAlertRuleSpecFieldWorkloadRule          = "workloadRule"
)

type ProjectAlertRuleSpec struct {
	DisplayName           string        `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	GroupID               string        `json:"groupId,omitempty" yaml:"group_id,omitempty"`
	GroupIntervalSeconds  int64         `json:"groupIntervalSeconds,omitempty" yaml:"group_interval_seconds,omitempty"`
	GroupWaitSeconds      int64         `json:"groupWaitSeconds,omitempty" yaml:"group_wait_seconds,omitempty"`
	Inherited             *bool         `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	MetricRule            *MetricRule   `json:"metricRule,omitempty" yaml:"metric_rule,omitempty"`
	PodRule               *PodRule      `json:"podRule,omitempty" yaml:"pod_rule,omitempty"`
	ProjectID             string        `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	RepeatIntervalSeconds int64         `json:"repeatIntervalSeconds,omitempty" yaml:"repeat_interval_seconds,omitempty"`
	Severity              string        `json:"severity,omitempty" yaml:"severity,omitempty"`
	WorkloadRule          *WorkloadRule `json:"workloadRule,omitempty" yaml:"workload_rule,omitempty"`
}
