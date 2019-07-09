package client

const (
	GoogleOauthConfigType                          = "googleOauthConfig"
	GoogleOauthConfigFieldAccessMode               = "accessMode"
	GoogleOauthConfigFieldAdminEmail               = "adminEmail"
	GoogleOauthConfigFieldAllowedPrincipalIDs      = "allowedPrincipalIds"
	GoogleOauthConfigFieldAnnotations              = "annotations"
	GoogleOauthConfigFieldCreated                  = "created"
	GoogleOauthConfigFieldCreatorID                = "creatorId"
	GoogleOauthConfigFieldEnabled                  = "enabled"
	GoogleOauthConfigFieldHostname                 = "hostname"
	GoogleOauthConfigFieldLabels                   = "labels"
	GoogleOauthConfigFieldName                     = "name"
	GoogleOauthConfigFieldOauthCredential          = "oauthCredential"
	GoogleOauthConfigFieldOwnerReferences          = "ownerReferences"
	GoogleOauthConfigFieldRemoved                  = "removed"
	GoogleOauthConfigFieldServiceAccountCredential = "serviceAccountCredential"
	GoogleOauthConfigFieldType                     = "type"
	GoogleOauthConfigFieldUUID                     = "uuid"
	GoogleOauthConfigFieldUserInfoEndpoint         = "userInfoEndpoint"
)

type GoogleOauthConfig struct {
	AccessMode               string            `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`
	AdminEmail               string            `json:"adminEmail,omitempty" yaml:"admin_email,omitempty"`
	AllowedPrincipalIDs      []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowed_principal_ids,omitempty"`
	Annotations              map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created                  string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Enabled                  bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Hostname                 string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Labels                   map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                     string            `json:"name,omitempty" yaml:"name,omitempty"`
	OauthCredential          string            `json:"oauthCredential,omitempty" yaml:"oauth_credential,omitempty"`
	OwnerReferences          []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Removed                  string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	ServiceAccountCredential string            `json:"serviceAccountCredential,omitempty" yaml:"service_account_credential,omitempty"`
	Type                     string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID                     string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UserInfoEndpoint         string            `json:"userInfoEndpoint,omitempty" yaml:"user_info_endpoint,omitempty"`
}
