package client

const (
	OpenLdapConfigType                                 = "openLdapConfig"
	OpenLdapConfigFieldAccessMode                      = "accessMode"
	OpenLdapConfigFieldAllowedPrincipalIDs             = "allowedPrincipalIds"
	OpenLdapConfigFieldAnnotations                     = "annotations"
	OpenLdapConfigFieldCertificate                     = "certificate"
	OpenLdapConfigFieldConnectionTimeout               = "connectionTimeout"
	OpenLdapConfigFieldCreated                         = "created"
	OpenLdapConfigFieldCreatorID                       = "creatorId"
	OpenLdapConfigFieldEnabled                         = "enabled"
	OpenLdapConfigFieldGroupDNAttribute                = "groupDNAttribute"
	OpenLdapConfigFieldGroupMemberMappingAttribute     = "groupMemberMappingAttribute"
	OpenLdapConfigFieldGroupMemberUserAttribute        = "groupMemberUserAttribute"
	OpenLdapConfigFieldGroupNameAttribute              = "groupNameAttribute"
	OpenLdapConfigFieldGroupObjectClass                = "groupObjectClass"
	OpenLdapConfigFieldGroupSearchAttribute            = "groupSearchAttribute"
	OpenLdapConfigFieldGroupSearchBase                 = "groupSearchBase"
	OpenLdapConfigFieldLabels                          = "labels"
	OpenLdapConfigFieldName                            = "name"
	OpenLdapConfigFieldNestedGroupMembershipEnabled    = "nestedGroupMembershipEnabled"
	OpenLdapConfigFieldOwnerReferences                 = "ownerReferences"
	OpenLdapConfigFieldPort                            = "port"
	OpenLdapConfigFieldRemoved                         = "removed"
	OpenLdapConfigFieldServers                         = "servers"
	OpenLdapConfigFieldServiceAccountDistinguishedName = "serviceAccountDistinguishedName"
	OpenLdapConfigFieldServiceAccountPassword          = "serviceAccountPassword"
	OpenLdapConfigFieldTLS                             = "tls"
	OpenLdapConfigFieldType                            = "type"
	OpenLdapConfigFieldUUID                            = "uuid"
	OpenLdapConfigFieldUserDisabledBitMask             = "userDisabledBitMask"
	OpenLdapConfigFieldUserEnabledAttribute            = "userEnabledAttribute"
	OpenLdapConfigFieldUserLoginAttribute              = "userLoginAttribute"
	OpenLdapConfigFieldUserMemberAttribute             = "userMemberAttribute"
	OpenLdapConfigFieldUserNameAttribute               = "userNameAttribute"
	OpenLdapConfigFieldUserObjectClass                 = "userObjectClass"
	OpenLdapConfigFieldUserSearchAttribute             = "userSearchAttribute"
	OpenLdapConfigFieldUserSearchBase                  = "userSearchBase"
)

type OpenLdapConfig struct {
	AccessMode                      string            `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`
	AllowedPrincipalIDs             []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowed_principal_ids,omitempty"`
	Annotations                     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Certificate                     string            `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	ConnectionTimeout               int64             `json:"connectionTimeout,omitempty" yaml:"connection_timeout,omitempty"`
	Created                         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                       string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Enabled                         bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	GroupDNAttribute                string            `json:"groupDNAttribute,omitempty" yaml:"group_dn_attribute,omitempty"`
	GroupMemberMappingAttribute     string            `json:"groupMemberMappingAttribute,omitempty" yaml:"group_member_mapping_attribute,omitempty"`
	GroupMemberUserAttribute        string            `json:"groupMemberUserAttribute,omitempty" yaml:"group_member_user_attribute,omitempty"`
	GroupNameAttribute              string            `json:"groupNameAttribute,omitempty" yaml:"group_name_attribute,omitempty"`
	GroupObjectClass                string            `json:"groupObjectClass,omitempty" yaml:"group_object_class,omitempty"`
	GroupSearchAttribute            string            `json:"groupSearchAttribute,omitempty" yaml:"group_search_attribute,omitempty"`
	GroupSearchBase                 string            `json:"groupSearchBase,omitempty" yaml:"group_search_base,omitempty"`
	Labels                          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NestedGroupMembershipEnabled    bool              `json:"nestedGroupMembershipEnabled,omitempty" yaml:"nested_group_membership_enabled,omitempty"`
	OwnerReferences                 []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Port                            int64             `json:"port,omitempty" yaml:"port,omitempty"`
	Removed                         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Servers                         []string          `json:"servers,omitempty" yaml:"servers,omitempty"`
	ServiceAccountDistinguishedName string            `json:"serviceAccountDistinguishedName,omitempty" yaml:"service_account_distinguished_name,omitempty"`
	ServiceAccountPassword          string            `json:"serviceAccountPassword,omitempty" yaml:"service_account_password,omitempty"`
	TLS                             bool              `json:"tls,omitempty" yaml:"tls,omitempty"`
	Type                            string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID                            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UserDisabledBitMask             int64             `json:"userDisabledBitMask,omitempty" yaml:"user_disabled_bit_mask,omitempty"`
	UserEnabledAttribute            string            `json:"userEnabledAttribute,omitempty" yaml:"user_enabled_attribute,omitempty"`
	UserLoginAttribute              string            `json:"userLoginAttribute,omitempty" yaml:"user_login_attribute,omitempty"`
	UserMemberAttribute             string            `json:"userMemberAttribute,omitempty" yaml:"user_member_attribute,omitempty"`
	UserNameAttribute               string            `json:"userNameAttribute,omitempty" yaml:"user_name_attribute,omitempty"`
	UserObjectClass                 string            `json:"userObjectClass,omitempty" yaml:"user_object_class,omitempty"`
	UserSearchAttribute             string            `json:"userSearchAttribute,omitempty" yaml:"user_search_attribute,omitempty"`
	UserSearchBase                  string            `json:"userSearchBase,omitempty" yaml:"user_search_base,omitempty"`
}
