package client

const (
	CronJobSpecType          = "cronJobSpec"
	CronJobSpecFieldCronJob  = "cronJob"
	CronJobSpecFieldSelector = "selector"
	CronJobSpecFieldTemplate = "template"
)

type CronJobSpec struct {
	CronJob  *CronJobConfig   `json:"cronJob,omitempty"`
	Selector *LabelSelector   `json:"selector,omitempty"`
	Template *PodTemplateSpec `json:"template,omitempty"`
}
