package client

const (
	DaemonSetStatusType                        = "daemonSetStatus"
	DaemonSetStatusFieldCollisionCount         = "collisionCount"
	DaemonSetStatusFieldConditions             = "conditions"
	DaemonSetStatusFieldCurrentNumberScheduled = "currentNumberScheduled"
	DaemonSetStatusFieldDesiredNumberScheduled = "desiredNumberScheduled"
	DaemonSetStatusFieldNumberAvailable        = "numberAvailable"
	DaemonSetStatusFieldNumberMisscheduled     = "numberMisscheduled"
	DaemonSetStatusFieldNumberReady            = "numberReady"
	DaemonSetStatusFieldNumberUnavailable      = "numberUnavailable"
	DaemonSetStatusFieldObservedGeneration     = "observedGeneration"
	DaemonSetStatusFieldUpdatedNumberScheduled = "updatedNumberScheduled"
)

type DaemonSetStatus struct {
	CollisionCount         *int64               `json:"collisionCount,omitempty" yaml:"collision_count,omitempty"`
	Conditions             []DaemonSetCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	CurrentNumberScheduled int64                `json:"currentNumberScheduled,omitempty" yaml:"current_number_scheduled,omitempty"`
	DesiredNumberScheduled int64                `json:"desiredNumberScheduled,omitempty" yaml:"desired_number_scheduled,omitempty"`
	NumberAvailable        int64                `json:"numberAvailable,omitempty" yaml:"number_available,omitempty"`
	NumberMisscheduled     int64                `json:"numberMisscheduled,omitempty" yaml:"number_misscheduled,omitempty"`
	NumberReady            int64                `json:"numberReady,omitempty" yaml:"number_ready,omitempty"`
	NumberUnavailable      int64                `json:"numberUnavailable,omitempty" yaml:"number_unavailable,omitempty"`
	ObservedGeneration     int64                `json:"observedGeneration,omitempty" yaml:"observed_generation,omitempty"`
	UpdatedNumberScheduled int64                `json:"updatedNumberScheduled,omitempty" yaml:"updated_number_scheduled,omitempty"`
}
