package client

const (
	JobConditionType                    = "jobCondition"
	JobConditionFieldLastProbeTime      = "lastProbeTime"
	JobConditionFieldLastTransitionTime = "lastTransitionTime"
	JobConditionFieldMessage            = "message"
	JobConditionFieldReason             = "reason"
	JobConditionFieldStatus             = "status"
	JobConditionFieldType               = "type"
)

type JobCondition struct {
	LastProbeTime      string `json:"lastProbeTime,omitempty" yaml:"last_probe_time,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"last_transition_time,omitempty"`
	Message            string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason             string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Status             string `json:"status,omitempty" yaml:"status,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
}
