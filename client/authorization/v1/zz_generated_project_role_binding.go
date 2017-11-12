package client

import (
	"github.com/rancher/norman/types"
)

const (
	ProjectRoleBindingType                  = "projectRoleBinding"
	ProjectRoleBindingFieldAPIVersion       = "apiVersion"
	ProjectRoleBindingFieldKind             = "kind"
	ProjectRoleBindingFieldObjectMeta       = "objectMeta"
	ProjectRoleBindingFieldProjectName      = "projectName"
	ProjectRoleBindingFieldRoleTemplateName = "roleTemplateName"
	ProjectRoleBindingFieldSubjects         = "subjects"
)

type ProjectRoleBinding struct {
	types.Resource
	APIVersion       string     `json:"apiVersion,omitempty"`
	Kind             string     `json:"kind,omitempty"`
	ObjectMeta       ObjectMeta `json:"objectMeta,omitempty"`
	ProjectName      string     `json:"projectName,omitempty"`
	RoleTemplateName string     `json:"roleTemplateName,omitempty"`
	Subjects         []Subject  `json:"subjects,omitempty"`
}
type ProjectRoleBindingCollection struct {
	types.Collection
	Data   []ProjectRoleBinding `json:"data,omitempty"`
	client *ProjectRoleBindingClient
}

type ProjectRoleBindingClient struct {
	apiClient *Client
}

type ProjectRoleBindingOperations interface {
	List(opts *types.ListOpts) (*ProjectRoleBindingCollection, error)
	Create(opts *ProjectRoleBinding) (*ProjectRoleBinding, error)
	Update(existing *ProjectRoleBinding, updates interface{}) (*ProjectRoleBinding, error)
	ByID(id string) (*ProjectRoleBinding, error)
	Delete(container *ProjectRoleBinding) error
}

func newProjectRoleBindingClient(apiClient *Client) *ProjectRoleBindingClient {
	return &ProjectRoleBindingClient{
		apiClient: apiClient,
	}
}

func (c *ProjectRoleBindingClient) Create(container *ProjectRoleBinding) (*ProjectRoleBinding, error) {
	resp := &ProjectRoleBinding{}
	err := c.apiClient.Ops.DoCreate(ProjectRoleBindingType, container, resp)
	return resp, err
}

func (c *ProjectRoleBindingClient) Update(existing *ProjectRoleBinding, updates interface{}) (*ProjectRoleBinding, error) {
	resp := &ProjectRoleBinding{}
	err := c.apiClient.Ops.DoUpdate(ProjectRoleBindingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectRoleBindingClient) List(opts *types.ListOpts) (*ProjectRoleBindingCollection, error) {
	resp := &ProjectRoleBindingCollection{}
	err := c.apiClient.Ops.DoList(ProjectRoleBindingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectRoleBindingCollection) Next() (*ProjectRoleBindingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectRoleBindingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectRoleBindingClient) ByID(id string) (*ProjectRoleBinding, error) {
	resp := &ProjectRoleBinding{}
	err := c.apiClient.Ops.DoByID(ProjectRoleBindingType, id, resp)
	return resp, err
}

func (c *ProjectRoleBindingClient) Delete(container *ProjectRoleBinding) error {
	return c.apiClient.Ops.DoResourceDelete(ProjectRoleBindingType, &container.Resource)
}
