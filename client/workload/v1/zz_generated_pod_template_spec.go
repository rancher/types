package client

const (
	PodTemplateSpecType            = "podTemplateSpec"
	PodTemplateSpecFieldObjectMeta = "objectMeta"
	PodTemplateSpecFieldSpec       = "spec"
)

type PodTemplateSpec struct {
	ObjectMeta ObjectMeta `json:"objectMeta,omitempty"`
	Spec       PodSpec    `json:"spec,omitempty"`
}
