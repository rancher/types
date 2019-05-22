package client

const (
	ExampleStatusType            = "exampleStatus"
	ExampleStatusFieldConditions = "conditions"
)

type ExampleStatus struct {
	Conditions []ComposeCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
