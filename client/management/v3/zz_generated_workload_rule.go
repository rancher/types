package client

const (
	WorkloadRuleType                     = "workloadRule"
	WorkloadRuleFieldAvailablePercentage = "availablePercentage"
	WorkloadRuleFieldSelector            = "selector"
	WorkloadRuleFieldWorkloadID          = "workloadId"
)

type WorkloadRule struct {
	AvailablePercentage int64             `json:"availablePercentage,omitempty" yaml:"available_percentage,omitempty"`
	Selector            map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	WorkloadID          string            `json:"workloadId,omitempty" yaml:"workload_id,omitempty"`
}
