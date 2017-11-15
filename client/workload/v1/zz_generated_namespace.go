package client

import (
	"github.com/rancher/norman/types"
)

const (
	NamespaceType                   = "namespace"
	NamespaceFieldAPIVersion        = "apiVersion"
	NamespaceFieldAnnotations       = "annotations"
	NamespaceFieldCreationTimestamp = "creationTimestamp"
	NamespaceFieldDeletionTimestamp = "deletionTimestamp"
	NamespaceFieldKind              = "kind"
	NamespaceFieldLabels            = "labels"
	NamespaceFieldName              = "name"
	NamespaceFieldNamespace         = "namespace"
	NamespaceFieldStatus            = "status"
	NamespaceFieldUID               = "uid"
)

type Namespace struct {
	types.Resource
	APIVersion        string            `json:"apiVersion,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp string            `json:"deletionTimestamp,omitempty"`
	Kind              string            `json:"kind,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	Status            NamespaceStatus   `json:"status,omitempty"`
	UID               string            `json:"uid,omitempty"`
}
type NamespaceCollection struct {
	types.Collection
	Data   []Namespace `json:"data,omitempty"`
	client *NamespaceClient
}

type NamespaceClient struct {
	apiClient *Client
}

type NamespaceOperations interface {
	List(opts *types.ListOpts) (*NamespaceCollection, error)
	Create(opts *Namespace) (*Namespace, error)
	Update(existing *Namespace, updates interface{}) (*Namespace, error)
	ByID(id string) (*Namespace, error)
	Delete(container *Namespace) error
}

func newNamespaceClient(apiClient *Client) *NamespaceClient {
	return &NamespaceClient{
		apiClient: apiClient,
	}
}

func (c *NamespaceClient) Create(container *Namespace) (*Namespace, error) {
	resp := &Namespace{}
	err := c.apiClient.Ops.DoCreate(NamespaceType, container, resp)
	return resp, err
}

func (c *NamespaceClient) Update(existing *Namespace, updates interface{}) (*Namespace, error) {
	resp := &Namespace{}
	err := c.apiClient.Ops.DoUpdate(NamespaceType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NamespaceClient) List(opts *types.ListOpts) (*NamespaceCollection, error) {
	resp := &NamespaceCollection{}
	err := c.apiClient.Ops.DoList(NamespaceType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NamespaceCollection) Next() (*NamespaceCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NamespaceCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NamespaceClient) ByID(id string) (*Namespace, error) {
	resp := &Namespace{}
	err := c.apiClient.Ops.DoByID(NamespaceType, id, resp)
	return resp, err
}

func (c *NamespaceClient) Delete(container *Namespace) error {
	return c.apiClient.Ops.DoResourceDelete(NamespaceType, &container.Resource)
}
