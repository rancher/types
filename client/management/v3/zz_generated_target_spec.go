package client

const (
	TargetSpecType               = "targetSpec"
	TargetSpecFieldAnswers       = "answers"
	TargetSpecFieldClusterConfig = "clusterConfig"
	TargetSpecFieldClusterID     = "clusterId"
)

type TargetSpec struct {
	Answers       map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterConfig *ClusterConfig    `json:"clusterConfig,omitempty" yaml:"clusterConfig,omitempty"`
	ClusterID     string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
}
