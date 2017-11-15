package client

import (
	"github.com/rancher/norman/types"
)

const (
	NodeType                   = "node"
	NodeFieldAPIVersion        = "apiVersion"
	NodeFieldAnnotations       = "annotations"
	NodeFieldConfigSource      = "configSource"
	NodeFieldCreationTimestamp = "creationTimestamp"
	NodeFieldDeletionTimestamp = "deletionTimestamp"
	NodeFieldExternalID        = "externalID"
	NodeFieldKind              = "kind"
	NodeFieldLabels            = "labels"
	NodeFieldName              = "name"
	NodeFieldNamespace         = "namespace"
	NodeFieldPodCIDR           = "podCIDR"
	NodeFieldProviderID        = "providerID"
	NodeFieldStatus            = "status"
	NodeFieldTaints            = "taints"
	NodeFieldUID               = "uid"
	NodeFieldUnschedulable     = "unschedulable"
)

type Node struct {
	types.Resource
	APIVersion        string            `json:"apiVersion,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	ConfigSource      *NodeConfigSource `json:"configSource,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp string            `json:"deletionTimestamp,omitempty"`
	ExternalID        string            `json:"externalID,omitempty"`
	Kind              string            `json:"kind,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	PodCIDR           string            `json:"podCIDR,omitempty"`
	ProviderID        string            `json:"providerID,omitempty"`
	Status            NodeStatus        `json:"status,omitempty"`
	Taints            []Taint           `json:"taints,omitempty"`
	UID               string            `json:"uid,omitempty"`
	Unschedulable     bool              `json:"unschedulable,omitempty"`
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
