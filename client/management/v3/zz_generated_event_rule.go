package client

const (
	EventRuleType              = "eventRule"
	EventRuleFieldEventType    = "eventType"
	EventRuleFieldResourceKind = "resourceKind"
)

type EventRule struct {
	EventType    string `json:"eventType,omitempty" yaml:"event_type,omitempty"`
	ResourceKind string `json:"resourceKind,omitempty" yaml:"resource_kind,omitempty"`
}
