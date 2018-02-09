package client

import (
	"github.com/rancher/norman/types"
)

const (
	DeploymentType                      = "deployment"
	DeploymentFieldAnnotations          = "annotations"
	DeploymentFieldCreated              = "created"
	DeploymentFieldCreatorID            = "creatorId"
	DeploymentFieldDeployment           = "deployment"
	DeploymentFieldDeploymentStatus     = "deploymentStatus"
	DeploymentFieldLabels               = "labels"
	DeploymentFieldName                 = "name"
	DeploymentFieldNamespaceId          = "namespaceId"
	DeploymentFieldOwnerReferences      = "ownerReferences"
	DeploymentFieldProjectID            = "projectId"
	DeploymentFieldRemoved              = "removed"
	DeploymentFieldScale                = "scale"
	DeploymentFieldSelector             = "selector"
	DeploymentFieldState                = "state"
	DeploymentFieldTemplate             = "template"
	DeploymentFieldTransitioning        = "transitioning"
	DeploymentFieldTransitioningMessage = "transitioningMessage"
	DeploymentFieldUuid                 = "uuid"
)

type Deployment struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty"`
	Created              string            `json:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty"`
	Deployment           *DeploymentConfig `json:"deployment,omitempty"`
	DeploymentStatus     *DeploymentStatus `json:"deploymentStatus,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	Name                 string            `json:"name,omitempty"`
	NamespaceId          string            `json:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty"`
	ProjectID            string            `json:"projectId,omitempty"`
	Removed              string            `json:"removed,omitempty"`
	Scale                *int64            `json:"scale,omitempty"`
	Selector             *LabelSelector    `json:"selector,omitempty"`
	State                string            `json:"state,omitempty"`
	Template             *PodTemplateSpec  `json:"template,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty"`
	Uuid                 string            `json:"uuid,omitempty"`
}
type DeploymentCollection struct {
	types.Collection
	Data   []Deployment `json:"data,omitempty"`
	client *DeploymentClient
}

type DeploymentClient struct {
	apiClient *Client
}

type DeploymentOperations interface {
	List(opts *types.ListOpts) (*DeploymentCollection, error)
	Create(opts *Deployment) (*Deployment, error)
	Update(existing *Deployment, updates interface{}) (*Deployment, error)
	ByID(id string) (*Deployment, error)
	Delete(container *Deployment) error
}

func newDeploymentClient(apiClient *Client) *DeploymentClient {
	return &DeploymentClient{
		apiClient: apiClient,
	}
}

func (c *DeploymentClient) Create(container *Deployment) (*Deployment, error) {
	resp := &Deployment{}
	err := c.apiClient.Ops.DoCreate(DeploymentType, container, resp)
	return resp, err
}

func (c *DeploymentClient) Update(existing *Deployment, updates interface{}) (*Deployment, error) {
	resp := &Deployment{}
	err := c.apiClient.Ops.DoUpdate(DeploymentType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DeploymentClient) List(opts *types.ListOpts) (*DeploymentCollection, error) {
	resp := &DeploymentCollection{}
	err := c.apiClient.Ops.DoList(DeploymentType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DeploymentCollection) Next() (*DeploymentCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DeploymentCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DeploymentClient) ByID(id string) (*Deployment, error) {
	resp := &Deployment{}
	err := c.apiClient.Ops.DoByID(DeploymentType, id, resp)
	return resp, err
}

func (c *DeploymentClient) Delete(container *Deployment) error {
	return c.apiClient.Ops.DoResourceDelete(DeploymentType, &container.Resource)
}
