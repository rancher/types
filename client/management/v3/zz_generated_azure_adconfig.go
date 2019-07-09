package client

const (
	AzureADConfigType                     = "azureADConfig"
	AzureADConfigFieldAccessMode          = "accessMode"
	AzureADConfigFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	AzureADConfigFieldAnnotations         = "annotations"
	AzureADConfigFieldApplicationID       = "applicationId"
	AzureADConfigFieldApplicationSecret   = "applicationSecret"
	AzureADConfigFieldAuthEndpoint        = "authEndpoint"
	AzureADConfigFieldCreated             = "created"
	AzureADConfigFieldCreatorID           = "creatorId"
	AzureADConfigFieldEnabled             = "enabled"
	AzureADConfigFieldEndpoint            = "endpoint"
	AzureADConfigFieldGraphEndpoint       = "graphEndpoint"
	AzureADConfigFieldLabels              = "labels"
	AzureADConfigFieldName                = "name"
	AzureADConfigFieldOwnerReferences     = "ownerReferences"
	AzureADConfigFieldRancherURL          = "rancherUrl"
	AzureADConfigFieldRemoved             = "removed"
	AzureADConfigFieldTenantID            = "tenantId"
	AzureADConfigFieldTokenEndpoint       = "tokenEndpoint"
	AzureADConfigFieldType                = "type"
	AzureADConfigFieldUUID                = "uuid"
)

type AzureADConfig struct {
	AccessMode          string            `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowed_principal_ids,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ApplicationID       string            `json:"applicationId,omitempty" yaml:"application_id,omitempty"`
	ApplicationSecret   string            `json:"applicationSecret,omitempty" yaml:"application_secret,omitempty"`
	AuthEndpoint        string            `json:"authEndpoint,omitempty" yaml:"auth_endpoint,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Endpoint            string            `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	GraphEndpoint       string            `json:"graphEndpoint,omitempty" yaml:"graph_endpoint,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	RancherURL          string            `json:"rancherUrl,omitempty" yaml:"rancher_url,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TenantID            string            `json:"tenantId,omitempty" yaml:"tenant_id,omitempty"`
	TokenEndpoint       string            `json:"tokenEndpoint,omitempty" yaml:"token_endpoint,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
