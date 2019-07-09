package client

const (
	CronJobStatusType                  = "cronJobStatus"
	CronJobStatusFieldActive           = "active"
	CronJobStatusFieldLastScheduleTime = "lastScheduleTime"
)

type CronJobStatus struct {
	Active           []ObjectReference `json:"active,omitempty" yaml:"active,omitempty"`
	LastScheduleTime string            `json:"lastScheduleTime,omitempty" yaml:"last_schedule_time,omitempty"`
}
