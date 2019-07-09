package client

const (
	NodeSelectorType                   = "nodeSelector"
	NodeSelectorFieldNodeSelectorTerms = "nodeSelectorTerms"
)

type NodeSelector struct {
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms,omitempty" yaml:"node_selector_terms,omitempty"`
}
