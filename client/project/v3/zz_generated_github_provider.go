package client

const (
	GithubProviderType                 = "githubProvider"
	GithubProviderFieldAnnotations     = "annotations"
	GithubProviderFieldCreated         = "created"
	GithubProviderFieldCreatorID       = "creatorId"
	GithubProviderFieldLabels          = "labels"
	GithubProviderFieldName            = "name"
	GithubProviderFieldOwnerReferences = "ownerReferences"
	GithubProviderFieldProjectID       = "projectId"
	GithubProviderFieldRedirectURL     = "redirectUrl"
	GithubProviderFieldRemoved         = "removed"
	GithubProviderFieldType            = "type"
	GithubProviderFieldUUID            = "uuid"
)

type GithubProvider struct {
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	ProjectID       string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	RedirectURL     string            `json:"redirectUrl,omitempty" yaml:"redirect_url,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Type            string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
