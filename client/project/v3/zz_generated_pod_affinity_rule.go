package client

const (
	PodAffinityRuleType           = "podAffinityRule"
	PodAffinityRuleFieldPreferred = "preferred"
	PodAffinityRuleFieldRequired  = "required"
	PodAffinityRuleFieldTopology  = "topology"
)

type PodAffinityRule struct {
	Preferred []string `json:"preferred,omitempty" yaml:"preferred,omitempty"`
	Required  []string `json:"required,omitempty" yaml:"required,omitempty"`
	Topology  string   `json:"topology,omitempty" yaml:"topology,omitempty"`
}
