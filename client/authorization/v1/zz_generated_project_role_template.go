package client

import (
	"github.com/rancher/norman/types"
)

const (
	ProjectRoleTemplateType                        = "projectRoleTemplate"
	ProjectRoleTemplateFieldAnnotations            = "annotations"
	ProjectRoleTemplateFieldCreated                = "created"
	ProjectRoleTemplateFieldLabels                 = "labels"
	ProjectRoleTemplateFieldName                   = "name"
	ProjectRoleTemplateFieldOwnerReferences        = "ownerReferences"
	ProjectRoleTemplateFieldProjectRoleTemplateIds = "projectRoleTemplateIds"
	ProjectRoleTemplateFieldRemoved                = "removed"
	ProjectRoleTemplateFieldResourcePath           = "resourcePath"
	ProjectRoleTemplateFieldRules                  = "rules"
	ProjectRoleTemplateFieldUuid                   = "uuid"
)

type ProjectRoleTemplate struct {
	types.Resource
	Annotations            map[string]string `json:"annotations,omitempty"`
	Created                string            `json:"created,omitempty"`
	Labels                 map[string]string `json:"labels,omitempty"`
	Name                   string            `json:"name,omitempty"`
	OwnerReferences        []OwnerReference  `json:"ownerReferences,omitempty"`
	ProjectRoleTemplateIds []string          `json:"projectRoleTemplateIds,omitempty"`
	Removed                string            `json:"removed,omitempty"`
	ResourcePath           string            `json:"resourcePath,omitempty"`
	Rules                  []PolicyRule      `json:"rules,omitempty"`
	Uuid                   string            `json:"uuid,omitempty"`
}
type ProjectRoleTemplateCollection struct {
	types.Collection
	Data   []ProjectRoleTemplate `json:"data,omitempty"`
	client *ProjectRoleTemplateClient
}

type ProjectRoleTemplateClient struct {
	apiClient *Client
}

type ProjectRoleTemplateOperations interface {
	List(opts *types.ListOpts) (*ProjectRoleTemplateCollection, error)
	Create(opts *ProjectRoleTemplate) (*ProjectRoleTemplate, error)
	Update(existing *ProjectRoleTemplate, updates interface{}) (*ProjectRoleTemplate, error)
	ByID(id string) (*ProjectRoleTemplate, error)
	Delete(container *ProjectRoleTemplate) error
}

func newProjectRoleTemplateClient(apiClient *Client) *ProjectRoleTemplateClient {
	return &ProjectRoleTemplateClient{
		apiClient: apiClient,
	}
}

func (c *ProjectRoleTemplateClient) Create(container *ProjectRoleTemplate) (*ProjectRoleTemplate, error) {
	resp := &ProjectRoleTemplate{}
	err := c.apiClient.Ops.DoCreate(ProjectRoleTemplateType, container, resp)
	return resp, err
}

func (c *ProjectRoleTemplateClient) Update(existing *ProjectRoleTemplate, updates interface{}) (*ProjectRoleTemplate, error) {
	resp := &ProjectRoleTemplate{}
	err := c.apiClient.Ops.DoUpdate(ProjectRoleTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectRoleTemplateClient) List(opts *types.ListOpts) (*ProjectRoleTemplateCollection, error) {
	resp := &ProjectRoleTemplateCollection{}
	err := c.apiClient.Ops.DoList(ProjectRoleTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectRoleTemplateCollection) Next() (*ProjectRoleTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectRoleTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectRoleTemplateClient) ByID(id string) (*ProjectRoleTemplate, error) {
	resp := &ProjectRoleTemplate{}
	err := c.apiClient.Ops.DoByID(ProjectRoleTemplateType, id, resp)
	return resp, err
}

func (c *ProjectRoleTemplateClient) Delete(container *ProjectRoleTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(ProjectRoleTemplateType, &container.Resource)
}
