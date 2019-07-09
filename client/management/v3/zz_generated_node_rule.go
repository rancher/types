package client

const (
	NodeRuleType              = "nodeRule"
	NodeRuleFieldCPUThreshold = "cpuThreshold"
	NodeRuleFieldCondition    = "condition"
	NodeRuleFieldMemThreshold = "memThreshold"
	NodeRuleFieldNodeID       = "nodeId"
	NodeRuleFieldSelector     = "selector"
)

type NodeRule struct {
	CPUThreshold int64             `json:"cpuThreshold,omitempty" yaml:"cpu_threshold,omitempty"`
	Condition    string            `json:"condition,omitempty" yaml:"condition,omitempty"`
	MemThreshold int64             `json:"memThreshold,omitempty" yaml:"mem_threshold,omitempty"`
	NodeID       string            `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	Selector     map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
}
