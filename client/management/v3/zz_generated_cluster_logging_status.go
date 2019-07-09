package client

const (
	ClusterLoggingStatusType             = "clusterLoggingStatus"
	ClusterLoggingStatusFieldAppliedSpec = "appliedSpec"
	ClusterLoggingStatusFieldConditions  = "conditions"
	ClusterLoggingStatusFieldFailedSpec  = "failedSpec"
)

type ClusterLoggingStatus struct {
	AppliedSpec *ClusterLoggingSpec `json:"appliedSpec,omitempty" yaml:"applied_spec,omitempty"`
	Conditions  []LoggingCondition  `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	FailedSpec  *ClusterLoggingSpec `json:"failedSpec,omitempty" yaml:"failed_spec,omitempty"`
}
