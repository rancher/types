package client

const (
	NodeConditionType                    = "nodeCondition"
	NodeConditionFieldLastHeartbeatTime  = "lastHeartbeatTime"
	NodeConditionFieldLastTransitionTime = "lastTransitionTime"
	NodeConditionFieldMessage            = "message"
	NodeConditionFieldReason             = "reason"
	NodeConditionFieldStatus             = "status"
	NodeConditionFieldType               = "type"
)

type NodeCondition struct {
	LastHeartbeatTime  string `json:"lastHeartbeatTime,omitempty" yaml:"last_heartbeat_time,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"last_transition_time,omitempty"`
	Message            string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason             string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Status             string `json:"status,omitempty" yaml:"status,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
}
