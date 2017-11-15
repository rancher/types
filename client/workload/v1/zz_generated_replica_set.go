package client

import (
	"github.com/rancher/norman/types"
)

const (
	ReplicaSetType                   = "replicaSet"
	ReplicaSetFieldAPIVersion        = "apiVersion"
	ReplicaSetFieldAnnotations       = "annotations"
	ReplicaSetFieldCreationTimestamp = "creationTimestamp"
	ReplicaSetFieldDeletionTimestamp = "deletionTimestamp"
	ReplicaSetFieldDeploy            = "deploy"
	ReplicaSetFieldKind              = "kind"
	ReplicaSetFieldLabels            = "labels"
	ReplicaSetFieldName              = "name"
	ReplicaSetFieldNamespace         = "namespace"
	ReplicaSetFieldStatus            = "status"
	ReplicaSetFieldTemplate          = "template"
	ReplicaSetFieldUID               = "uid"
)

type ReplicaSet struct {
	types.Resource
	APIVersion        string            `json:"apiVersion,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp string            `json:"deletionTimestamp,omitempty"`
	Deploy            *DeployParams     `json:"deploy,omitempty"`
	Kind              string            `json:"kind,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	Status            ReplicaSetStatus  `json:"status,omitempty"`
	Template          PodTemplateSpec   `json:"template,omitempty"`
	UID               string            `json:"uid,omitempty"`
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
