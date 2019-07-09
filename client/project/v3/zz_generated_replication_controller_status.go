package client

const (
	ReplicationControllerStatusType                      = "replicationControllerStatus"
	ReplicationControllerStatusFieldAvailableReplicas    = "availableReplicas"
	ReplicationControllerStatusFieldConditions           = "conditions"
	ReplicationControllerStatusFieldFullyLabeledReplicas = "fullyLabeledReplicas"
	ReplicationControllerStatusFieldObservedGeneration   = "observedGeneration"
	ReplicationControllerStatusFieldReadyReplicas        = "readyReplicas"
	ReplicationControllerStatusFieldReplicas             = "replicas"
)

type ReplicationControllerStatus struct {
	AvailableReplicas    int64                            `json:"availableReplicas,omitempty" yaml:"available_replicas,omitempty"`
	Conditions           []ReplicationControllerCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	FullyLabeledReplicas int64                            `json:"fullyLabeledReplicas,omitempty" yaml:"fully_labeled_replicas,omitempty"`
	ObservedGeneration   int64                            `json:"observedGeneration,omitempty" yaml:"observed_generation,omitempty"`
	ReadyReplicas        int64                            `json:"readyReplicas,omitempty" yaml:"ready_replicas,omitempty"`
	Replicas             int64                            `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}
