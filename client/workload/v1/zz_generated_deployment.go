package client

import (
	"github.com/rancher/norman/types"
)

const (
	DeploymentType                   = "deployment"
	DeploymentFieldAPIVersion        = "apiVersion"
	DeploymentFieldAnnotations       = "annotations"
	DeploymentFieldCreationTimestamp = "creationTimestamp"
	DeploymentFieldDeletionTimestamp = "deletionTimestamp"
	DeploymentFieldDeploy            = "deploy"
	DeploymentFieldKind              = "kind"
	DeploymentFieldLabels            = "labels"
	DeploymentFieldName              = "name"
	DeploymentFieldNamespace         = "namespace"
	DeploymentFieldPaused            = "paused"
	DeploymentFieldStatus            = "status"
	DeploymentFieldStrategy          = "strategy"
	DeploymentFieldTemplate          = "template"
	DeploymentFieldUID               = "uid"
)

type Deployment struct {
	types.Resource
	APIVersion        string             `json:"apiVersion,omitempty"`
	Annotations       map[string]string  `json:"annotations,omitempty"`
	CreationTimestamp string             `json:"creationTimestamp,omitempty"`
	DeletionTimestamp string             `json:"deletionTimestamp,omitempty"`
	Deploy            *DeployParams      `json:"deploy,omitempty"`
	Kind              string             `json:"kind,omitempty"`
	Labels            map[string]string  `json:"labels,omitempty"`
	Name              string             `json:"name,omitempty"`
	Namespace         string             `json:"namespace,omitempty"`
	Paused            bool               `json:"paused,omitempty"`
	Status            DeploymentStatus   `json:"status,omitempty"`
	Strategy          DeploymentStrategy `json:"strategy,omitempty"`
	Template          PodTemplateSpec    `json:"template,omitempty"`
	UID               string             `json:"uid,omitempty"`
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
