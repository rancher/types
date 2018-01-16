package client

import (
	"github.com/rancher/norman/types"
)

const (
	MachineDriverType                      = "machineDriver"
	MachineDriverFieldActive               = "active"
	MachineDriverFieldAnnotations          = "annotations"
	MachineDriverFieldBuiltin              = "builtin"
	MachineDriverFieldChecksum             = "checksum"
	MachineDriverFieldCreated              = "created"
	MachineDriverFieldCreatorID            = "creatorId"
	MachineDriverFieldDescription          = "description"
	MachineDriverFieldExternalID           = "externalId"
	MachineDriverFieldLabels               = "labels"
	MachineDriverFieldName                 = "name"
	MachineDriverFieldOwnerReferences      = "ownerReferences"
	MachineDriverFieldRemoved              = "removed"
	MachineDriverFieldState                = "state"
	MachineDriverFieldStatus               = "status"
	MachineDriverFieldTransitioning        = "transitioning"
	MachineDriverFieldTransitioningMessage = "transitioningMessage"
	MachineDriverFieldUIURL                = "uiUrl"
	MachineDriverFieldURL                  = "url"
	MachineDriverFieldUuid                 = "uuid"
)

type MachineDriver struct {
	types.Resource
	Active               *bool                `json:"active,omitempty"`
	Annotations          map[string]string    `json:"annotations,omitempty"`
	Builtin              *bool                `json:"builtin,omitempty"`
	Checksum             string               `json:"checksum,omitempty"`
	Created              string               `json:"created,omitempty"`
	CreatorID            string               `json:"creatorId,omitempty"`
	Description          string               `json:"description,omitempty"`
	ExternalID           string               `json:"externalId,omitempty"`
	Labels               map[string]string    `json:"labels,omitempty"`
	Name                 string               `json:"name,omitempty"`
	OwnerReferences      []OwnerReference     `json:"ownerReferences,omitempty"`
	Removed              string               `json:"removed,omitempty"`
	State                string               `json:"state,omitempty"`
	Status               *MachineDriverStatus `json:"status,omitempty"`
	Transitioning        string               `json:"transitioning,omitempty"`
	TransitioningMessage string               `json:"transitioningMessage,omitempty"`
	UIURL                string               `json:"uiUrl,omitempty"`
	URL                  string               `json:"url,omitempty"`
	Uuid                 string               `json:"uuid,omitempty"`
}
type MachineDriverCollection struct {
	types.Collection
	Data   []MachineDriver `json:"data,omitempty"`
	client *MachineDriverClient
}

type MachineDriverClient struct {
	apiClient *Client
}

type MachineDriverOperations interface {
	List(opts *types.ListOpts) (*MachineDriverCollection, error)
	Create(opts *MachineDriver) (*MachineDriver, error)
	Update(existing *MachineDriver, updates interface{}) (*MachineDriver, error)
	ByID(id string) (*MachineDriver, error)
	Delete(container *MachineDriver) error

	ActionActivate(*MachineDriver) (*MachineDriver, error)

	ActionDeactivate(*MachineDriver) (*MachineDriver, error)
}

func newMachineDriverClient(apiClient *Client) *MachineDriverClient {
	return &MachineDriverClient{
		apiClient: apiClient,
	}
}

func (c *MachineDriverClient) Create(container *MachineDriver) (*MachineDriver, error) {
	resp := &MachineDriver{}
	err := c.apiClient.Ops.DoCreate(MachineDriverType, container, resp)
	return resp, err
}

func (c *MachineDriverClient) Update(existing *MachineDriver, updates interface{}) (*MachineDriver, error) {
	resp := &MachineDriver{}
	err := c.apiClient.Ops.DoUpdate(MachineDriverType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MachineDriverClient) List(opts *types.ListOpts) (*MachineDriverCollection, error) {
	resp := &MachineDriverCollection{}
	err := c.apiClient.Ops.DoList(MachineDriverType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *MachineDriverCollection) Next() (*MachineDriverCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &MachineDriverCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *MachineDriverClient) ByID(id string) (*MachineDriver, error) {
	resp := &MachineDriver{}
	err := c.apiClient.Ops.DoByID(MachineDriverType, id, resp)
	return resp, err
}

func (c *MachineDriverClient) Delete(container *MachineDriver) error {
	return c.apiClient.Ops.DoResourceDelete(MachineDriverType, &container.Resource)
}

func (c *MachineDriverClient) ActionActivate(resource *MachineDriver) (*MachineDriver, error) {

	resp := &MachineDriver{}

	err := c.apiClient.Ops.DoAction(MachineDriverType, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *MachineDriverClient) ActionDeactivate(resource *MachineDriver) (*MachineDriver, error) {

	resp := &MachineDriver{}

	err := c.apiClient.Ops.DoAction(MachineDriverType, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}
