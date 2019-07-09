package client

const (
	DeploymentStatusType                     = "deploymentStatus"
	DeploymentStatusFieldAvailableReplicas   = "availableReplicas"
	DeploymentStatusFieldCollisionCount      = "collisionCount"
	DeploymentStatusFieldConditions          = "conditions"
	DeploymentStatusFieldObservedGeneration  = "observedGeneration"
	DeploymentStatusFieldReadyReplicas       = "readyReplicas"
	DeploymentStatusFieldReplicas            = "replicas"
	DeploymentStatusFieldUnavailableReplicas = "unavailableReplicas"
	DeploymentStatusFieldUpdatedReplicas     = "updatedReplicas"
)

type DeploymentStatus struct {
	AvailableReplicas   int64                 `json:"availableReplicas,omitempty" yaml:"available_replicas,omitempty"`
	CollisionCount      *int64                `json:"collisionCount,omitempty" yaml:"collision_count,omitempty"`
	Conditions          []DeploymentCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	ObservedGeneration  int64                 `json:"observedGeneration,omitempty" yaml:"observed_generation,omitempty"`
	ReadyReplicas       int64                 `json:"readyReplicas,omitempty" yaml:"ready_replicas,omitempty"`
	Replicas            int64                 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	UnavailableReplicas int64                 `json:"unavailableReplicas,omitempty" yaml:"unavailable_replicas,omitempty"`
	UpdatedReplicas     int64                 `json:"updatedReplicas,omitempty" yaml:"updated_replicas,omitempty"`
}
