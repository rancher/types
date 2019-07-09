package client

import (
	"github.com/rancher/norman/types"
)

const (
	ServiceMonitorType                   = "serviceMonitor"
	ServiceMonitorFieldAnnotations       = "annotations"
	ServiceMonitorFieldCreated           = "created"
	ServiceMonitorFieldCreatorID         = "creatorId"
	ServiceMonitorFieldEndpoints         = "endpoints"
	ServiceMonitorFieldJobLabel          = "jobLabel"
	ServiceMonitorFieldLabels            = "labels"
	ServiceMonitorFieldName              = "name"
	ServiceMonitorFieldNamespaceId       = "namespaceId"
	ServiceMonitorFieldNamespaceSelector = "namespaceSelector"
	ServiceMonitorFieldOwnerReferences   = "ownerReferences"
	ServiceMonitorFieldPodTargetLabels   = "podTargetLabels"
	ServiceMonitorFieldProjectID         = "projectId"
	ServiceMonitorFieldRemoved           = "removed"
	ServiceMonitorFieldSampleLimit       = "sampleLimit"
	ServiceMonitorFieldSelector          = "selector"
	ServiceMonitorFieldTargetLabels      = "targetLabels"
	ServiceMonitorFieldTargetService     = "targetService"
	ServiceMonitorFieldTargetWorkload    = "targetWorkload"
	ServiceMonitorFieldUUID              = "uuid"
)

type ServiceMonitor struct {
	types.Resource
	Annotations       map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created           string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID         string            `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Endpoints         []Endpoint        `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	JobLabel          string            `json:"jobLabel,omitempty" yaml:"job_label,omitempty"`
	Labels            map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name              string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId       string            `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	NamespaceSelector []string          `json:"namespaceSelector,omitempty" yaml:"namespace_selector,omitempty"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	PodTargetLabels   []string          `json:"podTargetLabels,omitempty" yaml:"pod_target_labels,omitempty"`
	ProjectID         string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Removed           string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	SampleLimit       int64             `json:"sampleLimit,omitempty" yaml:"sample_limit,omitempty"`
	Selector          *LabelSelector    `json:"selector,omitempty" yaml:"selector,omitempty"`
	TargetLabels      []string          `json:"targetLabels,omitempty" yaml:"target_labels,omitempty"`
	TargetService     string            `json:"targetService,omitempty" yaml:"target_service,omitempty"`
	TargetWorkload    string            `json:"targetWorkload,omitempty" yaml:"target_workload,omitempty"`
	UUID              string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ServiceMonitorCollection struct {
	types.Collection
	Data   []ServiceMonitor `json:"data,omitempty"`
	client *ServiceMonitorClient
}

type ServiceMonitorClient struct {
	apiClient *Client
}

type ServiceMonitorOperations interface {
	List(opts *types.ListOpts) (*ServiceMonitorCollection, error)
	Create(opts *ServiceMonitor) (*ServiceMonitor, error)
	Update(existing *ServiceMonitor, updates interface{}) (*ServiceMonitor, error)
	Replace(existing *ServiceMonitor) (*ServiceMonitor, error)
	ByID(id string) (*ServiceMonitor, error)
	Delete(container *ServiceMonitor) error
}

func newServiceMonitorClient(apiClient *Client) *ServiceMonitorClient {
	return &ServiceMonitorClient{
		apiClient: apiClient,
	}
}

func (c *ServiceMonitorClient) Create(container *ServiceMonitor) (*ServiceMonitor, error) {
	resp := &ServiceMonitor{}
	err := c.apiClient.Ops.DoCreate(ServiceMonitorType, container, resp)
	return resp, err
}

func (c *ServiceMonitorClient) Update(existing *ServiceMonitor, updates interface{}) (*ServiceMonitor, error) {
	resp := &ServiceMonitor{}
	err := c.apiClient.Ops.DoUpdate(ServiceMonitorType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceMonitorClient) Replace(obj *ServiceMonitor) (*ServiceMonitor, error) {
	resp := &ServiceMonitor{}
	err := c.apiClient.Ops.DoReplace(ServiceMonitorType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ServiceMonitorClient) List(opts *types.ListOpts) (*ServiceMonitorCollection, error) {
	resp := &ServiceMonitorCollection{}
	err := c.apiClient.Ops.DoList(ServiceMonitorType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceMonitorCollection) Next() (*ServiceMonitorCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceMonitorCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceMonitorClient) ByID(id string) (*ServiceMonitor, error) {
	resp := &ServiceMonitor{}
	err := c.apiClient.Ops.DoByID(ServiceMonitorType, id, resp)
	return resp, err
}

func (c *ServiceMonitorClient) Delete(container *ServiceMonitor) error {
	return c.apiClient.Ops.DoResourceDelete(ServiceMonitorType, &container.Resource)
}
