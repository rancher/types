package client

const (
	SourceCodeRepositorySpecType                        = "sourceCodeRepositorySpec"
	SourceCodeRepositorySpecFieldDefaultBranch          = "defaultBranch"
	SourceCodeRepositorySpecFieldLanguage               = "language"
	SourceCodeRepositorySpecFieldPermissions            = "permissions"
	SourceCodeRepositorySpecFieldProjectID              = "projectId"
	SourceCodeRepositorySpecFieldSourceCodeCredentialID = "sourceCodeCredentialId"
	SourceCodeRepositorySpecFieldSourceCodeType         = "sourceCodeType"
	SourceCodeRepositorySpecFieldURL                    = "url"
	SourceCodeRepositorySpecFieldUserID                 = "userId"
)

type SourceCodeRepositorySpec struct {
	DefaultBranch          string    `json:"defaultBranch,omitempty" yaml:"default_branch,omitempty"`
	Language               string    `json:"language,omitempty" yaml:"language,omitempty"`
	Permissions            *RepoPerm `json:"permissions,omitempty" yaml:"permissions,omitempty"`
	ProjectID              string    `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	SourceCodeCredentialID string    `json:"sourceCodeCredentialId,omitempty" yaml:"source_code_credential_id,omitempty"`
	SourceCodeType         string    `json:"sourceCodeType,omitempty" yaml:"source_code_type,omitempty"`
	URL                    string    `json:"url,omitempty" yaml:"url,omitempty"`
	UserID                 string    `json:"userId,omitempty" yaml:"user_id,omitempty"`
}
