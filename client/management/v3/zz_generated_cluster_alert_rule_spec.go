package client

const (
	ClusterAlertRuleSpecType                       = "clusterAlertRuleSpec"
	ClusterAlertRuleSpecFieldClusterID             = "clusterId"
	ClusterAlertRuleSpecFieldDisplayName           = "displayName"
	ClusterAlertRuleSpecFieldEventRule             = "eventRule"
	ClusterAlertRuleSpecFieldGroupID               = "groupId"
	ClusterAlertRuleSpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ClusterAlertRuleSpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	ClusterAlertRuleSpecFieldInherited             = "inherited"
	ClusterAlertRuleSpecFieldMetricRule            = "metricRule"
	ClusterAlertRuleSpecFieldNodeRule              = "nodeRule"
	ClusterAlertRuleSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ClusterAlertRuleSpecFieldSeverity              = "severity"
	ClusterAlertRuleSpecFieldSystemServiceRule     = "systemServiceRule"
)

type ClusterAlertRuleSpec struct {
	ClusterID             string             `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	DisplayName           string             `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	EventRule             *EventRule         `json:"eventRule,omitempty" yaml:"event_rule,omitempty"`
	GroupID               string             `json:"groupId,omitempty" yaml:"group_id,omitempty"`
	GroupIntervalSeconds  int64              `json:"groupIntervalSeconds,omitempty" yaml:"group_interval_seconds,omitempty"`
	GroupWaitSeconds      int64              `json:"groupWaitSeconds,omitempty" yaml:"group_wait_seconds,omitempty"`
	Inherited             *bool              `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	MetricRule            *MetricRule        `json:"metricRule,omitempty" yaml:"metric_rule,omitempty"`
	NodeRule              *NodeRule          `json:"nodeRule,omitempty" yaml:"node_rule,omitempty"`
	RepeatIntervalSeconds int64              `json:"repeatIntervalSeconds,omitempty" yaml:"repeat_interval_seconds,omitempty"`
	Severity              string             `json:"severity,omitempty" yaml:"severity,omitempty"`
	SystemServiceRule     *SystemServiceRule `json:"systemServiceRule,omitempty" yaml:"system_service_rule,omitempty"`
}
