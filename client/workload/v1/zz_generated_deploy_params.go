package client

const (
	DeployParamsType                         = "deployParams"
	DeployParamsFieldMinReadySeconds         = "minReadySeconds"
	DeployParamsFieldPodManagementPolicy     = "podManagementPolicy"
	DeployParamsFieldProgressDeadlineSeconds = "progressDeadlineSeconds"
	DeployParamsFieldReplicas                = "replicas"
	DeployParamsFieldRevisionHistoryLimit    = "revisionHistoryLimit"
)

type DeployParams struct {
	MinReadySeconds         int64  `json:"minReadySeconds,omitempty"`
	PodManagementPolicy     string `json:"podManagementPolicy,omitempty"`
	ProgressDeadlineSeconds *int64 `json:"progressDeadlineSeconds,omitempty"`
	Replicas                *int64 `json:"replicas,omitempty"`
	RevisionHistoryLimit    *int64 `json:"revisionHistoryLimit,omitempty"`
}
