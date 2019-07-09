package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	DeploymentConfigType                         = "deploymentConfig"
	DeploymentConfigFieldMaxSurge                = "maxSurge"
	DeploymentConfigFieldMaxUnavailable          = "maxUnavailable"
	DeploymentConfigFieldMinReadySeconds         = "minReadySeconds"
	DeploymentConfigFieldProgressDeadlineSeconds = "progressDeadlineSeconds"
	DeploymentConfigFieldRevisionHistoryLimit    = "revisionHistoryLimit"
	DeploymentConfigFieldStrategy                = "strategy"
)

type DeploymentConfig struct {
	MaxSurge                intstr.IntOrString `json:"maxSurge,omitempty" yaml:"max_surge,omitempty"`
	MaxUnavailable          intstr.IntOrString `json:"maxUnavailable,omitempty" yaml:"max_unavailable,omitempty"`
	MinReadySeconds         int64              `json:"minReadySeconds,omitempty" yaml:"min_ready_seconds,omitempty"`
	ProgressDeadlineSeconds *int64             `json:"progressDeadlineSeconds,omitempty" yaml:"progress_deadline_seconds,omitempty"`
	RevisionHistoryLimit    *int64             `json:"revisionHistoryLimit,omitempty" yaml:"revision_history_limit,omitempty"`
	Strategy                string             `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
