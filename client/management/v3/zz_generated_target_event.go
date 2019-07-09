package client

const (
	TargetEventType              = "targetEvent"
	TargetEventFieldEventType    = "eventType"
	TargetEventFieldResourceKind = "resourceKind"
)

type TargetEvent struct {
	EventType    string `json:"eventType,omitempty" yaml:"event_type,omitempty"`
	ResourceKind string `json:"resourceKind,omitempty" yaml:"resource_kind,omitempty"`
}
