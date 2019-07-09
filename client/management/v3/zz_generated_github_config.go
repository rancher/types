package client

const (
	GithubConfigType                     = "githubConfig"
	GithubConfigFieldAccessMode          = "accessMode"
	GithubConfigFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	GithubConfigFieldAnnotations         = "annotations"
	GithubConfigFieldClientID            = "clientId"
	GithubConfigFieldClientSecret        = "clientSecret"
	GithubConfigFieldCreated             = "created"
	GithubConfigFieldCreatorID           = "creatorId"
	GithubConfigFieldEnabled             = "enabled"
	GithubConfigFieldHostname            = "hostname"
	GithubConfigFieldLabels              = "labels"
	GithubConfigFieldName                = "name"
	GithubConfigFieldOwnerReferences     = "ownerReferences"
	GithubConfigFieldRemoved             = "removed"
	GithubConfigFieldTLS                 = "tls"
	GithubConfigFieldType                = "type"
	GithubConfigFieldUUID                = "uuid"
)

type GithubConfig struct {
	AccessMode          string            `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowed_principal_ids,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClientID            string            `json:"clientId,omitempty" yaml:"client_id,omitempty"`
	ClientSecret        string            `json:"clientSecret,omitempty" yaml:"client_secret,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Hostname            string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TLS                 bool              `json:"tls,omitempty" yaml:"tls,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
