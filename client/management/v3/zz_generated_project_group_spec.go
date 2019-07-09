package client

const (
	ProjectGroupSpecType                       = "projectGroupSpec"
	ProjectGroupSpecFieldDescription           = "description"
	ProjectGroupSpecFieldDisplayName           = "displayName"
	ProjectGroupSpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ProjectGroupSpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	ProjectGroupSpecFieldProjectID             = "projectId"
	ProjectGroupSpecFieldRecipients            = "recipients"
	ProjectGroupSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
)

type ProjectGroupSpec struct {
	Description           string      `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string      `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	GroupIntervalSeconds  int64       `json:"groupIntervalSeconds,omitempty" yaml:"group_interval_seconds,omitempty"`
	GroupWaitSeconds      int64       `json:"groupWaitSeconds,omitempty" yaml:"group_wait_seconds,omitempty"`
	ProjectID             string      `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Recipients            []Recipient `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64       `json:"repeatIntervalSeconds,omitempty" yaml:"repeat_interval_seconds,omitempty"`
}
