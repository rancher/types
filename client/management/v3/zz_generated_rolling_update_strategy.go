package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	RollingUpdateStrategyType                = "rollingUpdateStrategy"
	RollingUpdateStrategyFieldMaxUnavailable = "maxUnavailable"
)

type RollingUpdateStrategy struct {
	MaxUnavailable intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
}
