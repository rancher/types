package client

const (
    RollingUpdateDaemonSetType = "rollingUpdateDaemonSet"
	RollingUpdateDaemonSetFieldMaxUnavailable = "maxUnavailable"
)

type RollingUpdateDaemonSet struct {
        MaxUnavailable intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
}

