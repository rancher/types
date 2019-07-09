package client

const (
	PodRuleType                        = "podRule"
	PodRuleFieldCondition              = "condition"
	PodRuleFieldPodID                  = "podId"
	PodRuleFieldRestartIntervalSeconds = "restartIntervalSeconds"
	PodRuleFieldRestartTimes           = "restartTimes"
)

type PodRule struct {
	Condition              string `json:"condition,omitempty" yaml:"condition,omitempty"`
	PodID                  string `json:"podId,omitempty" yaml:"pod_id,omitempty"`
	RestartIntervalSeconds int64  `json:"restartIntervalSeconds,omitempty" yaml:"restart_interval_seconds,omitempty"`
	RestartTimes           int64  `json:"restartTimes,omitempty" yaml:"restart_times,omitempty"`
}
