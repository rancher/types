package client

import (
	"github.com/rancher/norman/types"
)

const (
	ServiceType                          = "service"
	ServiceFieldAnnotations              = "annotations"
	ServiceFieldClusterIp                = "clusterIp"
	ServiceFieldCreated                  = "created"
	ServiceFieldCreatorID                = "creatorId"
	ServiceFieldDescription              = "description"
	ServiceFieldExternalIPs              = "externalIPs"
	ServiceFieldExternalTrafficPolicy    = "externalTrafficPolicy"
	ServiceFieldHealthCheckNodePort      = "healthCheckNodePort"
	ServiceFieldHostname                 = "hostname"
	ServiceFieldIPAddresses              = "ipAddresses"
	ServiceFieldKind                     = "kind"
	ServiceFieldLabels                   = "labels"
	ServiceFieldLoadBalancerIP           = "loadBalancerIP"
	ServiceFieldLoadBalancerSourceRanges = "loadBalancerSourceRanges"
	ServiceFieldName                     = "name"
	ServiceFieldNamespaceId              = "namespaceId"
	ServiceFieldOwnerReferences          = "ownerReferences"
	ServiceFieldPorts                    = "ports"
	ServiceFieldProjectID                = "projectId"
	ServiceFieldPublicEndpoints          = "publicEndpoints"
	ServiceFieldPublishNotReadyAddresses = "publishNotReadyAddresses"
	ServiceFieldRemoved                  = "removed"
	ServiceFieldSelector                 = "selector"
	ServiceFieldSessionAffinity          = "sessionAffinity"
	ServiceFieldSessionAffinityConfig    = "sessionAffinityConfig"
	ServiceFieldState                    = "state"
	ServiceFieldTargetDNSRecordIDs       = "targetDnsRecordIds"
	ServiceFieldTargetWorkloadIDs        = "targetWorkloadIds"
	ServiceFieldTransitioning            = "transitioning"
	ServiceFieldTransitioningMessage     = "transitioningMessage"
	ServiceFieldUUID                     = "uuid"
	ServiceFieldWorkloadID               = "workloadId"
)

type Service struct {
	types.Resource
	Annotations              map[string]string      `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterIp                string                 `json:"clusterIp,omitempty" yaml:"cluster_ip,omitempty"`
	Created                  string                 `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                string                 `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Description              string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalIPs              []string               `json:"externalIPs,omitempty" yaml:"external_i_ps,omitempty"`
	ExternalTrafficPolicy    string                 `json:"externalTrafficPolicy,omitempty" yaml:"external_traffic_policy,omitempty"`
	HealthCheckNodePort      int64                  `json:"healthCheckNodePort,omitempty" yaml:"health_check_node_port,omitempty"`
	Hostname                 string                 `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	IPAddresses              []string               `json:"ipAddresses,omitempty" yaml:"ip_addresses,omitempty"`
	Kind                     string                 `json:"kind,omitempty" yaml:"kind,omitempty"`
	Labels                   map[string]string      `json:"labels,omitempty" yaml:"labels,omitempty"`
	LoadBalancerIP           string                 `json:"loadBalancerIP,omitempty" yaml:"load_balancer_ip,omitempty"`
	LoadBalancerSourceRanges []string               `json:"loadBalancerSourceRanges,omitempty" yaml:"load_balancer_source_ranges,omitempty"`
	Name                     string                 `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId              string                 `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	OwnerReferences          []OwnerReference       `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Ports                    []ServicePort          `json:"ports,omitempty" yaml:"ports,omitempty"`
	ProjectID                string                 `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	PublicEndpoints          []PublicEndpoint       `json:"publicEndpoints,omitempty" yaml:"public_endpoints,omitempty"`
	PublishNotReadyAddresses bool                   `json:"publishNotReadyAddresses,omitempty" yaml:"publish_not_ready_addresses,omitempty"`
	Removed                  string                 `json:"removed,omitempty" yaml:"removed,omitempty"`
	Selector                 map[string]string      `json:"selector,omitempty" yaml:"selector,omitempty"`
	SessionAffinity          string                 `json:"sessionAffinity,omitempty" yaml:"session_affinity,omitempty"`
	SessionAffinityConfig    *SessionAffinityConfig `json:"sessionAffinityConfig,omitempty" yaml:"session_affinity_config,omitempty"`
	State                    string                 `json:"state,omitempty" yaml:"state,omitempty"`
	TargetDNSRecordIDs       []string               `json:"targetDnsRecordIds,omitempty" yaml:"target_dns_record_ids,omitempty"`
	TargetWorkloadIDs        []string               `json:"targetWorkloadIds,omitempty" yaml:"target_workload_ids,omitempty"`
	Transitioning            string                 `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage     string                 `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                     string                 `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	WorkloadID               string                 `json:"workloadId,omitempty" yaml:"workload_id,omitempty"`
}

type ServiceCollection struct {
	types.Collection
	Data   []Service `json:"data,omitempty"`
	client *ServiceClient
}

type ServiceClient struct {
	apiClient *Client
}

type ServiceOperations interface {
	List(opts *types.ListOpts) (*ServiceCollection, error)
	Create(opts *Service) (*Service, error)
	Update(existing *Service, updates interface{}) (*Service, error)
	Replace(existing *Service) (*Service, error)
	ByID(id string) (*Service, error)
	Delete(container *Service) error
}

func newServiceClient(apiClient *Client) *ServiceClient {
	return &ServiceClient{
		apiClient: apiClient,
	}
}

func (c *ServiceClient) Create(container *Service) (*Service, error) {
	resp := &Service{}
	err := c.apiClient.Ops.DoCreate(ServiceType, container, resp)
	return resp, err
}

func (c *ServiceClient) Update(existing *Service, updates interface{}) (*Service, error) {
	resp := &Service{}
	err := c.apiClient.Ops.DoUpdate(ServiceType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceClient) Replace(obj *Service) (*Service, error) {
	resp := &Service{}
	err := c.apiClient.Ops.DoReplace(ServiceType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ServiceClient) List(opts *types.ListOpts) (*ServiceCollection, error) {
	resp := &ServiceCollection{}
	err := c.apiClient.Ops.DoList(ServiceType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceCollection) Next() (*ServiceCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceClient) ByID(id string) (*Service, error) {
	resp := &Service{}
	err := c.apiClient.Ops.DoByID(ServiceType, id, resp)
	return resp, err
}

func (c *ServiceClient) Delete(container *Service) error {
	return c.apiClient.Ops.DoResourceDelete(ServiceType, &container.Resource)
}
