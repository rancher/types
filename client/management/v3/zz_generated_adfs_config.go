package client

const (
	ADFSConfigType                     = "adfsConfig"
	ADFSConfigFieldAccessMode          = "accessMode"
	ADFSConfigFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	ADFSConfigFieldAnnotations         = "annotations"
	ADFSConfigFieldCreated             = "created"
	ADFSConfigFieldCreatorID           = "creatorId"
	ADFSConfigFieldDisplayNameField    = "displayNameField"
	ADFSConfigFieldEnabled             = "enabled"
	ADFSConfigFieldGroupsField         = "groupsField"
	ADFSConfigFieldIDPMetadataContent  = "idpMetadataContent"
	ADFSConfigFieldLabels              = "labels"
	ADFSConfigFieldName                = "name"
	ADFSConfigFieldOwnerReferences     = "ownerReferences"
	ADFSConfigFieldRancherAPIHost      = "rancherApiHost"
	ADFSConfigFieldRemoved             = "removed"
	ADFSConfigFieldSpCert              = "spCert"
	ADFSConfigFieldSpKey               = "spKey"
	ADFSConfigFieldType                = "type"
	ADFSConfigFieldUIDField            = "uidField"
	ADFSConfigFieldUUID                = "uuid"
	ADFSConfigFieldUserNameField       = "userNameField"
)

type ADFSConfig struct {
	AccessMode          string            `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowed_principal_ids,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DisplayNameField    string            `json:"displayNameField,omitempty" yaml:"display_name_field,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	GroupsField         string            `json:"groupsField,omitempty" yaml:"groups_field,omitempty"`
	IDPMetadataContent  string            `json:"idpMetadataContent,omitempty" yaml:"idp_metadata_content,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	RancherAPIHost      string            `json:"rancherApiHost,omitempty" yaml:"rancher_api_host,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	SpCert              string            `json:"spCert,omitempty" yaml:"sp_cert,omitempty"`
	SpKey               string            `json:"spKey,omitempty" yaml:"sp_key,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
	UIDField            string            `json:"uidField,omitempty" yaml:"uid_field,omitempty"`
	UUID                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UserNameField       string            `json:"userNameField,omitempty" yaml:"user_name_field,omitempty"`
}
