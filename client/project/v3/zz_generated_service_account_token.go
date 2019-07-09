package client

import (
	"github.com/rancher/norman/types"
)

const (
	ServiceAccountTokenType                 = "serviceAccountToken"
	ServiceAccountTokenFieldAccountName     = "accountName"
	ServiceAccountTokenFieldAccountUID      = "accountUid"
	ServiceAccountTokenFieldAnnotations     = "annotations"
	ServiceAccountTokenFieldCACRT           = "caCrt"
	ServiceAccountTokenFieldCreated         = "created"
	ServiceAccountTokenFieldCreatorID       = "creatorId"
	ServiceAccountTokenFieldDescription     = "description"
	ServiceAccountTokenFieldLabels          = "labels"
	ServiceAccountTokenFieldName            = "name"
	ServiceAccountTokenFieldNamespaceId     = "namespaceId"
	ServiceAccountTokenFieldOwnerReferences = "ownerReferences"
	ServiceAccountTokenFieldProjectID       = "projectId"
	ServiceAccountTokenFieldRemoved         = "removed"
	ServiceAccountTokenFieldToken           = "token"
	ServiceAccountTokenFieldUUID            = "uuid"
)

type ServiceAccountToken struct {
	types.Resource
	AccountName     string            `json:"accountName,omitempty" yaml:"account_name,omitempty"`
	AccountUID      string            `json:"accountUid,omitempty" yaml:"account_uid,omitempty"`
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	CACRT           string            `json:"caCrt,omitempty" yaml:"ca_crt,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Description     string            `json:"description,omitempty" yaml:"description,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	ProjectID       string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Token           string            `json:"token,omitempty" yaml:"token,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ServiceAccountTokenCollection struct {
	types.Collection
	Data   []ServiceAccountToken `json:"data,omitempty"`
	client *ServiceAccountTokenClient
}

type ServiceAccountTokenClient struct {
	apiClient *Client
}

type ServiceAccountTokenOperations interface {
	List(opts *types.ListOpts) (*ServiceAccountTokenCollection, error)
	Create(opts *ServiceAccountToken) (*ServiceAccountToken, error)
	Update(existing *ServiceAccountToken, updates interface{}) (*ServiceAccountToken, error)
	Replace(existing *ServiceAccountToken) (*ServiceAccountToken, error)
	ByID(id string) (*ServiceAccountToken, error)
	Delete(container *ServiceAccountToken) error
}

func newServiceAccountTokenClient(apiClient *Client) *ServiceAccountTokenClient {
	return &ServiceAccountTokenClient{
		apiClient: apiClient,
	}
}

func (c *ServiceAccountTokenClient) Create(container *ServiceAccountToken) (*ServiceAccountToken, error) {
	resp := &ServiceAccountToken{}
	err := c.apiClient.Ops.DoCreate(ServiceAccountTokenType, container, resp)
	return resp, err
}

func (c *ServiceAccountTokenClient) Update(existing *ServiceAccountToken, updates interface{}) (*ServiceAccountToken, error) {
	resp := &ServiceAccountToken{}
	err := c.apiClient.Ops.DoUpdate(ServiceAccountTokenType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceAccountTokenClient) Replace(obj *ServiceAccountToken) (*ServiceAccountToken, error) {
	resp := &ServiceAccountToken{}
	err := c.apiClient.Ops.DoReplace(ServiceAccountTokenType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ServiceAccountTokenClient) List(opts *types.ListOpts) (*ServiceAccountTokenCollection, error) {
	resp := &ServiceAccountTokenCollection{}
	err := c.apiClient.Ops.DoList(ServiceAccountTokenType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceAccountTokenCollection) Next() (*ServiceAccountTokenCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceAccountTokenCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceAccountTokenClient) ByID(id string) (*ServiceAccountToken, error) {
	resp := &ServiceAccountToken{}
	err := c.apiClient.Ops.DoByID(ServiceAccountTokenType, id, resp)
	return resp, err
}

func (c *ServiceAccountTokenClient) Delete(container *ServiceAccountToken) error {
	return c.apiClient.Ops.DoResourceDelete(ServiceAccountTokenType, &container.Resource)
}
