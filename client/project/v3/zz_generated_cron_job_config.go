package client

const (
	CronJobConfigType                            = "cronJobConfig"
	CronJobConfigFieldConcurrencyPolicy          = "concurrencyPolicy"
	CronJobConfigFieldFailedJobsHistoryLimit     = "failedJobsHistoryLimit"
	CronJobConfigFieldJobAnnotations             = "jobAnnotations"
	CronJobConfigFieldJobConfig                  = "jobConfig"
	CronJobConfigFieldJobLabels                  = "jobLabels"
	CronJobConfigFieldSchedule                   = "schedule"
	CronJobConfigFieldStartingDeadlineSeconds    = "startingDeadlineSeconds"
	CronJobConfigFieldSuccessfulJobsHistoryLimit = "successfulJobsHistoryLimit"
	CronJobConfigFieldSuspend                    = "suspend"
)

type CronJobConfig struct {
	ConcurrencyPolicy          string            `json:"concurrencyPolicy,omitempty" yaml:"concurrency_policy,omitempty"`
	FailedJobsHistoryLimit     *int64            `json:"failedJobsHistoryLimit,omitempty" yaml:"failed_jobs_history_limit,omitempty"`
	JobAnnotations             map[string]string `json:"jobAnnotations,omitempty" yaml:"job_annotations,omitempty"`
	JobConfig                  *JobConfig        `json:"jobConfig,omitempty" yaml:"job_config,omitempty"`
	JobLabels                  map[string]string `json:"jobLabels,omitempty" yaml:"job_labels,omitempty"`
	Schedule                   string            `json:"schedule,omitempty" yaml:"schedule,omitempty"`
	StartingDeadlineSeconds    *int64            `json:"startingDeadlineSeconds,omitempty" yaml:"starting_deadline_seconds,omitempty"`
	SuccessfulJobsHistoryLimit *int64            `json:"successfulJobsHistoryLimit,omitempty" yaml:"successful_jobs_history_limit,omitempty"`
	Suspend                    *bool             `json:"suspend,omitempty" yaml:"suspend,omitempty"`
}
