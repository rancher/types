package client

import (
	"github.com/rancher/norman/types"
)

const (
	PhabricatorConfigType                     = "phabricatorConfig"
	PhabricatorConfigFieldAccessMode          = "accessMode"
	PhabricatorConfigFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	PhabricatorConfigFieldAnnotations         = "annotations"
	PhabricatorConfigFieldClientID            = "clientId"
	PhabricatorConfigFieldClientSecret        = "clientSecret"
	PhabricatorConfigFieldCreated             = "created"
	PhabricatorConfigFieldCreatorID           = "creatorId"
	PhabricatorConfigFieldEnabled             = "enabled"
	PhabricatorConfigFieldHostname            = "hostname"
	PhabricatorConfigFieldLabels              = "labels"
	PhabricatorConfigFieldName                = "name"
	PhabricatorConfigFieldOwnerReferences     = "ownerReferences"
	PhabricatorConfigFieldRemoved             = "removed"
	PhabricatorConfigFieldTLS                 = "tls"
	PhabricatorConfigFieldType                = "type"
	PhabricatorConfigFieldUuid                = "uuid"
)

type PhabricatorConfig struct {
	types.Resource
	AccessMode          string            `json:"accessMode,omitempty" yaml:"accessMode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowedPrincipalIds,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClientID            string            `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	ClientSecret        string            `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Hostname            string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TLS                 bool              `json:"tls,omitempty" yaml:"tls,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
	Uuid                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type PhabricatorConfigCollection struct {
	types.Collection
	Data   []PhabricatorConfig `json:"data,omitempty"`
	client *PhabricatorConfigClient
}

type PhabricatorConfigClient struct {
	apiClient *Client
}

type PhabricatorConfigOperations interface {
	List(opts *types.ListOpts) (*PhabricatorConfigCollection, error)
	Create(opts *PhabricatorConfig) (*PhabricatorConfig, error)
	Update(existing *PhabricatorConfig, updates interface{}) (*PhabricatorConfig, error)
	ByID(id string) (*PhabricatorConfig, error)
	Delete(container *PhabricatorConfig) error
}

func newPhabricatorConfigClient(apiClient *Client) *PhabricatorConfigClient {
	return &PhabricatorConfigClient{
		apiClient: apiClient,
	}
}

func (c *PhabricatorConfigClient) Create(container *PhabricatorConfig) (*PhabricatorConfig, error) {
	resp := &PhabricatorConfig{}
	err := c.apiClient.Ops.DoCreate(PhabricatorConfigType, container, resp)
	return resp, err
}

func (c *PhabricatorConfigClient) Update(existing *PhabricatorConfig, updates interface{}) (*PhabricatorConfig, error) {
	resp := &PhabricatorConfig{}
	err := c.apiClient.Ops.DoUpdate(PhabricatorConfigType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PhabricatorConfigClient) List(opts *types.ListOpts) (*PhabricatorConfigCollection, error) {
	resp := &PhabricatorConfigCollection{}
	err := c.apiClient.Ops.DoList(PhabricatorConfigType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PhabricatorConfigCollection) Next() (*PhabricatorConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PhabricatorConfigCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PhabricatorConfigClient) ByID(id string) (*PhabricatorConfig, error) {
	resp := &PhabricatorConfig{}
	err := c.apiClient.Ops.DoByID(PhabricatorConfigType, id, resp)
	return resp, err
}

func (c *PhabricatorConfigClient) Delete(container *PhabricatorConfig) error {
	return c.apiClient.Ops.DoResourceDelete(PhabricatorConfigType, &container.Resource)
}
