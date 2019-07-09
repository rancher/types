package client

const (
	ComposeConditionType                    = "composeCondition"
	ComposeConditionFieldLastTransitionTime = "lastTransitionTime"
	ComposeConditionFieldLastUpdateTime     = "lastUpdateTime"
	ComposeConditionFieldMessage            = "message"
	ComposeConditionFieldReason             = "reason"
	ComposeConditionFieldStatus             = "status"
	ComposeConditionFieldType               = "type"
)

type ComposeCondition struct {
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"last_transition_time,omitempty"`
	LastUpdateTime     string `json:"lastUpdateTime,omitempty" yaml:"last_update_time,omitempty"`
	Message            string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason             string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Status             string `json:"status,omitempty" yaml:"status,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
}
