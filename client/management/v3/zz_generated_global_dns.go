package client

import (
	"github.com/rancher/norman/types"
)

const (
	GlobalDNSType                      = "globalDNS"
	GlobalDNSFieldAnnotations          = "annotations"
	GlobalDNSFieldClusterIDs           = "clusterIds"
	GlobalDNSFieldCreated              = "created"
	GlobalDNSFieldCreatorID            = "creatorId"
	GlobalDNSFieldLabels               = "labels"
	GlobalDNSFieldName                 = "name"
	GlobalDNSFieldOwnerReferences      = "ownerReferences"
	GlobalDNSFieldProviderName         = "providerName"
	GlobalDNSFieldRemoved              = "removed"
	GlobalDNSFieldRootDomain           = "rootDomain"
	GlobalDNSFieldState                = "state"
	GlobalDNSFieldStatus               = "status"
	GlobalDNSFieldTransitioning        = "transitioning"
	GlobalDNSFieldTransitioningMessage = "transitioningMessage"
	GlobalDNSFieldUUID                 = "uuid"
)

type GlobalDNS struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterIDs           []string          `json:"clusterIds,omitempty" yaml:"clusterIds,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProviderName         string            `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	RootDomain           string            `json:"rootDomain,omitempty" yaml:"rootDomain,omitempty"`
	State                string            `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *GlobalDNSStatus  `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type GlobalDNSCollection struct {
	types.Collection
	Data   []GlobalDNS `json:"data,omitempty"`
	client *GlobalDNSClient
}

type GlobalDNSClient struct {
	apiClient *Client
}

type GlobalDNSOperations interface {
	List(opts *types.ListOpts) (*GlobalDNSCollection, error)
	Create(opts *GlobalDNS) (*GlobalDNS, error)
	Update(existing *GlobalDNS, updates interface{}) (*GlobalDNS, error)
	Replace(existing *GlobalDNS) (*GlobalDNS, error)
	ByID(id string) (*GlobalDNS, error)
	Delete(container *GlobalDNS) error
}

func newGlobalDNSClient(apiClient *Client) *GlobalDNSClient {
	return &GlobalDNSClient{
		apiClient: apiClient,
	}
}

func (c *GlobalDNSClient) Create(container *GlobalDNS) (*GlobalDNS, error) {
	resp := &GlobalDNS{}
	err := c.apiClient.Ops.DoCreate(GlobalDNSType, container, resp)
	return resp, err
}

func (c *GlobalDNSClient) Update(existing *GlobalDNS, updates interface{}) (*GlobalDNS, error) {
	resp := &GlobalDNS{}
	err := c.apiClient.Ops.DoUpdate(GlobalDNSType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GlobalDNSClient) Replace(obj *GlobalDNS) (*GlobalDNS, error) {
	resp := &GlobalDNS{}
	err := c.apiClient.Ops.DoReplace(GlobalDNSType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *GlobalDNSClient) List(opts *types.ListOpts) (*GlobalDNSCollection, error) {
	resp := &GlobalDNSCollection{}
	err := c.apiClient.Ops.DoList(GlobalDNSType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *GlobalDNSCollection) Next() (*GlobalDNSCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &GlobalDNSCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *GlobalDNSClient) ByID(id string) (*GlobalDNS, error) {
	resp := &GlobalDNS{}
	err := c.apiClient.Ops.DoByID(GlobalDNSType, id, resp)
	return resp, err
}

func (c *GlobalDNSClient) Delete(container *GlobalDNS) error {
	return c.apiClient.Ops.DoResourceDelete(GlobalDNSType, &container.Resource)
}
