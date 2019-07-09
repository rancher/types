package client

const (
	AppSpecType                   = "appSpec"
	AppSpecFieldAnswers           = "answers"
	AppSpecFieldAppRevisionID     = "appRevisionId"
	AppSpecFieldDescription       = "description"
	AppSpecFieldExternalID        = "externalId"
	AppSpecFieldFiles             = "files"
	AppSpecFieldMultiClusterAppID = "multiClusterAppId"
	AppSpecFieldProjectID         = "projectId"
	AppSpecFieldPrune             = "prune"
	AppSpecFieldTargetNamespace   = "targetNamespace"
	AppSpecFieldValuesYaml        = "valuesYaml"
)

type AppSpec struct {
	Answers           map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	AppRevisionID     string            `json:"appRevisionId,omitempty" yaml:"app_revision_id,omitempty"`
	Description       string            `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalID        string            `json:"externalId,omitempty" yaml:"external_id,omitempty"`
	Files             map[string]string `json:"files,omitempty" yaml:"files,omitempty"`
	MultiClusterAppID string            `json:"multiClusterAppId,omitempty" yaml:"multi_cluster_app_id,omitempty"`
	ProjectID         string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Prune             bool              `json:"prune,omitempty" yaml:"prune,omitempty"`
	TargetNamespace   string            `json:"targetNamespace,omitempty" yaml:"target_namespace,omitempty"`
	ValuesYaml        string            `json:"valuesYaml,omitempty" yaml:"values_yaml,omitempty"`
}
