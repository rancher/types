package client

import (
	"github.com/rancher/norman/types"
)

const (
	DaemonSetType                      = "daemonSet"
	DaemonSetFieldAnnotations          = "annotations"
	DaemonSetFieldCreated              = "created"
	DaemonSetFieldCreatorID            = "creatorId"
	DaemonSetFieldDaemonSet            = "daemonSet"
	DaemonSetFieldDaemonSetStatus      = "daemonSetStatus"
	DaemonSetFieldLabels               = "labels"
	DaemonSetFieldName                 = "name"
	DaemonSetFieldNamespaceId          = "namespaceId"
	DaemonSetFieldOwnerReferences      = "ownerReferences"
	DaemonSetFieldProjectID            = "projectId"
	DaemonSetFieldRemoved              = "removed"
	DaemonSetFieldSelector             = "selector"
	DaemonSetFieldState                = "state"
	DaemonSetFieldTemplate             = "template"
	DaemonSetFieldTransitioning        = "transitioning"
	DaemonSetFieldTransitioningMessage = "transitioningMessage"
	DaemonSetFieldUuid                 = "uuid"
)

type DaemonSet struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty"`
	Created              string            `json:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty"`
	DaemonSet            *DaemonSetConfig  `json:"daemonSet,omitempty"`
	DaemonSetStatus      *DaemonSetStatus  `json:"daemonSetStatus,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	Name                 string            `json:"name,omitempty"`
	NamespaceId          string            `json:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty"`
	ProjectID            string            `json:"projectId,omitempty"`
	Removed              string            `json:"removed,omitempty"`
	Selector             *LabelSelector    `json:"selector,omitempty"`
	State                string            `json:"state,omitempty"`
	Template             *PodTemplateSpec  `json:"template,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty"`
	Uuid                 string            `json:"uuid,omitempty"`
}
type DaemonSetCollection struct {
	types.Collection
	Data   []DaemonSet `json:"data,omitempty"`
	client *DaemonSetClient
}

type DaemonSetClient struct {
	apiClient *Client
}

type DaemonSetOperations interface {
	List(opts *types.ListOpts) (*DaemonSetCollection, error)
	Create(opts *DaemonSet) (*DaemonSet, error)
	Update(existing *DaemonSet, updates interface{}) (*DaemonSet, error)
	ByID(id string) (*DaemonSet, error)
	Delete(container *DaemonSet) error
}

func newDaemonSetClient(apiClient *Client) *DaemonSetClient {
	return &DaemonSetClient{
		apiClient: apiClient,
	}
}

func (c *DaemonSetClient) Create(container *DaemonSet) (*DaemonSet, error) {
	resp := &DaemonSet{}
	err := c.apiClient.Ops.DoCreate(DaemonSetType, container, resp)
	return resp, err
}

func (c *DaemonSetClient) Update(existing *DaemonSet, updates interface{}) (*DaemonSet, error) {
	resp := &DaemonSet{}
	err := c.apiClient.Ops.DoUpdate(DaemonSetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DaemonSetClient) List(opts *types.ListOpts) (*DaemonSetCollection, error) {
	resp := &DaemonSetCollection{}
	err := c.apiClient.Ops.DoList(DaemonSetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DaemonSetCollection) Next() (*DaemonSetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DaemonSetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DaemonSetClient) ByID(id string) (*DaemonSet, error) {
	resp := &DaemonSet{}
	err := c.apiClient.Ops.DoByID(DaemonSetType, id, resp)
	return resp, err
}

func (c *DaemonSetClient) Delete(container *DaemonSet) error {
	return c.apiClient.Ops.DoResourceDelete(DaemonSetType, &container.Resource)
}
