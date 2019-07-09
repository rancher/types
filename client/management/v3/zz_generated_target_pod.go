package client

const (
	TargetPodType                        = "targetPod"
	TargetPodFieldCondition              = "condition"
	TargetPodFieldPodID                  = "podId"
	TargetPodFieldRestartIntervalSeconds = "restartIntervalSeconds"
	TargetPodFieldRestartTimes           = "restartTimes"
)

type TargetPod struct {
	Condition              string `json:"condition,omitempty" yaml:"condition,omitempty"`
	PodID                  string `json:"podId,omitempty" yaml:"pod_id,omitempty"`
	RestartIntervalSeconds int64  `json:"restartIntervalSeconds,omitempty" yaml:"restart_interval_seconds,omitempty"`
	RestartTimes           int64  `json:"restartTimes,omitempty" yaml:"restart_times,omitempty"`
}
