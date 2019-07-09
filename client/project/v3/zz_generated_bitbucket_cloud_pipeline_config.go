package client

const (
	BitbucketCloudPipelineConfigType                 = "bitbucketCloudPipelineConfig"
	BitbucketCloudPipelineConfigFieldAnnotations     = "annotations"
	BitbucketCloudPipelineConfigFieldClientID        = "clientId"
	BitbucketCloudPipelineConfigFieldClientSecret    = "clientSecret"
	BitbucketCloudPipelineConfigFieldCreated         = "created"
	BitbucketCloudPipelineConfigFieldCreatorID       = "creatorId"
	BitbucketCloudPipelineConfigFieldEnabled         = "enabled"
	BitbucketCloudPipelineConfigFieldLabels          = "labels"
	BitbucketCloudPipelineConfigFieldName            = "name"
	BitbucketCloudPipelineConfigFieldNamespaceId     = "namespaceId"
	BitbucketCloudPipelineConfigFieldOwnerReferences = "ownerReferences"
	BitbucketCloudPipelineConfigFieldProjectID       = "projectId"
	BitbucketCloudPipelineConfigFieldRedirectURL     = "redirectUrl"
	BitbucketCloudPipelineConfigFieldRemoved         = "removed"
	BitbucketCloudPipelineConfigFieldType            = "type"
	BitbucketCloudPipelineConfigFieldUUID            = "uuid"
)

type BitbucketCloudPipelineConfig struct {
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClientID        string            `json:"clientId,omitempty" yaml:"client_id,omitempty"`
	ClientSecret    string            `json:"clientSecret,omitempty" yaml:"client_secret,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Enabled         bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	ProjectID       string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	RedirectURL     string            `json:"redirectUrl,omitempty" yaml:"redirect_url,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Type            string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
