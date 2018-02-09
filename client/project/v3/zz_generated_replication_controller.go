package client

import (
	"github.com/rancher/norman/types"
)

const (
	ReplicationControllerType                             = "replicationController"
	ReplicationControllerFieldAnnotations                 = "annotations"
	ReplicationControllerFieldCreated                     = "created"
	ReplicationControllerFieldCreatorID                   = "creatorId"
	ReplicationControllerFieldLabels                      = "labels"
	ReplicationControllerFieldName                        = "name"
	ReplicationControllerFieldNamespaceId                 = "namespaceId"
	ReplicationControllerFieldOwnerReferences             = "ownerReferences"
	ReplicationControllerFieldProjectID                   = "projectId"
	ReplicationControllerFieldRemoved                     = "removed"
	ReplicationControllerFieldReplicationController       = "replicationController"
	ReplicationControllerFieldReplicationControllerStatus = "replicationControllerStatus"
	ReplicationControllerFieldSelector                    = "selector"
	ReplicationControllerFieldState                       = "state"
	ReplicationControllerFieldTemplate                    = "template"
	ReplicationControllerFieldTransitioning               = "transitioning"
	ReplicationControllerFieldTransitioningMessage        = "transitioningMessage"
	ReplicationControllerFieldUuid                        = "uuid"
)

type ReplicationController struct {
	types.Resource
	Annotations                 map[string]string            `json:"annotations,omitempty"`
	Created                     string                       `json:"created,omitempty"`
	CreatorID                   string                       `json:"creatorId,omitempty"`
	Labels                      map[string]string            `json:"labels,omitempty"`
	Name                        string                       `json:"name,omitempty"`
	NamespaceId                 string                       `json:"namespaceId,omitempty"`
	OwnerReferences             []OwnerReference             `json:"ownerReferences,omitempty"`
	ProjectID                   string                       `json:"projectId,omitempty"`
	Removed                     string                       `json:"removed,omitempty"`
	ReplicationController       *ReplicationControllerConfig `json:"replicationController,omitempty"`
	ReplicationControllerStatus *ReplicationControllerStatus `json:"replicationControllerStatus,omitempty"`
	Selector                    map[string]string            `json:"selector,omitempty"`
	State                       string                       `json:"state,omitempty"`
	Template                    *PodTemplateSpec             `json:"template,omitempty"`
	Transitioning               string                       `json:"transitioning,omitempty"`
	TransitioningMessage        string                       `json:"transitioningMessage,omitempty"`
	Uuid                        string                       `json:"uuid,omitempty"`
}
type ReplicationControllerCollection struct {
	types.Collection
	Data   []ReplicationController `json:"data,omitempty"`
	client *ReplicationControllerClient
}

type ReplicationControllerClient struct {
	apiClient *Client
}

type ReplicationControllerOperations interface {
	List(opts *types.ListOpts) (*ReplicationControllerCollection, error)
	Create(opts *ReplicationController) (*ReplicationController, error)
	Update(existing *ReplicationController, updates interface{}) (*ReplicationController, error)
	ByID(id string) (*ReplicationController, error)
	Delete(container *ReplicationController) error
}

func newReplicationControllerClient(apiClient *Client) *ReplicationControllerClient {
	return &ReplicationControllerClient{
		apiClient: apiClient,
	}
}

func (c *ReplicationControllerClient) Create(container *ReplicationController) (*ReplicationController, error) {
	resp := &ReplicationController{}
	err := c.apiClient.Ops.DoCreate(ReplicationControllerType, container, resp)
	return resp, err
}

func (c *ReplicationControllerClient) Update(existing *ReplicationController, updates interface{}) (*ReplicationController, error) {
	resp := &ReplicationController{}
	err := c.apiClient.Ops.DoUpdate(ReplicationControllerType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ReplicationControllerClient) List(opts *types.ListOpts) (*ReplicationControllerCollection, error) {
	resp := &ReplicationControllerCollection{}
	err := c.apiClient.Ops.DoList(ReplicationControllerType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ReplicationControllerCollection) Next() (*ReplicationControllerCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ReplicationControllerCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ReplicationControllerClient) ByID(id string) (*ReplicationController, error) {
	resp := &ReplicationController{}
	err := c.apiClient.Ops.DoByID(ReplicationControllerType, id, resp)
	return resp, err
}

func (c *ReplicationControllerClient) Delete(container *ReplicationController) error {
	return c.apiClient.Ops.DoResourceDelete(ReplicationControllerType, &container.Resource)
}
