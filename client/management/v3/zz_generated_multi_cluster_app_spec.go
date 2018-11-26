package client

const (
	MultiClusterAppSpecType                   = "multiClusterAppSpec"
	MultiClusterAppSpecFieldAnswers           = "answers"
	MultiClusterAppSpecFieldCatalogID         = "catalogId"
	MultiClusterAppSpecFieldTargets           = "targets"
	MultiClusterAppSpecFieldTemplateVersionID = "templateVersionId"
)

type MultiClusterAppSpec struct {
	Answers           []Answer `json:"answers,omitempty" yaml:"answers,omitempty"`
	CatalogID         string   `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`
	Targets           []Target `json:"targets,omitempty" yaml:"targets,omitempty"`
	TemplateVersionID string   `json:"templateVersionId,omitempty" yaml:"templateVersionId,omitempty"`
}
