package client

const (
	DeploymentSpecType            = "deploymentSpec"
	DeploymentSpecFieldDeployment = "deployment"
	DeploymentSpecFieldScale      = "scale"
	DeploymentSpecFieldSelector   = "selector"
	DeploymentSpecFieldTemplate   = "template"
)

type DeploymentSpec struct {
	Deployment *DeploymentConfig `json:"deployment,omitempty"`
	Scale      *int64            `json:"scale,omitempty"`
	Selector   *LabelSelector    `json:"selector,omitempty"`
	Template   *PodTemplateSpec  `json:"template,omitempty"`
}
