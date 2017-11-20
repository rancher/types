package client

import (
	"github.com/rancher/norman/types"
)

const (
	ProjectRoleTemplateBindingType                         = "projectRoleTemplateBinding"
	ProjectRoleTemplateBindingFieldAPIVersion              = "apiVersion"
	ProjectRoleTemplateBindingFieldKind                    = "kind"
	ProjectRoleTemplateBindingFieldObjectMeta              = "objectMeta"
	ProjectRoleTemplateBindingFieldProjectName             = "projectName"
	ProjectRoleTemplateBindingFieldProjectRoleTemplateName = "projectRoleTemplateName"
	ProjectRoleTemplateBindingFieldSubject                 = "subject"
)

type ProjectRoleTemplateBinding struct {
	types.Resource
	APIVersion              string     `json:"apiVersion,omitempty"`
	Kind                    string     `json:"kind,omitempty"`
	ObjectMeta              ObjectMeta `json:"objectMeta,omitempty"`
	ProjectName             string     `json:"projectName,omitempty"`
	ProjectRoleTemplateName string     `json:"projectRoleTemplateName,omitempty"`
	Subject                 Subject    `json:"subject,omitempty"`
}
type ProjectRoleTemplateBindingCollection struct {
	types.Collection
	Data   []ProjectRoleTemplateBinding `json:"data,omitempty"`
	client *ProjectRoleTemplateBindingClient
}

type ProjectRoleTemplateBindingClient struct {
	apiClient *Client
}

type ProjectRoleTemplateBindingOperations interface {
	List(opts *types.ListOpts) (*ProjectRoleTemplateBindingCollection, error)
	Create(opts *ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error)
	Update(existing *ProjectRoleTemplateBinding, updates interface{}) (*ProjectRoleTemplateBinding, error)
	ByID(id string) (*ProjectRoleTemplateBinding, error)
	Delete(container *ProjectRoleTemplateBinding) error
}

func newProjectRoleTemplateBindingClient(apiClient *Client) *ProjectRoleTemplateBindingClient {
	return &ProjectRoleTemplateBindingClient{
		apiClient: apiClient,
	}
}

func (c *ProjectRoleTemplateBindingClient) Create(container *ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error) {
	resp := &ProjectRoleTemplateBinding{}
	err := c.apiClient.Ops.DoCreate(ProjectRoleTemplateBindingType, container, resp)
	return resp, err
}

func (c *ProjectRoleTemplateBindingClient) Update(existing *ProjectRoleTemplateBinding, updates interface{}) (*ProjectRoleTemplateBinding, error) {
	resp := &ProjectRoleTemplateBinding{}
	err := c.apiClient.Ops.DoUpdate(ProjectRoleTemplateBindingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectRoleTemplateBindingClient) List(opts *types.ListOpts) (*ProjectRoleTemplateBindingCollection, error) {
	resp := &ProjectRoleTemplateBindingCollection{}
	err := c.apiClient.Ops.DoList(ProjectRoleTemplateBindingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectRoleTemplateBindingCollection) Next() (*ProjectRoleTemplateBindingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectRoleTemplateBindingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectRoleTemplateBindingClient) ByID(id string) (*ProjectRoleTemplateBinding, error) {
	resp := &ProjectRoleTemplateBinding{}
	err := c.apiClient.Ops.DoByID(ProjectRoleTemplateBindingType, id, resp)
	return resp, err
}

func (c *ProjectRoleTemplateBindingClient) Delete(container *ProjectRoleTemplateBinding) error {
	return c.apiClient.Ops.DoResourceDelete(ProjectRoleTemplateBindingType, &container.Resource)
}
