package client

const (
	AffinityType                 = "affinity"
	AffinityFieldNodeAffinity    = "nodeAffinity"
	AffinityFieldPodAffinity     = "podAffinity"
	AffinityFieldPodAntiAffinity = "podAntiAffinity"
)

type Affinity struct {
	NodeAffinity    *NodeAffinity    `json:"nodeAffinity,omitempty" yaml:"node_affinity,omitempty"`
	PodAffinity     *PodAffinity     `json:"podAffinity,omitempty" yaml:"pod_affinity,omitempty"`
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty" yaml:"pod_anti_affinity,omitempty"`
}
