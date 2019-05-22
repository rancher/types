package client

import (
	"github.com/rancher/norman/types"
)

const (
	ExampleConfigType                      = "exampleConfig"
	ExampleConfigFieldAnnotations          = "annotations"
	ExampleConfigFieldCreated              = "created"
	ExampleConfigFieldCreatorID            = "creatorId"
	ExampleConfigFieldExampleString        = "rancherCompose"
	ExampleConfigFieldLabels               = "labels"
	ExampleConfigFieldName                 = "name"
	ExampleConfigFieldOwnerReferences      = "ownerReferences"
	ExampleConfigFieldRemoved              = "removed"
	ExampleConfigFieldState                = "state"
	ExampleConfigFieldStatus               = "status"
	ExampleConfigFieldTransitioning        = "transitioning"
	ExampleConfigFieldTransitioningMessage = "transitioningMessage"
	ExampleConfigFieldUUID                 = "uuid"
)

type ExampleConfig struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	ExampleString        string            `json:"rancherCompose,omitempty" yaml:"rancherCompose,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string            `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *ExampleStatus    `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExampleConfigCollection struct {
	types.Collection
	Data   []ExampleConfig `json:"data,omitempty"`
	client *ExampleConfigClient
}

type ExampleConfigClient struct {
	apiClient *Client
}

type ExampleConfigOperations interface {
	List(opts *types.ListOpts) (*ExampleConfigCollection, error)
	Create(opts *ExampleConfig) (*ExampleConfig, error)
	Update(existing *ExampleConfig, updates interface{}) (*ExampleConfig, error)
	Replace(existing *ExampleConfig) (*ExampleConfig, error)
	ByID(id string) (*ExampleConfig, error)
	Delete(container *ExampleConfig) error
}

func newExampleConfigClient(apiClient *Client) *ExampleConfigClient {
	return &ExampleConfigClient{
		apiClient: apiClient,
	}
}

func (c *ExampleConfigClient) Create(container *ExampleConfig) (*ExampleConfig, error) {
	resp := &ExampleConfig{}
	err := c.apiClient.Ops.DoCreate(ExampleConfigType, container, resp)
	return resp, err
}

func (c *ExampleConfigClient) Update(existing *ExampleConfig, updates interface{}) (*ExampleConfig, error) {
	resp := &ExampleConfig{}
	err := c.apiClient.Ops.DoUpdate(ExampleConfigType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExampleConfigClient) Replace(obj *ExampleConfig) (*ExampleConfig, error) {
	resp := &ExampleConfig{}
	err := c.apiClient.Ops.DoReplace(ExampleConfigType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ExampleConfigClient) List(opts *types.ListOpts) (*ExampleConfigCollection, error) {
	resp := &ExampleConfigCollection{}
	err := c.apiClient.Ops.DoList(ExampleConfigType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExampleConfigCollection) Next() (*ExampleConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExampleConfigCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExampleConfigClient) ByID(id string) (*ExampleConfig, error) {
	resp := &ExampleConfig{}
	err := c.apiClient.Ops.DoByID(ExampleConfigType, id, resp)
	return resp, err
}

func (c *ExampleConfigClient) Delete(container *ExampleConfig) error {
	return c.apiClient.Ops.DoResourceDelete(ExampleConfigType, &container.Resource)
}
