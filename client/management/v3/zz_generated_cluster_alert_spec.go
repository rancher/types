package client

const (
	ClusterAlertSpecType                       = "clusterAlertSpec"
	ClusterAlertSpecFieldClusterID             = "clusterId"
	ClusterAlertSpecFieldDescription           = "description"
	ClusterAlertSpecFieldDisplayName           = "displayName"
	ClusterAlertSpecFieldInitialWaitSeconds    = "initialWaitSeconds"
	ClusterAlertSpecFieldRecipients            = "recipients"
	ClusterAlertSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ClusterAlertSpecFieldSeverity              = "severity"
	ClusterAlertSpecFieldTargetEvent           = "targetEvent"
	ClusterAlertSpecFieldTargetNode            = "targetNode"
	ClusterAlertSpecFieldTargetSystemService   = "targetSystemService"
)

type ClusterAlertSpec struct {
	ClusterID             string               `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Description           string               `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string               `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	InitialWaitSeconds    int64                `json:"initialWaitSeconds,omitempty" yaml:"initial_wait_seconds,omitempty"`
	Recipients            []Recipient          `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64                `json:"repeatIntervalSeconds,omitempty" yaml:"repeat_interval_seconds,omitempty"`
	Severity              string               `json:"severity,omitempty" yaml:"severity,omitempty"`
	TargetEvent           *TargetEvent         `json:"targetEvent,omitempty" yaml:"target_event,omitempty"`
	TargetNode            *TargetNode          `json:"targetNode,omitempty" yaml:"target_node,omitempty"`
	TargetSystemService   *TargetSystemService `json:"targetSystemService,omitempty" yaml:"target_system_service,omitempty"`
}
