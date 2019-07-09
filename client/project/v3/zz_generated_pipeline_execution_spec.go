package client

const (
	PipelineExecutionSpecType                = "pipelineExecutionSpec"
	PipelineExecutionSpecFieldAuthor         = "author"
	PipelineExecutionSpecFieldAvatarURL      = "avatarUrl"
	PipelineExecutionSpecFieldBranch         = "branch"
	PipelineExecutionSpecFieldCommit         = "commit"
	PipelineExecutionSpecFieldEmail          = "email"
	PipelineExecutionSpecFieldEvent          = "event"
	PipelineExecutionSpecFieldHTMLLink       = "htmlLink"
	PipelineExecutionSpecFieldMessage        = "message"
	PipelineExecutionSpecFieldPipelineConfig = "pipelineConfig"
	PipelineExecutionSpecFieldPipelineID     = "pipelineId"
	PipelineExecutionSpecFieldProjectID      = "projectId"
	PipelineExecutionSpecFieldRef            = "ref"
	PipelineExecutionSpecFieldRepositoryURL  = "repositoryUrl"
	PipelineExecutionSpecFieldRun            = "run"
	PipelineExecutionSpecFieldTitle          = "title"
	PipelineExecutionSpecFieldTriggerUserID  = "triggerUserId"
	PipelineExecutionSpecFieldTriggeredBy    = "triggeredBy"
)

type PipelineExecutionSpec struct {
	Author         string          `json:"author,omitempty" yaml:"author,omitempty"`
	AvatarURL      string          `json:"avatarUrl,omitempty" yaml:"avatar_url,omitempty"`
	Branch         string          `json:"branch,omitempty" yaml:"branch,omitempty"`
	Commit         string          `json:"commit,omitempty" yaml:"commit,omitempty"`
	Email          string          `json:"email,omitempty" yaml:"email,omitempty"`
	Event          string          `json:"event,omitempty" yaml:"event,omitempty"`
	HTMLLink       string          `json:"htmlLink,omitempty" yaml:"html_link,omitempty"`
	Message        string          `json:"message,omitempty" yaml:"message,omitempty"`
	PipelineConfig *PipelineConfig `json:"pipelineConfig,omitempty" yaml:"pipeline_config,omitempty"`
	PipelineID     string          `json:"pipelineId,omitempty" yaml:"pipeline_id,omitempty"`
	ProjectID      string          `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Ref            string          `json:"ref,omitempty" yaml:"ref,omitempty"`
	RepositoryURL  string          `json:"repositoryUrl,omitempty" yaml:"repository_url,omitempty"`
	Run            int64           `json:"run,omitempty" yaml:"run,omitempty"`
	Title          string          `json:"title,omitempty" yaml:"title,omitempty"`
	TriggerUserID  string          `json:"triggerUserId,omitempty" yaml:"trigger_user_id,omitempty"`
	TriggeredBy    string          `json:"triggeredBy,omitempty" yaml:"triggered_by,omitempty"`
}
