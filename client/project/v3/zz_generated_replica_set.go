package client

import (
	"github.com/rancher/norman/types"
)

const (
	ReplicaSetType                      = "replicaSet"
	ReplicaSetFieldAnnotations          = "annotations"
	ReplicaSetFieldCreated              = "created"
	ReplicaSetFieldCreatorID            = "creatorId"
	ReplicaSetFieldLabels               = "labels"
	ReplicaSetFieldName                 = "name"
	ReplicaSetFieldNamespaceId          = "namespaceId"
	ReplicaSetFieldOwnerReferences      = "ownerReferences"
	ReplicaSetFieldProjectID            = "projectId"
	ReplicaSetFieldRemoved              = "removed"
	ReplicaSetFieldReplicaSet           = "replicaSet"
	ReplicaSetFieldReplicaSetStatus     = "replicaSetStatus"
	ReplicaSetFieldSelector             = "selector"
	ReplicaSetFieldState                = "state"
	ReplicaSetFieldTemplate             = "template"
	ReplicaSetFieldTransitioning        = "transitioning"
	ReplicaSetFieldTransitioningMessage = "transitioningMessage"
	ReplicaSetFieldUuid                 = "uuid"
)

type ReplicaSet struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty"`
	Created              string            `json:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	Name                 string            `json:"name,omitempty"`
	NamespaceId          string            `json:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty"`
	ProjectID            string            `json:"projectId,omitempty"`
	Removed              string            `json:"removed,omitempty"`
	ReplicaSet           *ReplicaSetConfig `json:"replicaSet,omitempty"`
	ReplicaSetStatus     *ReplicaSetStatus `json:"replicaSetStatus,omitempty"`
	Selector             *LabelSelector    `json:"selector,omitempty"`
	State                string            `json:"state,omitempty"`
	Template             *PodTemplateSpec  `json:"template,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty"`
	Uuid                 string            `json:"uuid,omitempty"`
}
type ReplicaSetCollection struct {
	types.Collection
	Data   []ReplicaSet `json:"data,omitempty"`
	client *ReplicaSetClient
}

type ReplicaSetClient struct {
	apiClient *Client
}

type ReplicaSetOperations interface {
	List(opts *types.ListOpts) (*ReplicaSetCollection, error)
	Create(opts *ReplicaSet) (*ReplicaSet, error)
	Update(existing *ReplicaSet, updates interface{}) (*ReplicaSet, error)
	ByID(id string) (*ReplicaSet, error)
	Delete(container *ReplicaSet) error
}

func newReplicaSetClient(apiClient *Client) *ReplicaSetClient {
	return &ReplicaSetClient{
		apiClient: apiClient,
	}
}

func (c *ReplicaSetClient) Create(container *ReplicaSet) (*ReplicaSet, error) {
	resp := &ReplicaSet{}
	err := c.apiClient.Ops.DoCreate(ReplicaSetType, container, resp)
	return resp, err
}

func (c *ReplicaSetClient) Update(existing *ReplicaSet, updates interface{}) (*ReplicaSet, error) {
	resp := &ReplicaSet{}
	err := c.apiClient.Ops.DoUpdate(ReplicaSetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ReplicaSetClient) List(opts *types.ListOpts) (*ReplicaSetCollection, error) {
	resp := &ReplicaSetCollection{}
	err := c.apiClient.Ops.DoList(ReplicaSetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ReplicaSetCollection) Next() (*ReplicaSetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ReplicaSetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ReplicaSetClient) ByID(id string) (*ReplicaSet, error) {
	resp := &ReplicaSet{}
	err := c.apiClient.Ops.DoByID(ReplicaSetType, id, resp)
	return resp, err
}

func (c *ReplicaSetClient) Delete(container *ReplicaSet) error {
	return c.apiClient.Ops.DoResourceDelete(ReplicaSetType, &container.Resource)
}
