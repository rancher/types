package client

const (
	StackSpecType                   = "stackSpec"
	StackSpecFieldAnswers           = "answers"
	StackSpecFieldDescription       = "description"
	StackSpecFieldGroups            = "groups"
	StackSpecFieldInstallNamespace  = "installNamespace"
	StackSpecFieldProjectId         = "projectId"
	StackSpecFieldPrune             = "prune"
	StackSpecFieldTag               = "tag"
	StackSpecFieldTemplateVersionID = "templateVersionId"
	StackSpecFieldTemplates         = "templates"
	StackSpecFieldUser              = "user"
)

type StackSpec struct {
	Answers           map[string]string `json:"answers,omitempty"`
	Description       string            `json:"description,omitempty"`
	Groups            []string          `json:"groups,omitempty"`
	InstallNamespace  string            `json:"installNamespace,omitempty"`
	ProjectId         string            `json:"projectId,omitempty"`
	Prune             *bool             `json:"prune,omitempty"`
	Tag               map[string]string `json:"tag,omitempty"`
	TemplateVersionID string            `json:"templateVersionId,omitempty"`
	Templates         map[string]string `json:"templates,omitempty"`
	User              string            `json:"user,omitempty"`
}
