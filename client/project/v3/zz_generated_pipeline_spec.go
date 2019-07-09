package client

const (
	PipelineSpecType                        = "pipelineSpec"
	PipelineSpecFieldDisplayName            = "displayName"
	PipelineSpecFieldProjectID              = "projectId"
	PipelineSpecFieldRepositoryURL          = "repositoryUrl"
	PipelineSpecFieldSourceCodeCredentialID = "sourceCodeCredentialId"
	PipelineSpecFieldTriggerWebhookPr       = "triggerWebhookPr"
	PipelineSpecFieldTriggerWebhookPush     = "triggerWebhookPush"
	PipelineSpecFieldTriggerWebhookTag      = "triggerWebhookTag"
)

type PipelineSpec struct {
	DisplayName            string `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	ProjectID              string `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	RepositoryURL          string `json:"repositoryUrl,omitempty" yaml:"repository_url,omitempty"`
	SourceCodeCredentialID string `json:"sourceCodeCredentialId,omitempty" yaml:"source_code_credential_id,omitempty"`
	TriggerWebhookPr       bool   `json:"triggerWebhookPr,omitempty" yaml:"trigger_webhook_pr,omitempty"`
	TriggerWebhookPush     bool   `json:"triggerWebhookPush,omitempty" yaml:"trigger_webhook_push,omitempty"`
	TriggerWebhookTag      bool   `json:"triggerWebhookTag,omitempty" yaml:"trigger_webhook_tag,omitempty"`
}
