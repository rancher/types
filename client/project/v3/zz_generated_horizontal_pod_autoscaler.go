package client

import (
	"github.com/rancher/norman/types"
)

const (
	HorizontalPodAutoscalerType                      = "horizontalPodAutoscaler"
	HorizontalPodAutoscalerFieldAnnotations          = "annotations"
	HorizontalPodAutoscalerFieldConditions           = "conditions"
	HorizontalPodAutoscalerFieldCreated              = "created"
	HorizontalPodAutoscalerFieldCreatorID            = "creatorId"
	HorizontalPodAutoscalerFieldCurrentReplicas      = "currentReplicas"
	HorizontalPodAutoscalerFieldDescription          = "description"
	HorizontalPodAutoscalerFieldDesiredReplicas      = "desiredReplicas"
	HorizontalPodAutoscalerFieldLabels               = "labels"
	HorizontalPodAutoscalerFieldLastScaleTime        = "lastScaleTime"
	HorizontalPodAutoscalerFieldMaxReplicas          = "maxReplicas"
	HorizontalPodAutoscalerFieldMetrics              = "metrics"
	HorizontalPodAutoscalerFieldMinReplicas          = "minReplicas"
	HorizontalPodAutoscalerFieldName                 = "name"
	HorizontalPodAutoscalerFieldNamespaceId          = "namespaceId"
	HorizontalPodAutoscalerFieldObservedGeneration   = "observedGeneration"
	HorizontalPodAutoscalerFieldOwnerReferences      = "ownerReferences"
	HorizontalPodAutoscalerFieldProjectID            = "projectId"
	HorizontalPodAutoscalerFieldRemoved              = "removed"
	HorizontalPodAutoscalerFieldState                = "state"
	HorizontalPodAutoscalerFieldTransitioning        = "transitioning"
	HorizontalPodAutoscalerFieldTransitioningMessage = "transitioningMessage"
	HorizontalPodAutoscalerFieldUUID                 = "uuid"
	HorizontalPodAutoscalerFieldWorkloadId           = "workloadId"
)

type HorizontalPodAutoscaler struct {
	types.Resource
	Annotations          map[string]string                  `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Conditions           []HorizontalPodAutoscalerCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created              string                             `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string                             `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	CurrentReplicas      int64                              `json:"currentReplicas,omitempty" yaml:"current_replicas,omitempty"`
	Description          string                             `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredReplicas      int64                              `json:"desiredReplicas,omitempty" yaml:"desired_replicas,omitempty"`
	Labels               map[string]string                  `json:"labels,omitempty" yaml:"labels,omitempty"`
	LastScaleTime        string                             `json:"lastScaleTime,omitempty" yaml:"last_scale_time,omitempty"`
	MaxReplicas          int64                              `json:"maxReplicas,omitempty" yaml:"max_replicas,omitempty"`
	Metrics              []Metric                           `json:"metrics,omitempty" yaml:"metrics,omitempty"`
	MinReplicas          *int64                             `json:"minReplicas,omitempty" yaml:"min_replicas,omitempty"`
	Name                 string                             `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string                             `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	ObservedGeneration   *int64                             `json:"observedGeneration,omitempty" yaml:"observed_generation,omitempty"`
	OwnerReferences      []OwnerReference                   `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	ProjectID            string                             `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Removed              string                             `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string                             `json:"state,omitempty" yaml:"state,omitempty"`
	Transitioning        string                             `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string                             `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                 string                             `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	WorkloadId           string                             `json:"workloadId,omitempty" yaml:"workload_id,omitempty"`
}

type HorizontalPodAutoscalerCollection struct {
	types.Collection
	Data   []HorizontalPodAutoscaler `json:"data,omitempty"`
	client *HorizontalPodAutoscalerClient
}

type HorizontalPodAutoscalerClient struct {
	apiClient *Client
}

type HorizontalPodAutoscalerOperations interface {
	List(opts *types.ListOpts) (*HorizontalPodAutoscalerCollection, error)
	Create(opts *HorizontalPodAutoscaler) (*HorizontalPodAutoscaler, error)
	Update(existing *HorizontalPodAutoscaler, updates interface{}) (*HorizontalPodAutoscaler, error)
	Replace(existing *HorizontalPodAutoscaler) (*HorizontalPodAutoscaler, error)
	ByID(id string) (*HorizontalPodAutoscaler, error)
	Delete(container *HorizontalPodAutoscaler) error
}

func newHorizontalPodAutoscalerClient(apiClient *Client) *HorizontalPodAutoscalerClient {
	return &HorizontalPodAutoscalerClient{
		apiClient: apiClient,
	}
}

func (c *HorizontalPodAutoscalerClient) Create(container *HorizontalPodAutoscaler) (*HorizontalPodAutoscaler, error) {
	resp := &HorizontalPodAutoscaler{}
	err := c.apiClient.Ops.DoCreate(HorizontalPodAutoscalerType, container, resp)
	return resp, err
}

func (c *HorizontalPodAutoscalerClient) Update(existing *HorizontalPodAutoscaler, updates interface{}) (*HorizontalPodAutoscaler, error) {
	resp := &HorizontalPodAutoscaler{}
	err := c.apiClient.Ops.DoUpdate(HorizontalPodAutoscalerType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HorizontalPodAutoscalerClient) Replace(obj *HorizontalPodAutoscaler) (*HorizontalPodAutoscaler, error) {
	resp := &HorizontalPodAutoscaler{}
	err := c.apiClient.Ops.DoReplace(HorizontalPodAutoscalerType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *HorizontalPodAutoscalerClient) List(opts *types.ListOpts) (*HorizontalPodAutoscalerCollection, error) {
	resp := &HorizontalPodAutoscalerCollection{}
	err := c.apiClient.Ops.DoList(HorizontalPodAutoscalerType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HorizontalPodAutoscalerCollection) Next() (*HorizontalPodAutoscalerCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HorizontalPodAutoscalerCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HorizontalPodAutoscalerClient) ByID(id string) (*HorizontalPodAutoscaler, error) {
	resp := &HorizontalPodAutoscaler{}
	err := c.apiClient.Ops.DoByID(HorizontalPodAutoscalerType, id, resp)
	return resp, err
}

func (c *HorizontalPodAutoscalerClient) Delete(container *HorizontalPodAutoscaler) error {
	return c.apiClient.Ops.DoResourceDelete(HorizontalPodAutoscalerType, &container.Resource)
}
