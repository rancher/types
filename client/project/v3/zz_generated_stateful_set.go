package client

import (
	"github.com/rancher/norman/types"
)

const (
	StatefulSetType                      = "statefulSet"
	StatefulSetFieldAnnotations          = "annotations"
	StatefulSetFieldCreated              = "created"
	StatefulSetFieldCreatorID            = "creatorId"
	StatefulSetFieldLabels               = "labels"
	StatefulSetFieldName                 = "name"
	StatefulSetFieldNamespaceId          = "namespaceId"
	StatefulSetFieldOwnerReferences      = "ownerReferences"
	StatefulSetFieldProjectID            = "projectId"
	StatefulSetFieldRemoved              = "removed"
	StatefulSetFieldSelector             = "selector"
	StatefulSetFieldState                = "state"
	StatefulSetFieldStatefulSet          = "statefulSet"
	StatefulSetFieldStatefulSetStatus    = "statefulSetStatus"
	StatefulSetFieldTemplate             = "template"
	StatefulSetFieldTransitioning        = "transitioning"
	StatefulSetFieldTransitioningMessage = "transitioningMessage"
	StatefulSetFieldUuid                 = "uuid"
)

type StatefulSet struct {
	types.Resource
	Annotations          map[string]string  `json:"annotations,omitempty"`
	Created              string             `json:"created,omitempty"`
	CreatorID            string             `json:"creatorId,omitempty"`
	Labels               map[string]string  `json:"labels,omitempty"`
	Name                 string             `json:"name,omitempty"`
	NamespaceId          string             `json:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference   `json:"ownerReferences,omitempty"`
	ProjectID            string             `json:"projectId,omitempty"`
	Removed              string             `json:"removed,omitempty"`
	Selector             *LabelSelector     `json:"selector,omitempty"`
	State                string             `json:"state,omitempty"`
	StatefulSet          *StatefulSetConfig `json:"statefulSet,omitempty"`
	StatefulSetStatus    *StatefulSetStatus `json:"statefulSetStatus,omitempty"`
	Template             *PodTemplateSpec   `json:"template,omitempty"`
	Transitioning        string             `json:"transitioning,omitempty"`
	TransitioningMessage string             `json:"transitioningMessage,omitempty"`
	Uuid                 string             `json:"uuid,omitempty"`
}
type StatefulSetCollection struct {
	types.Collection
	Data   []StatefulSet `json:"data,omitempty"`
	client *StatefulSetClient
}

type StatefulSetClient struct {
	apiClient *Client
}

type StatefulSetOperations interface {
	List(opts *types.ListOpts) (*StatefulSetCollection, error)
	Create(opts *StatefulSet) (*StatefulSet, error)
	Update(existing *StatefulSet, updates interface{}) (*StatefulSet, error)
	ByID(id string) (*StatefulSet, error)
	Delete(container *StatefulSet) error
}

func newStatefulSetClient(apiClient *Client) *StatefulSetClient {
	return &StatefulSetClient{
		apiClient: apiClient,
	}
}

func (c *StatefulSetClient) Create(container *StatefulSet) (*StatefulSet, error) {
	resp := &StatefulSet{}
	err := c.apiClient.Ops.DoCreate(StatefulSetType, container, resp)
	return resp, err
}

func (c *StatefulSetClient) Update(existing *StatefulSet, updates interface{}) (*StatefulSet, error) {
	resp := &StatefulSet{}
	err := c.apiClient.Ops.DoUpdate(StatefulSetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StatefulSetClient) List(opts *types.ListOpts) (*StatefulSetCollection, error) {
	resp := &StatefulSetCollection{}
	err := c.apiClient.Ops.DoList(StatefulSetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *StatefulSetCollection) Next() (*StatefulSetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &StatefulSetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *StatefulSetClient) ByID(id string) (*StatefulSet, error) {
	resp := &StatefulSet{}
	err := c.apiClient.Ops.DoByID(StatefulSetType, id, resp)
	return resp, err
}

func (c *StatefulSetClient) Delete(container *StatefulSet) error {
	return c.apiClient.Ops.DoResourceDelete(StatefulSetType, &container.Resource)
}
