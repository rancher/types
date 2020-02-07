package client

const (
    RollingUpdateDeploymentType = "rollingUpdateDeployment"
	RollingUpdateDeploymentFieldMaxSurge = "maxSurge"
	RollingUpdateDeploymentFieldMaxUnavailable = "maxUnavailable"
)

type RollingUpdateDeployment struct {
        MaxSurge intstr.IntOrString `json:"maxSurge,omitempty" yaml:"maxSurge,omitempty"`
        MaxUnavailable intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
}

