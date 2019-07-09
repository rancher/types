package client

import (
	"github.com/rancher/norman/types"
)

const (
	LdapConfigType                                 = "ldapConfig"
	LdapConfigFieldAccessMode                      = "accessMode"
	LdapConfigFieldAllowedPrincipalIDs             = "allowedPrincipalIds"
	LdapConfigFieldAnnotations                     = "annotations"
	LdapConfigFieldCertificate                     = "certificate"
	LdapConfigFieldConnectionTimeout               = "connectionTimeout"
	LdapConfigFieldCreated                         = "created"
	LdapConfigFieldCreatorID                       = "creatorId"
	LdapConfigFieldEnabled                         = "enabled"
	LdapConfigFieldGroupDNAttribute                = "groupDNAttribute"
	LdapConfigFieldGroupMemberMappingAttribute     = "groupMemberMappingAttribute"
	LdapConfigFieldGroupMemberUserAttribute        = "groupMemberUserAttribute"
	LdapConfigFieldGroupNameAttribute              = "groupNameAttribute"
	LdapConfigFieldGroupObjectClass                = "groupObjectClass"
	LdapConfigFieldGroupSearchAttribute            = "groupSearchAttribute"
	LdapConfigFieldGroupSearchBase                 = "groupSearchBase"
	LdapConfigFieldLabels                          = "labels"
	LdapConfigFieldName                            = "name"
	LdapConfigFieldNestedGroupMembershipEnabled    = "nestedGroupMembershipEnabled"
	LdapConfigFieldOwnerReferences                 = "ownerReferences"
	LdapConfigFieldPort                            = "port"
	LdapConfigFieldRemoved                         = "removed"
	LdapConfigFieldServers                         = "servers"
	LdapConfigFieldServiceAccountDistinguishedName = "serviceAccountDistinguishedName"
	LdapConfigFieldServiceAccountPassword          = "serviceAccountPassword"
	LdapConfigFieldTLS                             = "tls"
	LdapConfigFieldType                            = "type"
	LdapConfigFieldUUID                            = "uuid"
	LdapConfigFieldUserDisabledBitMask             = "userDisabledBitMask"
	LdapConfigFieldUserEnabledAttribute            = "userEnabledAttribute"
	LdapConfigFieldUserLoginAttribute              = "userLoginAttribute"
	LdapConfigFieldUserMemberAttribute             = "userMemberAttribute"
	LdapConfigFieldUserNameAttribute               = "userNameAttribute"
	LdapConfigFieldUserObjectClass                 = "userObjectClass"
	LdapConfigFieldUserSearchAttribute             = "userSearchAttribute"
	LdapConfigFieldUserSearchBase                  = "userSearchBase"
)

type LdapConfig struct {
	types.Resource
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

type LdapConfigCollection struct {
	types.Collection
	Data   []LdapConfig `json:"data,omitempty"`
	client *LdapConfigClient
}

type LdapConfigClient struct {
	apiClient *Client
}

type LdapConfigOperations interface {
	List(opts *types.ListOpts) (*LdapConfigCollection, error)
	Create(opts *LdapConfig) (*LdapConfig, error)
	Update(existing *LdapConfig, updates interface{}) (*LdapConfig, error)
	Replace(existing *LdapConfig) (*LdapConfig, error)
	ByID(id string) (*LdapConfig, error)
	Delete(container *LdapConfig) error
}

func newLdapConfigClient(apiClient *Client) *LdapConfigClient {
	return &LdapConfigClient{
		apiClient: apiClient,
	}
}

func (c *LdapConfigClient) Create(container *LdapConfig) (*LdapConfig, error) {
	resp := &LdapConfig{}
	err := c.apiClient.Ops.DoCreate(LdapConfigType, container, resp)
	return resp, err
}

func (c *LdapConfigClient) Update(existing *LdapConfig, updates interface{}) (*LdapConfig, error) {
	resp := &LdapConfig{}
	err := c.apiClient.Ops.DoUpdate(LdapConfigType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LdapConfigClient) Replace(obj *LdapConfig) (*LdapConfig, error) {
	resp := &LdapConfig{}
	err := c.apiClient.Ops.DoReplace(LdapConfigType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *LdapConfigClient) List(opts *types.ListOpts) (*LdapConfigCollection, error) {
	resp := &LdapConfigCollection{}
	err := c.apiClient.Ops.DoList(LdapConfigType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LdapConfigCollection) Next() (*LdapConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LdapConfigCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LdapConfigClient) ByID(id string) (*LdapConfig, error) {
	resp := &LdapConfig{}
	err := c.apiClient.Ops.DoByID(LdapConfigType, id, resp)
	return resp, err
}

func (c *LdapConfigClient) Delete(container *LdapConfig) error {
	return c.apiClient.Ops.DoResourceDelete(LdapConfigType, &container.Resource)
}
