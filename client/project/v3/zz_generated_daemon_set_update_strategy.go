package client

const (
    DaemonSetUpdateStrategyType = "daemonSetUpdateStrategy"
	DaemonSetUpdateStrategyFieldMaxUnavailable = "maxUnavailable"
	DaemonSetUpdateStrategyFieldStrategy = "strategy"
)

type DaemonSetUpdateStrategy struct {
        MaxUnavailable intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
        Strategy string `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}

