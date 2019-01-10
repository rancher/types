package client

const (
	MultiClusterAppSpecType                   = "multiClusterAppSpec"
	MultiClusterAppSpecFieldAnswers           = "answers"
	MultiClusterAppSpecFieldMembers           = "members"
	MultiClusterAppSpecFieldTargets           = "targets"
	MultiClusterAppSpecFieldTemplateVersionID = "templateVersionId"
)

type MultiClusterAppSpec struct {
	Answers           []Answer `json:"answers,omitempty" yaml:"answers,omitempty"`
	Members           []Member `json:"members,omitempty" yaml:"members,omitempty"`
	Targets           []Target `json:"targets,omitempty" yaml:"targets,omitempty"`
	TemplateVersionID string   `json:"templateVersionId,omitempty" yaml:"templateVersionId,omitempty"`
}
