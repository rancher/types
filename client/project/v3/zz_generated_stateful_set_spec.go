package client

const (
	StatefulSetSpecType             = "statefulSetSpec"
	StatefulSetSpecFieldSelector    = "selector"
	StatefulSetSpecFieldStatefulSet = "statefulSet"
	StatefulSetSpecFieldTemplate    = "template"
)

type StatefulSetSpec struct {
	Selector    *LabelSelector     `json:"selector,omitempty"`
	StatefulSet *StatefulSetConfig `json:"statefulSet,omitempty"`
	Template    *PodTemplateSpec   `json:"template,omitempty"`
}
