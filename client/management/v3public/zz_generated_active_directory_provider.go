package client

import (
	"github.com/rancher/norman/types"
)

const (
	ActiveDirectoryProviderType                    = "activeDirectoryProvider"
	ActiveDirectoryProviderFieldAnnotations        = "annotations"
	ActiveDirectoryProviderFieldCreated            = "created"
	ActiveDirectoryProviderFieldCreatorID          = "creatorId"
	ActiveDirectoryProviderFieldDefaultLoginDomain = "defaultLoginDomain"
	ActiveDirectoryProviderFieldLabels             = "labels"
	ActiveDirectoryProviderFieldName               = "name"
	ActiveDirectoryProviderFieldOwnerReferences    = "ownerReferences"
	ActiveDirectoryProviderFieldRemoved            = "removed"
	ActiveDirectoryProviderFieldType               = "type"
	ActiveDirectoryProviderFieldUUID               = "uuid"
)

type ActiveDirectoryProvider struct {
	types.Resource
	Annotations        map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created            string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID          string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	DefaultLoginDomain string            `json:"defaultLoginDomain,omitempty" yaml:"defaultLoginDomain,omitempty"`
	Labels             map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name               string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences    []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed            string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Type               string            `json:"type,omitempty" yaml:"type,omitempty"`
	UUID               string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ActiveDirectoryProviderCollection struct {
	types.Collection
	Data   []ActiveDirectoryProvider `json:"data,omitempty"`
	client *ActiveDirectoryProviderClient
}

type ActiveDirectoryProviderClient struct {
	apiClient *Client
}

type ActiveDirectoryProviderOperations interface {
	List(opts *types.ListOpts) (*ActiveDirectoryProviderCollection, error)
	Create(opts *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	Update(existing *ActiveDirectoryProvider, updates interface{}) (*ActiveDirectoryProvider, error)
	Replace(existing *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	ByID(id string) (*ActiveDirectoryProvider, error)
	Delete(container *ActiveDirectoryProvider) error

	ActionLogin(resource *ActiveDirectoryProvider, input *BasicLogin) (*Token, error)
}

func newActiveDirectoryProviderClient(apiClient *Client) *ActiveDirectoryProviderClient {
	return &ActiveDirectoryProviderClient{
		apiClient: apiClient,
	}
}

func (c *ActiveDirectoryProviderClient) Create(container *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	resp := &ActiveDirectoryProvider{}
	err := c.apiClient.Ops.DoCreate(ActiveDirectoryProviderType, container, resp)
	return resp, err
}

func (c *ActiveDirectoryProviderClient) Update(existing *ActiveDirectoryProvider, updates interface{}) (*ActiveDirectoryProvider, error) {
	resp := &ActiveDirectoryProvider{}
	err := c.apiClient.Ops.DoUpdate(ActiveDirectoryProviderType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ActiveDirectoryProviderClient) Replace(obj *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	resp := &ActiveDirectoryProvider{}
	err := c.apiClient.Ops.DoReplace(ActiveDirectoryProviderType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ActiveDirectoryProviderClient) List(opts *types.ListOpts) (*ActiveDirectoryProviderCollection, error) {
	resp := &ActiveDirectoryProviderCollection{}
	err := c.apiClient.Ops.DoList(ActiveDirectoryProviderType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ActiveDirectoryProviderCollection) Next() (*ActiveDirectoryProviderCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ActiveDirectoryProviderCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ActiveDirectoryProviderClient) ByID(id string) (*ActiveDirectoryProvider, error) {
	resp := &ActiveDirectoryProvider{}
	err := c.apiClient.Ops.DoByID(ActiveDirectoryProviderType, id, resp)
	return resp, err
}

func (c *ActiveDirectoryProviderClient) Delete(container *ActiveDirectoryProvider) error {
	return c.apiClient.Ops.DoResourceDelete(ActiveDirectoryProviderType, &container.Resource)
}

func (c *ActiveDirectoryProviderClient) ActionLogin(resource *ActiveDirectoryProvider, input *BasicLogin) (*Token, error) {
	resp := &Token{}
	err := c.apiClient.Ops.DoAction(ActiveDirectoryProviderType, "login", &resource.Resource, input, resp)
	return resp, err
}
