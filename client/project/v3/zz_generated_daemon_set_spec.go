package client

const (
	DaemonSetSpecType           = "daemonSetSpec"
	DaemonSetSpecFieldDaemonSet = "daemonSet"
	DaemonSetSpecFieldSelector  = "selector"
	DaemonSetSpecFieldTemplate  = "template"
)

type DaemonSetSpec struct {
	DaemonSet *DaemonSetConfig `json:"daemonSet,omitempty"`
	Selector  *LabelSelector   `json:"selector,omitempty"`
	Template  *PodTemplateSpec `json:"template,omitempty"`
}
