package client

import (
	"github.com/rancher/norman/types"
)

const (
	ControllerRevisionType                 = "controllerRevision"
	ControllerRevisionFieldAnnotations     = "annotations"
	ControllerRevisionFieldCreated         = "created"
	ControllerRevisionFieldCreatorID       = "creatorId"
	ControllerRevisionFieldData            = "data"
	ControllerRevisionFieldLabels          = "labels"
	ControllerRevisionFieldName            = "name"
	ControllerRevisionFieldNamespaceId     = "namespaceId"
	ControllerRevisionFieldOwnerReferences = "ownerReferences"
	ControllerRevisionFieldProjectID       = "projectId"
	ControllerRevisionFieldPublicEndpoints = "publicEndpoints"
	ControllerRevisionFieldRemoved         = "removed"
	ControllerRevisionFieldRevision        = "revision"
	ControllerRevisionFieldUuid            = "uuid"
)

type ControllerRevision struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Data            *PodTemplateSpec  `json:"data,omitempty" yaml:"data,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProjectID       string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	PublicEndpoints []PublicEndpoint  `json:"publicEndpoints,omitempty" yaml:"publicEndpoints,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Revision        int64             `json:"revision,omitempty" yaml:"revision,omitempty"`
	Uuid            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type ControllerRevisionCollection struct {
	types.Collection
	Data   []ControllerRevision `json:"data,omitempty"`
	client *ControllerRevisionClient
}

type ControllerRevisionClient struct {
	apiClient *Client
}

type ControllerRevisionOperations interface {
	List(opts *types.ListOpts) (*ControllerRevisionCollection, error)
	Create(opts *ControllerRevision) (*ControllerRevision, error)
	Update(existing *ControllerRevision, updates interface{}) (*ControllerRevision, error)
	ByID(id string) (*ControllerRevision, error)
	Delete(container *ControllerRevision) error
}

func newControllerRevisionClient(apiClient *Client) *ControllerRevisionClient {
	return &ControllerRevisionClient{
		apiClient: apiClient,
	}
}

func (c *ControllerRevisionClient) Create(container *ControllerRevision) (*ControllerRevision, error) {
	resp := &ControllerRevision{}
	err := c.apiClient.Ops.DoCreate(ControllerRevisionType, container, resp)
	return resp, err
}

func (c *ControllerRevisionClient) Update(existing *ControllerRevision, updates interface{}) (*ControllerRevision, error) {
	resp := &ControllerRevision{}
	err := c.apiClient.Ops.DoUpdate(ControllerRevisionType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ControllerRevisionClient) List(opts *types.ListOpts) (*ControllerRevisionCollection, error) {
	resp := &ControllerRevisionCollection{}
	err := c.apiClient.Ops.DoList(ControllerRevisionType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ControllerRevisionCollection) Next() (*ControllerRevisionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ControllerRevisionCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ControllerRevisionClient) ByID(id string) (*ControllerRevision, error) {
	resp := &ControllerRevision{}
	err := c.apiClient.Ops.DoByID(ControllerRevisionType, id, resp)
	return resp, err
}

func (c *ControllerRevisionClient) Delete(container *ControllerRevision) error {
	return c.apiClient.Ops.DoResourceDelete(ControllerRevisionType, &container.Resource)
}
