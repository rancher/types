package client

const (
	AppRevisionStatusType            = "appRevisionStatus"
	AppRevisionStatusFieldAnswers    = "answers"
	AppRevisionStatusFieldDigest     = "digest"
	AppRevisionStatusFieldExternalID = "externalId"
	AppRevisionStatusFieldFiles      = "files"
	AppRevisionStatusFieldProjectID  = "projectId"
	AppRevisionStatusFieldValuesYaml = "valuesYaml"
)

type AppRevisionStatus struct {
	Answers    map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	Digest     string            `json:"digest,omitempty" yaml:"digest,omitempty"`
	ExternalID string            `json:"externalId,omitempty" yaml:"external_id,omitempty"`
	Files      map[string]string `json:"files,omitempty" yaml:"files,omitempty"`
	ProjectID  string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	ValuesYaml string            `json:"valuesYaml,omitempty" yaml:"values_yaml,omitempty"`
}
