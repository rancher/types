package client

const (
	DeploymentSpecType          = "deploymentSpec"
	DeploymentSpecFieldDeploy   = "deploy"
	DeploymentSpecFieldPaused   = "paused"
	DeploymentSpecFieldStrategy = "strategy"
	DeploymentSpecFieldTemplate = "template"
)

type DeploymentSpec struct {
	Deploy   *DeployParams      `json:"deploy,omitempty"`
	Paused   bool               `json:"paused,omitempty"`
	Strategy DeploymentStrategy `json:"strategy,omitempty"`
	Template PodTemplateSpec    `json:"template,omitempty"`
}
