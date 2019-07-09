package client

const (
	ContainerResourceLimitType                = "containerResourceLimit"
	ContainerResourceLimitFieldLimitsCPU      = "limitsCpu"
	ContainerResourceLimitFieldLimitsMemory   = "limitsMemory"
	ContainerResourceLimitFieldRequestsCPU    = "requestsCpu"
	ContainerResourceLimitFieldRequestsMemory = "requestsMemory"
)

type ContainerResourceLimit struct {
	LimitsCPU      string `json:"limitsCpu,omitempty" yaml:"limits_cpu,omitempty"`
	LimitsMemory   string `json:"limitsMemory,omitempty" yaml:"limits_memory,omitempty"`
	RequestsCPU    string `json:"requestsCpu,omitempty" yaml:"requests_cpu,omitempty"`
	RequestsMemory string `json:"requestsMemory,omitempty" yaml:"requests_memory,omitempty"`
}
