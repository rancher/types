package client

import (
	"github.com/rancher/norman/types"
)

const (
	TargetType                      = "target"
	TargetFieldAnnotations          = "annotations"
	TargetFieldAnswers              = "answers"
	TargetFieldClusterConfig        = "clusterConfig"
	TargetFieldClusterID            = "clusterId"
	TargetFieldCreated              = "created"
	TargetFieldCreatorID            = "creatorId"
	TargetFieldLabels               = "labels"
	TargetFieldName                 = "name"
	TargetFieldOwnerReferences      = "ownerReferences"
	TargetFieldRemoved              = "removed"
	TargetFieldState                = "state"
	TargetFieldStatus               = "status"
	TargetFieldTransitioning        = "transitioning"
	TargetFieldTransitioningMessage = "transitioningMessage"
	TargetFieldUUID                 = "uuid"
)

type Target struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Answers              map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterConfig        *ClusterConfig    `json:"clusterConfig,omitempty" yaml:"clusterConfig,omitempty"`
	ClusterID            string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string            `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *TargetStatus     `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type TargetCollection struct {
	types.Collection
	Data   []Target `json:"data,omitempty"`
	client *TargetClient
}

type TargetClient struct {
	apiClient *Client
}

type TargetOperations interface {
	List(opts *types.ListOpts) (*TargetCollection, error)
	Create(opts *Target) (*Target, error)
	Update(existing *Target, updates interface{}) (*Target, error)
	Replace(existing *Target) (*Target, error)
	ByID(id string) (*Target, error)
	Delete(container *Target) error
}

func newTargetClient(apiClient *Client) *TargetClient {
	return &TargetClient{
		apiClient: apiClient,
	}
}

func (c *TargetClient) Create(container *Target) (*Target, error) {
	resp := &Target{}
	err := c.apiClient.Ops.DoCreate(TargetType, container, resp)
	return resp, err
}

func (c *TargetClient) Update(existing *Target, updates interface{}) (*Target, error) {
	resp := &Target{}
	err := c.apiClient.Ops.DoUpdate(TargetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TargetClient) Replace(obj *Target) (*Target, error) {
	resp := &Target{}
	err := c.apiClient.Ops.DoReplace(TargetType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *TargetClient) List(opts *types.ListOpts) (*TargetCollection, error) {
	resp := &TargetCollection{}
	err := c.apiClient.Ops.DoList(TargetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TargetCollection) Next() (*TargetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TargetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TargetClient) ByID(id string) (*Target, error) {
	resp := &Target{}
	err := c.apiClient.Ops.DoByID(TargetType, id, resp)
	return resp, err
}

func (c *TargetClient) Delete(container *Target) error {
	return c.apiClient.Ops.DoResourceDelete(TargetType, &container.Resource)
}
