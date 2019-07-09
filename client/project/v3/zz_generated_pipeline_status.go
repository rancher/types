package client

const (
	PipelineStatusType                      = "pipelineStatus"
	PipelineStatusFieldLastExecutionID      = "lastExecutionId"
	PipelineStatusFieldLastRunState         = "lastRunState"
	PipelineStatusFieldLastStarted          = "lastStarted"
	PipelineStatusFieldNextRun              = "nextRun"
	PipelineStatusFieldNextStart            = "nextStart"
	PipelineStatusFieldPipelineState        = "pipelineState"
	PipelineStatusFieldSourceCodeCredential = "sourceCodeCredential"
	PipelineStatusFieldToken                = "token"
	PipelineStatusFieldWebHookID            = "webhookId"
)

type PipelineStatus struct {
	LastExecutionID      string                `json:"lastExecutionId,omitempty" yaml:"last_execution_id,omitempty"`
	LastRunState         string                `json:"lastRunState,omitempty" yaml:"last_run_state,omitempty"`
	LastStarted          string                `json:"lastStarted,omitempty" yaml:"last_started,omitempty"`
	NextRun              int64                 `json:"nextRun,omitempty" yaml:"next_run,omitempty"`
	NextStart            string                `json:"nextStart,omitempty" yaml:"next_start,omitempty"`
	PipelineState        string                `json:"pipelineState,omitempty" yaml:"pipeline_state,omitempty"`
	SourceCodeCredential *SourceCodeCredential `json:"sourceCodeCredential,omitempty" yaml:"source_code_credential,omitempty"`
	Token                string                `json:"token,omitempty" yaml:"token,omitempty"`
	WebHookID            string                `json:"webhookId,omitempty" yaml:"webhook_id,omitempty"`
}
