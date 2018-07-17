package client

const (
	PodSchedulingType              = "podScheduling"
	PodSchedulingFieldAffinity     = "affinity"
	PodSchedulingFieldAntiAffinity = "antiAffinity"
)

type PodScheduling struct {
	Affinity     []PodAffinityRule `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	AntiAffinity []PodAffinityRule `json:"antiAffinity,omitempty" yaml:"antiAffinity,omitempty"`
}
