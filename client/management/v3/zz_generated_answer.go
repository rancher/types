package client

const (
	AnswerType           = "answer"
	AnswerFieldClusterID = "clusterId"
	AnswerFieldProjectID = "projectId"
	AnswerFieldValues    = "values"
)

type Answer struct {
	ClusterID string            `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	ProjectID string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Values    map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
}
