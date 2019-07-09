package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterRoleTemplateBindingType                  = "clusterRoleTemplateBinding"
	ClusterRoleTemplateBindingFieldAnnotations      = "annotations"
	ClusterRoleTemplateBindingFieldClusterID        = "clusterId"
	ClusterRoleTemplateBindingFieldCreated          = "created"
	ClusterRoleTemplateBindingFieldCreatorID        = "creatorId"
	ClusterRoleTemplateBindingFieldGroupID          = "groupId"
	ClusterRoleTemplateBindingFieldGroupPrincipalID = "groupPrincipalId"
	ClusterRoleTemplateBindingFieldLabels           = "labels"
	ClusterRoleTemplateBindingFieldName             = "name"
	ClusterRoleTemplateBindingFieldNamespaceId      = "namespaceId"
	ClusterRoleTemplateBindingFieldOwnerReferences  = "ownerReferences"
	ClusterRoleTemplateBindingFieldRemoved          = "removed"
	ClusterRoleTemplateBindingFieldRoleTemplateID   = "roleTemplateId"
	ClusterRoleTemplateBindingFieldUUID             = "uuid"
	ClusterRoleTemplateBindingFieldUserID           = "userId"
	ClusterRoleTemplateBindingFieldUserPrincipalID  = "userPrincipalId"
)

type ClusterRoleTemplateBinding struct {
	types.Resource
	Annotations      map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterID        string            `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Created          string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID        string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	GroupID          string            `json:"groupId,omitempty" yaml:"group_id,omitempty"`
	GroupPrincipalID string            `json:"groupPrincipalId,omitempty" yaml:"group_principal_id,omitempty"`
	Labels           map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name             string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId      string            `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	OwnerReferences  []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Removed          string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	RoleTemplateID   string            `json:"roleTemplateId,omitempty" yaml:"role_template_id,omitempty"`
	UUID             string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UserID           string            `json:"userId,omitempty" yaml:"user_id,omitempty"`
	UserPrincipalID  string            `json:"userPrincipalId,omitempty" yaml:"user_principal_id,omitempty"`
}

type ClusterRoleTemplateBindingCollection struct {
	types.Collection
	Data   []ClusterRoleTemplateBinding `json:"data,omitempty"`
	client *ClusterRoleTemplateBindingClient
}

type ClusterRoleTemplateBindingClient struct {
	apiClient *Client
}

type ClusterRoleTemplateBindingOperations interface {
	List(opts *types.ListOpts) (*ClusterRoleTemplateBindingCollection, error)
	Create(opts *ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error)
	Update(existing *ClusterRoleTemplateBinding, updates interface{}) (*ClusterRoleTemplateBinding, error)
	Replace(existing *ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error)
	ByID(id string) (*ClusterRoleTemplateBinding, error)
	Delete(container *ClusterRoleTemplateBinding) error
}

func newClusterRoleTemplateBindingClient(apiClient *Client) *ClusterRoleTemplateBindingClient {
	return &ClusterRoleTemplateBindingClient{
		apiClient: apiClient,
	}
}

func (c *ClusterRoleTemplateBindingClient) Create(container *ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error) {
	resp := &ClusterRoleTemplateBinding{}
	err := c.apiClient.Ops.DoCreate(ClusterRoleTemplateBindingType, container, resp)
	return resp, err
}

func (c *ClusterRoleTemplateBindingClient) Update(existing *ClusterRoleTemplateBinding, updates interface{}) (*ClusterRoleTemplateBinding, error) {
	resp := &ClusterRoleTemplateBinding{}
	err := c.apiClient.Ops.DoUpdate(ClusterRoleTemplateBindingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterRoleTemplateBindingClient) Replace(obj *ClusterRoleTemplateBinding) (*ClusterRoleTemplateBinding, error) {
	resp := &ClusterRoleTemplateBinding{}
	err := c.apiClient.Ops.DoReplace(ClusterRoleTemplateBindingType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterRoleTemplateBindingClient) List(opts *types.ListOpts) (*ClusterRoleTemplateBindingCollection, error) {
	resp := &ClusterRoleTemplateBindingCollection{}
	err := c.apiClient.Ops.DoList(ClusterRoleTemplateBindingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterRoleTemplateBindingCollection) Next() (*ClusterRoleTemplateBindingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterRoleTemplateBindingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterRoleTemplateBindingClient) ByID(id string) (*ClusterRoleTemplateBinding, error) {
	resp := &ClusterRoleTemplateBinding{}
	err := c.apiClient.Ops.DoByID(ClusterRoleTemplateBindingType, id, resp)
	return resp, err
}

func (c *ClusterRoleTemplateBindingClient) Delete(container *ClusterRoleTemplateBinding) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterRoleTemplateBindingType, &container.Resource)
}
