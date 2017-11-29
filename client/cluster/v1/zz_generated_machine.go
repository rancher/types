package client

import (
	"github.com/rancher/norman/types"
)

const (
	MachineType                          = "machine"
	MachineFieldAmazonEC2Config          = "amazonEc2Config"
	MachineFieldAnnotations              = "annotations"
	MachineFieldAuthCertificateAuthority = "authCertificateAuthority"
	MachineFieldAuthKey                  = "authKey"
	MachineFieldAzureConfig              = "azureConfig"
	MachineFieldClusterId                = "clusterId"
	MachineFieldCreated                  = "created"
	MachineFieldDescription              = "description"
	MachineFieldDigitalOceanConfig       = "digitalOceanConfig"
	MachineFieldDockerVersion            = "dockerVersion"
	MachineFieldDriver                   = "driver"
	MachineFieldEngineEnv                = "engineEnv"
	MachineFieldEngineInsecureRegistry   = "engineInsecureRegistry"
	MachineFieldEngineInstallURL         = "engineInstallURL"
	MachineFieldEngineLabel              = "engineLabel"
	MachineFieldEngineOpt                = "engineOpt"
	MachineFieldEngineRegistryMirror     = "engineRegistryMirror"
	MachineFieldEngineStorageDriver      = "engineStorageDriver"
	MachineFieldExternalID               = "externalId"
	MachineFieldHostname                 = "hostname"
	MachineFieldId                       = "id"
	MachineFieldLabels                   = "labels"
	MachineFieldMachineTemplateId        = "machineTemplateId"
	MachineFieldName                     = "name"
	MachineFieldOwnerReferences          = "ownerReferences"
	MachineFieldRemoved                  = "removed"
	MachineFieldResourcePath             = "resourcePath"
	MachineFieldState                    = "state"
	MachineFieldStatus                   = "status"
	MachineFieldTransitioning            = "transitioning"
	MachineFieldTransitioningMessage     = "transitioningMessage"
	MachineFieldUuid                     = "uuid"
)

type Machine struct {
	types.Resource
	AmazonEC2Config          *AmazonEC2Config    `json:"amazonEc2Config,omitempty"`
	Annotations              map[string]string   `json:"annotations,omitempty"`
	AuthCertificateAuthority string              `json:"authCertificateAuthority,omitempty"`
	AuthKey                  string              `json:"authKey,omitempty"`
	AzureConfig              *AzureConfig        `json:"azureConfig,omitempty"`
	ClusterId                string              `json:"clusterId,omitempty"`
	Created                  string              `json:"created,omitempty"`
	Description              string              `json:"description,omitempty"`
	DigitalOceanConfig       *DigitalOceanConfig `json:"digitalOceanConfig,omitempty"`
	DockerVersion            string              `json:"dockerVersion,omitempty"`
	Driver                   string              `json:"driver,omitempty"`
	EngineEnv                map[string]string   `json:"engineEnv,omitempty"`
	EngineInsecureRegistry   []string            `json:"engineInsecureRegistry,omitempty"`
	EngineInstallURL         string              `json:"engineInstallURL,omitempty"`
	EngineLabel              map[string]string   `json:"engineLabel,omitempty"`
	EngineOpt                map[string]string   `json:"engineOpt,omitempty"`
	EngineRegistryMirror     []string            `json:"engineRegistryMirror,omitempty"`
	EngineStorageDriver      string              `json:"engineStorageDriver,omitempty"`
	ExternalID               string              `json:"externalId,omitempty"`
	Hostname                 string              `json:"hostname,omitempty"`
	Id                       string              `json:"id,omitempty"`
	Labels                   map[string]string   `json:"labels,omitempty"`
	MachineTemplateId        string              `json:"machineTemplateId,omitempty"`
	Name                     string              `json:"name,omitempty"`
	OwnerReferences          []OwnerReference    `json:"ownerReferences,omitempty"`
	Removed                  string              `json:"removed,omitempty"`
	ResourcePath             string              `json:"resourcePath,omitempty"`
	State                    string              `json:"state,omitempty"`
	Status                   *MachineStatus      `json:"status,omitempty"`
	Transitioning            string              `json:"transitioning,omitempty"`
	TransitioningMessage     string              `json:"transitioningMessage,omitempty"`
	Uuid                     string              `json:"uuid,omitempty"`
}
type MachineCollection struct {
	types.Collection
	Data   []Machine `json:"data,omitempty"`
	client *MachineClient
}

type MachineClient struct {
	apiClient *Client
}

type MachineOperations interface {
	List(opts *types.ListOpts) (*MachineCollection, error)
	Create(opts *Machine) (*Machine, error)
	Update(existing *Machine, updates interface{}) (*Machine, error)
	ByID(id string) (*Machine, error)
	Delete(container *Machine) error
}

func newMachineClient(apiClient *Client) *MachineClient {
	return &MachineClient{
		apiClient: apiClient,
	}
}

func (c *MachineClient) Create(container *Machine) (*Machine, error) {
	resp := &Machine{}
	err := c.apiClient.Ops.DoCreate(MachineType, container, resp)
	return resp, err
}

func (c *MachineClient) Update(existing *Machine, updates interface{}) (*Machine, error) {
	resp := &Machine{}
	err := c.apiClient.Ops.DoUpdate(MachineType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MachineClient) List(opts *types.ListOpts) (*MachineCollection, error) {
	resp := &MachineCollection{}
	err := c.apiClient.Ops.DoList(MachineType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *MachineCollection) Next() (*MachineCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &MachineCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *MachineClient) ByID(id string) (*Machine, error) {
	resp := &Machine{}
	err := c.apiClient.Ops.DoByID(MachineType, id, resp)
	return resp, err
}

func (c *MachineClient) Delete(container *Machine) error {
	return c.apiClient.Ops.DoResourceDelete(MachineType, &container.Resource)
}
