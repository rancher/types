package client

const (
	ReplicaSetSpecType          = "replicaSetSpec"
	ReplicaSetSpecFieldDeploy   = "deploy"
	ReplicaSetSpecFieldTemplate = "template"
)

type ReplicaSetSpec struct {
	Deploy   *DeployParams   `json:"deploy,omitempty"`
	Template PodTemplateSpec `json:"template,omitempty"`
}
