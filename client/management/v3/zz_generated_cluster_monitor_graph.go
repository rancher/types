package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterMonitorGraphType                        = "clusterMonitorGraph"
	ClusterMonitorGraphFieldAnnotations            = "annotations"
	ClusterMonitorGraphFieldClusterID              = "clusterId"
	ClusterMonitorGraphFieldCreated                = "created"
	ClusterMonitorGraphFieldCreatorID              = "creatorId"
	ClusterMonitorGraphFieldDescription            = "description"
	ClusterMonitorGraphFieldDetailsMetricsSelector = "detailsMetricsSelector"
	ClusterMonitorGraphFieldDisplayResourceType    = "displayResourceType"
	ClusterMonitorGraphFieldGraphType              = "graphType"
	ClusterMonitorGraphFieldLabels                 = "labels"
	ClusterMonitorGraphFieldMetricsSelector        = "metricsSelector"
	ClusterMonitorGraphFieldName                   = "name"
	ClusterMonitorGraphFieldNamespaceId            = "namespaceId"
	ClusterMonitorGraphFieldOwnerReferences        = "ownerReferences"
	ClusterMonitorGraphFieldPriority               = "priority"
	ClusterMonitorGraphFieldRemoved                = "removed"
	ClusterMonitorGraphFieldResourceType           = "resourceType"
	ClusterMonitorGraphFieldUUID                   = "uuid"
	ClusterMonitorGraphFieldYAxis                  = "yAxis"
)

type ClusterMonitorGraph struct {
	types.Resource
	Annotations            map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterID              string            `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Created                string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID              string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Description            string            `json:"description,omitempty" yaml:"description,omitempty"`
	DetailsMetricsSelector map[string]string `json:"detailsMetricsSelector,omitempty" yaml:"details_metrics_selector,omitempty"`
	DisplayResourceType    string            `json:"displayResourceType,omitempty" yaml:"display_resource_type,omitempty"`
	GraphType              string            `json:"graphType,omitempty" yaml:"graph_type,omitempty"`
	Labels                 map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	MetricsSelector        map[string]string `json:"metricsSelector,omitempty" yaml:"metrics_selector,omitempty"`
	Name                   string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId            string            `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	OwnerReferences        []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Priority               int64             `json:"priority,omitempty" yaml:"priority,omitempty"`
	Removed                string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	ResourceType           string            `json:"resourceType,omitempty" yaml:"resource_type,omitempty"`
	UUID                   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	YAxis                  *YAxis            `json:"yAxis,omitempty" yaml:"y_axis,omitempty"`
}

type ClusterMonitorGraphCollection struct {
	types.Collection
	Data   []ClusterMonitorGraph `json:"data,omitempty"`
	client *ClusterMonitorGraphClient
}

type ClusterMonitorGraphClient struct {
	apiClient *Client
}

type ClusterMonitorGraphOperations interface {
	List(opts *types.ListOpts) (*ClusterMonitorGraphCollection, error)
	Create(opts *ClusterMonitorGraph) (*ClusterMonitorGraph, error)
	Update(existing *ClusterMonitorGraph, updates interface{}) (*ClusterMonitorGraph, error)
	Replace(existing *ClusterMonitorGraph) (*ClusterMonitorGraph, error)
	ByID(id string) (*ClusterMonitorGraph, error)
	Delete(container *ClusterMonitorGraph) error

	CollectionActionQuery(resource *ClusterMonitorGraphCollection, input *QueryGraphInput) (*QueryClusterGraphOutput, error)
}

func newClusterMonitorGraphClient(apiClient *Client) *ClusterMonitorGraphClient {
	return &ClusterMonitorGraphClient{
		apiClient: apiClient,
	}
}

func (c *ClusterMonitorGraphClient) Create(container *ClusterMonitorGraph) (*ClusterMonitorGraph, error) {
	resp := &ClusterMonitorGraph{}
	err := c.apiClient.Ops.DoCreate(ClusterMonitorGraphType, container, resp)
	return resp, err
}

func (c *ClusterMonitorGraphClient) Update(existing *ClusterMonitorGraph, updates interface{}) (*ClusterMonitorGraph, error) {
	resp := &ClusterMonitorGraph{}
	err := c.apiClient.Ops.DoUpdate(ClusterMonitorGraphType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterMonitorGraphClient) Replace(obj *ClusterMonitorGraph) (*ClusterMonitorGraph, error) {
	resp := &ClusterMonitorGraph{}
	err := c.apiClient.Ops.DoReplace(ClusterMonitorGraphType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterMonitorGraphClient) List(opts *types.ListOpts) (*ClusterMonitorGraphCollection, error) {
	resp := &ClusterMonitorGraphCollection{}
	err := c.apiClient.Ops.DoList(ClusterMonitorGraphType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterMonitorGraphCollection) Next() (*ClusterMonitorGraphCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterMonitorGraphCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterMonitorGraphClient) ByID(id string) (*ClusterMonitorGraph, error) {
	resp := &ClusterMonitorGraph{}
	err := c.apiClient.Ops.DoByID(ClusterMonitorGraphType, id, resp)
	return resp, err
}

func (c *ClusterMonitorGraphClient) Delete(container *ClusterMonitorGraph) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterMonitorGraphType, &container.Resource)
}

func (c *ClusterMonitorGraphClient) CollectionActionQuery(resource *ClusterMonitorGraphCollection, input *QueryGraphInput) (*QueryClusterGraphOutput, error) {
	resp := &QueryClusterGraphOutput{}
	err := c.apiClient.Ops.DoCollectionAction(ClusterMonitorGraphType, "query", &resource.Collection, input, resp)
	return resp, err
}
