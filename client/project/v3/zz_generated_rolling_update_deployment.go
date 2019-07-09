package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	RollingUpdateDeploymentType                = "rollingUpdateDeployment"
	RollingUpdateDeploymentFieldMaxSurge       = "maxSurge"
	RollingUpdateDeploymentFieldMaxUnavailable = "maxUnavailable"
)

type RollingUpdateDeployment struct {
	MaxSurge       intstr.IntOrString `json:"maxSurge,omitempty" yaml:"max_surge,omitempty"`
	MaxUnavailable intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"max_unavailable,omitempty"`
}
