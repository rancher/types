package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterRoleTemplateType                      = "clusterRoleTemplate"
	ClusterRoleTemplateFieldAPIVersion           = "apiVersion"
	ClusterRoleTemplateFieldClusterRoleTemplates = "clusterRoleTemplates"
	ClusterRoleTemplateFieldKind                 = "kind"
	ClusterRoleTemplateFieldObjectMeta           = "objectMeta"
	ClusterRoleTemplateFieldRules                = "rules"
)

type ClusterRoleTemplate struct {
	types.Resource
	APIVersion           string       `json:"apiVersion,omitempty"`
	ClusterRoleTemplates []string     `json:"clusterRoleTemplates,omitempty"`
	Kind                 string       `json:"kind,omitempty"`
	ObjectMeta           ObjectMeta   `json:"objectMeta,omitempty"`
	Rules                []PolicyRule `json:"rules,omitempty"`
}
type ClusterRoleTemplateCollection struct {
	types.Collection
	Data   []ClusterRoleTemplate `json:"data,omitempty"`
	client *ClusterRoleTemplateClient
}

type ClusterRoleTemplateClient struct {
	apiClient *Client
}

type ClusterRoleTemplateOperations interface {
	List(opts *types.ListOpts) (*ClusterRoleTemplateCollection, error)
	Create(opts *ClusterRoleTemplate) (*ClusterRoleTemplate, error)
	Update(existing *ClusterRoleTemplate, updates interface{}) (*ClusterRoleTemplate, error)
	ByID(id string) (*ClusterRoleTemplate, error)
	Delete(container *ClusterRoleTemplate) error
}

func newClusterRoleTemplateClient(apiClient *Client) *ClusterRoleTemplateClient {
	return &ClusterRoleTemplateClient{
		apiClient: apiClient,
	}
}

func (c *ClusterRoleTemplateClient) Create(container *ClusterRoleTemplate) (*ClusterRoleTemplate, error) {
	resp := &ClusterRoleTemplate{}
	err := c.apiClient.Ops.DoCreate(ClusterRoleTemplateType, container, resp)
	return resp, err
}

func (c *ClusterRoleTemplateClient) Update(existing *ClusterRoleTemplate, updates interface{}) (*ClusterRoleTemplate, error) {
	resp := &ClusterRoleTemplate{}
	err := c.apiClient.Ops.DoUpdate(ClusterRoleTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterRoleTemplateClient) List(opts *types.ListOpts) (*ClusterRoleTemplateCollection, error) {
	resp := &ClusterRoleTemplateCollection{}
	err := c.apiClient.Ops.DoList(ClusterRoleTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterRoleTemplateCollection) Next() (*ClusterRoleTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterRoleTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterRoleTemplateClient) ByID(id string) (*ClusterRoleTemplate, error) {
	resp := &ClusterRoleTemplate{}
	err := c.apiClient.Ops.DoByID(ClusterRoleTemplateType, id, resp)
	return resp, err
}

func (c *ClusterRoleTemplateClient) Delete(container *ClusterRoleTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterRoleTemplateType, &container.Resource)
}
