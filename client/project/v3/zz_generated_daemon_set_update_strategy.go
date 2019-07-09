package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	DaemonSetUpdateStrategyType                = "daemonSetUpdateStrategy"
	DaemonSetUpdateStrategyFieldMaxUnavailable = "maxUnavailable"
	DaemonSetUpdateStrategyFieldStrategy       = "strategy"
)

type DaemonSetUpdateStrategy struct {
	MaxUnavailable intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"max_unavailable,omitempty"`
	Strategy       string             `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
