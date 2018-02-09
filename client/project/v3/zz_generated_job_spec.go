package client

const (
	JobSpecType          = "jobSpec"
	JobSpecFieldJob      = "job"
	JobSpecFieldSelector = "selector"
	JobSpecFieldTemplate = "template"
)

type JobSpec struct {
	Job      *JobConfig       `json:"job,omitempty"`
	Selector *LabelSelector   `json:"selector,omitempty"`
	Template *PodTemplateSpec `json:"template,omitempty"`
}
