package client

import (
	"github.com/rancher/norman/types"
)

const (
	PersistentVolumeClaimType                      = "persistentVolumeClaim"
	PersistentVolumeClaimFieldAccessModes          = "accessModes"
	PersistentVolumeClaimFieldAnnotations          = "annotations"
	PersistentVolumeClaimFieldCreated              = "created"
	PersistentVolumeClaimFieldCreatorID            = "creatorId"
	PersistentVolumeClaimFieldDataSource           = "dataSource"
	PersistentVolumeClaimFieldLabels               = "labels"
	PersistentVolumeClaimFieldName                 = "name"
	PersistentVolumeClaimFieldNamespaceId          = "namespaceId"
	PersistentVolumeClaimFieldOwnerReferences      = "ownerReferences"
	PersistentVolumeClaimFieldProjectID            = "projectId"
	PersistentVolumeClaimFieldRemoved              = "removed"
	PersistentVolumeClaimFieldResources            = "resources"
	PersistentVolumeClaimFieldSelector             = "selector"
	PersistentVolumeClaimFieldState                = "state"
	PersistentVolumeClaimFieldStatus               = "status"
	PersistentVolumeClaimFieldStorageClassID       = "storageClassId"
	PersistentVolumeClaimFieldTransitioning        = "transitioning"
	PersistentVolumeClaimFieldTransitioningMessage = "transitioningMessage"
	PersistentVolumeClaimFieldUUID                 = "uuid"
	PersistentVolumeClaimFieldVolumeID             = "volumeId"
	PersistentVolumeClaimFieldVolumeMode           = "volumeMode"
)

type PersistentVolumeClaim struct {
	types.Resource
	AccessModes          []string                     `json:"accessModes,omitempty" yaml:"access_modes,omitempty"`
	Annotations          map[string]string            `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created              string                       `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string                       `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DataSource           *TypedLocalObjectReference   `json:"dataSource,omitempty" yaml:"data_source,omitempty"`
	Labels               map[string]string            `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string                       `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string                       `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	OwnerReferences      []OwnerReference             `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	ProjectID            string                       `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Removed              string                       `json:"removed,omitempty" yaml:"removed,omitempty"`
	Resources            *ResourceRequirements        `json:"resources,omitempty" yaml:"resources,omitempty"`
	Selector             *LabelSelector               `json:"selector,omitempty" yaml:"selector,omitempty"`
	State                string                       `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *PersistentVolumeClaimStatus `json:"status,omitempty" yaml:"status,omitempty"`
	StorageClassID       string                       `json:"storageClassId,omitempty" yaml:"storage_class_id,omitempty"`
	Transitioning        string                       `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string                       `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                 string                       `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	VolumeID             string                       `json:"volumeId,omitempty" yaml:"volume_id,omitempty"`
	VolumeMode           string                       `json:"volumeMode,omitempty" yaml:"volume_mode,omitempty"`
}

type PersistentVolumeClaimCollection struct {
	types.Collection
	Data   []PersistentVolumeClaim `json:"data,omitempty"`
	client *PersistentVolumeClaimClient
}

type PersistentVolumeClaimClient struct {
	apiClient *Client
}

type PersistentVolumeClaimOperations interface {
	List(opts *types.ListOpts) (*PersistentVolumeClaimCollection, error)
	Create(opts *PersistentVolumeClaim) (*PersistentVolumeClaim, error)
	Update(existing *PersistentVolumeClaim, updates interface{}) (*PersistentVolumeClaim, error)
	Replace(existing *PersistentVolumeClaim) (*PersistentVolumeClaim, error)
	ByID(id string) (*PersistentVolumeClaim, error)
	Delete(container *PersistentVolumeClaim) error
}

func newPersistentVolumeClaimClient(apiClient *Client) *PersistentVolumeClaimClient {
	return &PersistentVolumeClaimClient{
		apiClient: apiClient,
	}
}

func (c *PersistentVolumeClaimClient) Create(container *PersistentVolumeClaim) (*PersistentVolumeClaim, error) {
	resp := &PersistentVolumeClaim{}
	err := c.apiClient.Ops.DoCreate(PersistentVolumeClaimType, container, resp)
	return resp, err
}

func (c *PersistentVolumeClaimClient) Update(existing *PersistentVolumeClaim, updates interface{}) (*PersistentVolumeClaim, error) {
	resp := &PersistentVolumeClaim{}
	err := c.apiClient.Ops.DoUpdate(PersistentVolumeClaimType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PersistentVolumeClaimClient) Replace(obj *PersistentVolumeClaim) (*PersistentVolumeClaim, error) {
	resp := &PersistentVolumeClaim{}
	err := c.apiClient.Ops.DoReplace(PersistentVolumeClaimType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *PersistentVolumeClaimClient) List(opts *types.ListOpts) (*PersistentVolumeClaimCollection, error) {
	resp := &PersistentVolumeClaimCollection{}
	err := c.apiClient.Ops.DoList(PersistentVolumeClaimType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PersistentVolumeClaimCollection) Next() (*PersistentVolumeClaimCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PersistentVolumeClaimCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PersistentVolumeClaimClient) ByID(id string) (*PersistentVolumeClaim, error) {
	resp := &PersistentVolumeClaim{}
	err := c.apiClient.Ops.DoByID(PersistentVolumeClaimType, id, resp)
	return resp, err
}

func (c *PersistentVolumeClaimClient) Delete(container *PersistentVolumeClaim) error {
	return c.apiClient.Ops.DoResourceDelete(PersistentVolumeClaimType, &container.Resource)
}
