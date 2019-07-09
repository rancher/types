package client

const (
	ProjectAlertSpecType                       = "projectAlertSpec"
	ProjectAlertSpecFieldDescription           = "description"
	ProjectAlertSpecFieldDisplayName           = "displayName"
	ProjectAlertSpecFieldInitialWaitSeconds    = "initialWaitSeconds"
	ProjectAlertSpecFieldProjectID             = "projectId"
	ProjectAlertSpecFieldRecipients            = "recipients"
	ProjectAlertSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ProjectAlertSpecFieldSeverity              = "severity"
	ProjectAlertSpecFieldTargetPod             = "targetPod"
	ProjectAlertSpecFieldTargetWorkload        = "targetWorkload"
)

type ProjectAlertSpec struct {
	Description           string          `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string          `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	InitialWaitSeconds    int64           `json:"initialWaitSeconds,omitempty" yaml:"initial_wait_seconds,omitempty"`
	ProjectID             string          `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Recipients            []Recipient     `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64           `json:"repeatIntervalSeconds,omitempty" yaml:"repeat_interval_seconds,omitempty"`
	Severity              string          `json:"severity,omitempty" yaml:"severity,omitempty"`
	TargetPod             *TargetPod      `json:"targetPod,omitempty" yaml:"target_pod,omitempty"`
	TargetWorkload        *TargetWorkload `json:"targetWorkload,omitempty" yaml:"target_workload,omitempty"`
}
