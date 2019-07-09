package client

const (
	SourceCodeCredentialSpecType                = "sourceCodeCredentialSpec"
	SourceCodeCredentialSpecFieldAccessToken    = "accessToken"
	SourceCodeCredentialSpecFieldAvatarURL      = "avatarUrl"
	SourceCodeCredentialSpecFieldDisplayName    = "displayName"
	SourceCodeCredentialSpecFieldExpiry         = "expiry"
	SourceCodeCredentialSpecFieldGitCloneToken  = "gitCloneToken"
	SourceCodeCredentialSpecFieldGitLoginName   = "gitLoginName"
	SourceCodeCredentialSpecFieldHTMLURL        = "htmlUrl"
	SourceCodeCredentialSpecFieldLoginName      = "loginName"
	SourceCodeCredentialSpecFieldProjectID      = "projectId"
	SourceCodeCredentialSpecFieldRefreshToken   = "refreshToken"
	SourceCodeCredentialSpecFieldSourceCodeType = "sourceCodeType"
	SourceCodeCredentialSpecFieldUserID         = "userId"
)

type SourceCodeCredentialSpec struct {
	AccessToken    string `json:"accessToken,omitempty" yaml:"access_token,omitempty"`
	AvatarURL      string `json:"avatarUrl,omitempty" yaml:"avatar_url,omitempty"`
	DisplayName    string `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	Expiry         string `json:"expiry,omitempty" yaml:"expiry,omitempty"`
	GitCloneToken  string `json:"gitCloneToken,omitempty" yaml:"git_clone_token,omitempty"`
	GitLoginName   string `json:"gitLoginName,omitempty" yaml:"git_login_name,omitempty"`
	HTMLURL        string `json:"htmlUrl,omitempty" yaml:"html_url,omitempty"`
	LoginName      string `json:"loginName,omitempty" yaml:"login_name,omitempty"`
	ProjectID      string `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	RefreshToken   string `json:"refreshToken,omitempty" yaml:"refresh_token,omitempty"`
	SourceCodeType string `json:"sourceCodeType,omitempty" yaml:"source_code_type,omitempty"`
	UserID         string `json:"userId,omitempty" yaml:"user_id,omitempty"`
}
