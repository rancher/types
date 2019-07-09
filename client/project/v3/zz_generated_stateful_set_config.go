package client

const (
	StatefulSetConfigType                      = "statefulSetConfig"
	StatefulSetConfigFieldPartition            = "partition"
	StatefulSetConfigFieldPodManagementPolicy  = "podManagementPolicy"
	StatefulSetConfigFieldRevisionHistoryLimit = "revisionHistoryLimit"
	StatefulSetConfigFieldServiceName          = "serviceName"
	StatefulSetConfigFieldStrategy             = "strategy"
	StatefulSetConfigFieldVolumeClaimTemplates = "volumeClaimTemplates"
)

type StatefulSetConfig struct {
	Partition            *int64                  `json:"partition,omitempty" yaml:"partition,omitempty"`
	PodManagementPolicy  string                  `json:"podManagementPolicy,omitempty" yaml:"pod_management_policy,omitempty"`
	RevisionHistoryLimit *int64                  `json:"revisionHistoryLimit,omitempty" yaml:"revision_history_limit,omitempty"`
	ServiceName          string                  `json:"serviceName,omitempty" yaml:"service_name,omitempty"`
	Strategy             string                  `json:"strategy,omitempty" yaml:"strategy,omitempty"`
	VolumeClaimTemplates []PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty" yaml:"volume_claim_templates,omitempty"`
}
