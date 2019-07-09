package client

const (
	InternalNodeSpecType               = "internalNodeSpec"
	InternalNodeSpecFieldPodCidr       = "podCidr"
	InternalNodeSpecFieldProviderId    = "providerId"
	InternalNodeSpecFieldTaints        = "taints"
	InternalNodeSpecFieldUnschedulable = "unschedulable"
)

type InternalNodeSpec struct {
	PodCidr       string  `json:"podCidr,omitempty" yaml:"pod_cidr,omitempty"`
	ProviderId    string  `json:"providerId,omitempty" yaml:"provider_id,omitempty"`
	Taints        []Taint `json:"taints,omitempty" yaml:"taints,omitempty"`
	Unschedulable bool    `json:"unschedulable,omitempty" yaml:"unschedulable,omitempty"`
}
