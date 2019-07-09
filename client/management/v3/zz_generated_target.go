package client

const (
	TargetType             = "target"
	TargetFieldAppID       = "appId"
	TargetFieldHealthstate = "healthState"
	TargetFieldProjectID   = "projectId"
	TargetFieldState       = "state"
)

type Target struct {
	AppID       string `json:"appId,omitempty" yaml:"app_id,omitempty"`
	Healthstate string `json:"healthState,omitempty" yaml:"health_state,omitempty"`
	ProjectID   string `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	State       string `json:"state,omitempty" yaml:"state,omitempty"`
}
