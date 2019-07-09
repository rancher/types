package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	DaemonSetConfigType                      = "daemonSetConfig"
	DaemonSetConfigFieldMaxUnavailable       = "maxUnavailable"
	DaemonSetConfigFieldMinReadySeconds      = "minReadySeconds"
	DaemonSetConfigFieldRevisionHistoryLimit = "revisionHistoryLimit"
	DaemonSetConfigFieldStrategy             = "strategy"
)

type DaemonSetConfig struct {
	MaxUnavailable       intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"max_unavailable,omitempty"`
	MinReadySeconds      int64              `json:"minReadySeconds,omitempty" yaml:"min_ready_seconds,omitempty"`
	RevisionHistoryLimit *int64             `json:"revisionHistoryLimit,omitempty" yaml:"revision_history_limit,omitempty"`
	Strategy             string             `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
