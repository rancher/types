package client

import (
	"github.com/rancher/norman/types"
)

const (
	NodeType                      = "node"
	NodeFieldAllocatable          = "allocatable"
	NodeFieldAnnotations          = "annotations"
	NodeFieldCapacity             = "capacity"
	NodeFieldCreated              = "created"
	NodeFieldCreatorID            = "creatorId"
	NodeFieldDescription          = "description"
	NodeFieldHostname             = "hostname"
	NodeFieldIPAddress            = "ipAddress"
	NodeFieldInfo                 = "info"
	NodeFieldLabels               = "labels"
	NodeFieldName                 = "name"
	NodeFieldNodeConditions       = "nodeConditions"
	NodeFieldOwnerReferences      = "ownerReferences"
	NodeFieldPodCidr              = "podCidr"
	NodeFieldProviderId           = "providerId"
	NodeFieldRemoved              = "removed"
	NodeFieldState                = "state"
	NodeFieldTaints               = "taints"
	NodeFieldTransitioning        = "transitioning"
	NodeFieldTransitioningMessage = "transitioningMessage"
	NodeFieldUnschedulable        = "unschedulable"
	NodeFieldUuid                 = "uuid"
	NodeFieldVolumesAttached      = "volumesAttached"
	NodeFieldVolumesInUse         = "volumesInUse"
)

type Node struct {
	types.Resource
	Allocatable          map[string]string         `json:"allocatable,omitempty"`
	Annotations          map[string]string         `json:"annotations,omitempty"`
	Capacity             map[string]string         `json:"capacity,omitempty"`
	Created              string                    `json:"created,omitempty"`
	CreatorID            string                    `json:"creatorId,omitempty"`
	Description          string                    `json:"description,omitempty"`
	Hostname             string                    `json:"hostname,omitempty"`
	IPAddress            string                    `json:"ipAddress,omitempty"`
	Info                 *NodeInfo                 `json:"info,omitempty"`
	Labels               map[string]string         `json:"labels,omitempty"`
	Name                 string                    `json:"name,omitempty"`
	NodeConditions       []NodeCondition           `json:"nodeConditions,omitempty"`
	OwnerReferences      []OwnerReference          `json:"ownerReferences,omitempty"`
	PodCidr              string                    `json:"podCidr,omitempty"`
	ProviderId           string                    `json:"providerId,omitempty"`
	Removed              string                    `json:"removed,omitempty"`
	State                string                    `json:"state,omitempty"`
	Taints               []Taint                   `json:"taints,omitempty"`
	Transitioning        string                    `json:"transitioning,omitempty"`
	TransitioningMessage string                    `json:"transitioningMessage,omitempty"`
	Unschedulable        *bool                     `json:"unschedulable,omitempty"`
	Uuid                 string                    `json:"uuid,omitempty"`
	VolumesAttached      map[string]AttachedVolume `json:"volumesAttached,omitempty"`
	VolumesInUse         []string                  `json:"volumesInUse,omitempty"`
}
type NodeCollection struct {
	types.Collection
	Data   []Node `json:"data,omitempty"`
	client *NodeClient
}

type NodeClient struct {
	apiClient *Client
}

type NodeOperations interface {
	List(opts *types.ListOpts) (*NodeCollection, error)
	Create(opts *Node) (*Node, error)
	Update(existing *Node, updates interface{}) (*Node, error)
	ByID(id string) (*Node, error)
	Delete(container *Node) error
}

func newNodeClient(apiClient *Client) *NodeClient {
	return &NodeClient{
		apiClient: apiClient,
	}
}

func (c *NodeClient) Create(container *Node) (*Node, error) {
	resp := &Node{}
	err := c.apiClient.Ops.DoCreate(NodeType, container, resp)
	return resp, err
}

func (c *NodeClient) Update(existing *Node, updates interface{}) (*Node, error) {
	resp := &Node{}
	err := c.apiClient.Ops.DoUpdate(NodeType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NodeClient) List(opts *types.ListOpts) (*NodeCollection, error) {
	resp := &NodeCollection{}
	err := c.apiClient.Ops.DoList(NodeType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NodeCollection) Next() (*NodeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NodeCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NodeClient) ByID(id string) (*Node, error) {
	resp := &Node{}
	err := c.apiClient.Ops.DoByID(NodeType, id, resp)
	return resp, err
}

func (c *NodeClient) Delete(container *Node) error {
	return c.apiClient.Ops.DoResourceDelete(NodeType, &container.Resource)
}
