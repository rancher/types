package client

const (
	TargetNodeType              = "targetNode"
	TargetNodeFieldCPUThreshold = "cpuThreshold"
	TargetNodeFieldCondition    = "condition"
	TargetNodeFieldMemThreshold = "memThreshold"
	TargetNodeFieldNodeID       = "nodeId"
	TargetNodeFieldSelector     = "selector"
)

type TargetNode struct {
	CPUThreshold int64             `json:"cpuThreshold,omitempty" yaml:"cpu_threshold,omitempty"`
	Condition    string            `json:"condition,omitempty" yaml:"condition,omitempty"`
	MemThreshold int64             `json:"memThreshold,omitempty" yaml:"mem_threshold,omitempty"`
	NodeID       string            `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	Selector     map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
}
