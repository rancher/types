package client

const (
	NodeSelectorTermType                  = "nodeSelectorTerm"
	NodeSelectorTermFieldMatchExpressions = "matchExpressions"
	NodeSelectorTermFieldMatchFields      = "matchFields"
)

type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty" yaml:"match_expressions,omitempty"`
	MatchFields      []NodeSelectorRequirement `json:"matchFields,omitempty" yaml:"match_fields,omitempty"`
}
