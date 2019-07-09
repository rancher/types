package client

const (
	AzureADProviderType                 = "azureADProvider"
	AzureADProviderFieldAnnotations     = "annotations"
	AzureADProviderFieldCreated         = "created"
	AzureADProviderFieldCreatorID       = "creatorId"
	AzureADProviderFieldLabels          = "labels"
	AzureADProviderFieldName            = "name"
	AzureADProviderFieldOwnerReferences = "ownerReferences"
	AzureADProviderFieldRedirectURL     = "redirectUrl"
	AzureADProviderFieldRemoved         = "removed"
	AzureADProviderFieldType            = "type"
	AzureADProviderFieldUUID            = "uuid"
)

type AzureADProvider struct {
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	RedirectURL     string            `json:"redirectUrl,omitempty" yaml:"redirect_url,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Type            string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
