package client

const (
	ClusterGroupSpecType                       = "clusterGroupSpec"
	ClusterGroupSpecFieldClusterID             = "clusterId"
	ClusterGroupSpecFieldDescription           = "description"
	ClusterGroupSpecFieldDisplayName           = "displayName"
	ClusterGroupSpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ClusterGroupSpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	ClusterGroupSpecFieldRecipients            = "recipients"
	ClusterGroupSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
)

type ClusterGroupSpec struct {
	ClusterID             string      `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Description           string      `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName           string      `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	GroupIntervalSeconds  int64       `json:"groupIntervalSeconds,omitempty" yaml:"group_interval_seconds,omitempty"`
	GroupWaitSeconds      int64       `json:"groupWaitSeconds,omitempty" yaml:"group_wait_seconds,omitempty"`
	Recipients            []Recipient `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	RepeatIntervalSeconds int64       `json:"repeatIntervalSeconds,omitempty" yaml:"repeat_interval_seconds,omitempty"`
}
