package client

const (
	AzureADConfigType                      = "azureADConfig"
	AzureADConfigFieldAccessMode           = "accessMode"
	AzureADConfigFieldAdminAccountPassword = "adminAccountPassword"
	AzureADConfigFieldAdminAccountUsername = "adminAccountUsername"
	AzureADConfigFieldAllowedPrincipalIDs  = "allowedPrincipalIds"
	AzureADConfigFieldAnnotations          = "annotations"
	AzureADConfigFieldClientId             = "clientid"
	AzureADConfigFieldCreated              = "created"
	AzureADConfigFieldCreatorID            = "creatorId"
	AzureADConfigFieldDomain               = "domain"
	AzureADConfigFieldEnabled              = "enabled"
	AzureADConfigFieldLabels               = "labels"
	AzureADConfigFieldName                 = "name"
	AzureADConfigFieldOwnerReferences      = "ownerReferences"
	AzureADConfigFieldRemoved              = "removed"
	AzureADConfigFieldTenantId             = "tenantid"
	AzureADConfigFieldType                 = "type"
	AzureADConfigFieldUuid                 = "uuid"
)

type AzureADConfig struct {
	AccessMode           string            `json:"accessMode,omitempty" yaml:"accessMode,omitempty"`
	AdminAccountPassword string            `json:"adminAccountPassword,omitempty" yaml:"adminAccountPassword,omitempty"`
	AdminAccountUsername string            `json:"adminAccountUsername,omitempty" yaml:"adminAccountUsername,omitempty"`
	AllowedPrincipalIDs  []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowedPrincipalIds,omitempty"`
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClientId             string            `json:"clientid,omitempty" yaml:"clientid,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Domain               string            `json:"domain,omitempty" yaml:"domain,omitempty"`
	Enabled              bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TenantId             string            `json:"tenantid,omitempty" yaml:"tenantid,omitempty"`
	Type                 string            `json:"type,omitempty" yaml:"type,omitempty"`
	Uuid                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
