package client

const (
	StatefulSetStatusType                    = "statefulSetStatus"
	StatefulSetStatusFieldCollisionCount     = "collisionCount"
	StatefulSetStatusFieldConditions         = "conditions"
	StatefulSetStatusFieldCurrentReplicas    = "currentReplicas"
	StatefulSetStatusFieldCurrentRevision    = "currentRevision"
	StatefulSetStatusFieldObservedGeneration = "observedGeneration"
	StatefulSetStatusFieldReadyReplicas      = "readyReplicas"
	StatefulSetStatusFieldReplicas           = "replicas"
	StatefulSetStatusFieldUpdateRevision     = "updateRevision"
	StatefulSetStatusFieldUpdatedReplicas    = "updatedReplicas"
)

type StatefulSetStatus struct {
	CollisionCount     *int64                 `json:"collisionCount,omitempty" yaml:"collision_count,omitempty"`
	Conditions         []StatefulSetCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	CurrentReplicas    int64                  `json:"currentReplicas,omitempty" yaml:"current_replicas,omitempty"`
	CurrentRevision    string                 `json:"currentRevision,omitempty" yaml:"current_revision,omitempty"`
	ObservedGeneration int64                  `json:"observedGeneration,omitempty" yaml:"observed_generation,omitempty"`
	ReadyReplicas      int64                  `json:"readyReplicas,omitempty" yaml:"ready_replicas,omitempty"`
	Replicas           int64                  `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	UpdateRevision     string                 `json:"updateRevision,omitempty" yaml:"update_revision,omitempty"`
	UpdatedReplicas    int64                  `json:"updatedReplicas,omitempty" yaml:"updated_replicas,omitempty"`
}
