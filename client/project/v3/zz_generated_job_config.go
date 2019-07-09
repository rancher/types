package client

const (
	JobConfigType                       = "jobConfig"
	JobConfigFieldActiveDeadlineSeconds = "activeDeadlineSeconds"
	JobConfigFieldBackoffLimit          = "backoffLimit"
	JobConfigFieldCompletions           = "completions"
	JobConfigFieldManualSelector        = "manualSelector"
	JobConfigFieldParallelism           = "parallelism"
)

type JobConfig struct {
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	BackoffLimit          *int64 `json:"backoffLimit,omitempty" yaml:"backoff_limit,omitempty"`
	Completions           *int64 `json:"completions,omitempty" yaml:"completions,omitempty"`
	ManualSelector        *bool  `json:"manualSelector,omitempty" yaml:"manual_selector,omitempty"`
	Parallelism           *int64 `json:"parallelism,omitempty" yaml:"parallelism,omitempty"`
}
