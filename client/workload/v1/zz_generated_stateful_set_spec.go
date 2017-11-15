package client

const (
	StatefulSetSpecType                      = "statefulSetSpec"
	StatefulSetSpecFieldDeploy               = "deploy"
	StatefulSetSpecFieldServiceName          = "serviceName"
	StatefulSetSpecFieldTemplate             = "template"
	StatefulSetSpecFieldUpdateStrategy       = "updateStrategy"
	StatefulSetSpecFieldVolumeClaimTemplates = "volumeClaimTemplates"
)

type StatefulSetSpec struct {
	Deploy               *DeployParams                    `json:"deploy,omitempty"`
	ServiceName          string                           `json:"serviceName,omitempty"`
	Template             PodTemplateSpec                  `json:"template,omitempty"`
	UpdateStrategy       StatefulSetUpdateStrategy        `json:"updateStrategy,omitempty"`
	VolumeClaimTemplates map[string]PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"`
}
