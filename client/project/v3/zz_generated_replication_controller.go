package client

import (
	"github.com/rancher/norman/types"
)

const (
	ReplicationControllerType                               = "replicationController"
	ReplicationControllerFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	ReplicationControllerFieldAnnotations                   = "annotations"
	ReplicationControllerFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	ReplicationControllerFieldContainers                    = "containers"
	ReplicationControllerFieldCreated                       = "created"
	ReplicationControllerFieldCreatorID                     = "creatorId"
	ReplicationControllerFieldDNSConfig                     = "dnsConfig"
	ReplicationControllerFieldDNSPolicy                     = "dnsPolicy"
	ReplicationControllerFieldFsgid                         = "fsgid"
	ReplicationControllerFieldGids                          = "gids"
	ReplicationControllerFieldHostAliases                   = "hostAliases"
	ReplicationControllerFieldHostIPC                       = "hostIPC"
	ReplicationControllerFieldHostNetwork                   = "hostNetwork"
	ReplicationControllerFieldHostPID                       = "hostPID"
	ReplicationControllerFieldHostname                      = "hostname"
	ReplicationControllerFieldImagePullSecrets              = "imagePullSecrets"
	ReplicationControllerFieldLabels                        = "labels"
	ReplicationControllerFieldName                          = "name"
	ReplicationControllerFieldNamespaceId                   = "namespaceId"
	ReplicationControllerFieldNodeID                        = "nodeId"
	ReplicationControllerFieldOwnerReferences               = "ownerReferences"
	ReplicationControllerFieldPriority                      = "priority"
	ReplicationControllerFieldPriorityClassName             = "priorityClassName"
	ReplicationControllerFieldProjectID                     = "projectId"
	ReplicationControllerFieldPublicEndpoints               = "publicEndpoints"
	ReplicationControllerFieldReadinessGates                = "readinessGates"
	ReplicationControllerFieldRemoved                       = "removed"
	ReplicationControllerFieldReplicationControllerConfig   = "replicationControllerConfig"
	ReplicationControllerFieldReplicationControllerStatus   = "replicationControllerStatus"
	ReplicationControllerFieldRestartPolicy                 = "restartPolicy"
	ReplicationControllerFieldRunAsGroup                    = "runAsGroup"
	ReplicationControllerFieldRunAsNonRoot                  = "runAsNonRoot"
	ReplicationControllerFieldRuntimeClassName              = "runtimeClassName"
	ReplicationControllerFieldScale                         = "scale"
	ReplicationControllerFieldSchedulerName                 = "schedulerName"
	ReplicationControllerFieldScheduling                    = "scheduling"
	ReplicationControllerFieldSelector                      = "selector"
	ReplicationControllerFieldServiceAccountName            = "serviceAccountName"
	ReplicationControllerFieldShareProcessNamespace         = "shareProcessNamespace"
	ReplicationControllerFieldState                         = "state"
	ReplicationControllerFieldSubdomain                     = "subdomain"
	ReplicationControllerFieldSysctls                       = "sysctls"
	ReplicationControllerFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	ReplicationControllerFieldTransitioning                 = "transitioning"
	ReplicationControllerFieldTransitioningMessage          = "transitioningMessage"
	ReplicationControllerFieldUUID                          = "uuid"
	ReplicationControllerFieldUid                           = "uid"
	ReplicationControllerFieldVolumes                       = "volumes"
	ReplicationControllerFieldWorkloadAnnotations           = "workloadAnnotations"
	ReplicationControllerFieldWorkloadLabels                = "workloadLabels"
	ReplicationControllerFieldWorkloadMetrics               = "workloadMetrics"
)

type ReplicationController struct {
	types.Resource
	ActiveDeadlineSeconds         *int64                       `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	Annotations                   map[string]string            `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AutomountServiceAccountToken  *bool                        `json:"automountServiceAccountToken,omitempty" yaml:"automount_service_account_token,omitempty"`
	Containers                    []Container                  `json:"containers,omitempty" yaml:"containers,omitempty"`
	Created                       string                       `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                     string                       `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DNSConfig                     *PodDNSConfig                `json:"dnsConfig,omitempty" yaml:"dns_config,omitempty"`
	DNSPolicy                     string                       `json:"dnsPolicy,omitempty" yaml:"dns_policy,omitempty"`
	Fsgid                         *int64                       `json:"fsgid,omitempty" yaml:"fsgid,omitempty"`
	Gids                          []int64                      `json:"gids,omitempty" yaml:"gids,omitempty"`
	HostAliases                   []HostAlias                  `json:"hostAliases,omitempty" yaml:"host_aliases,omitempty"`
	HostIPC                       bool                         `json:"hostIPC,omitempty" yaml:"host_ipc,omitempty"`
	HostNetwork                   bool                         `json:"hostNetwork,omitempty" yaml:"host_network,omitempty"`
	HostPID                       bool                         `json:"hostPID,omitempty" yaml:"host_pid,omitempty"`
	Hostname                      string                       `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	ImagePullSecrets              []LocalObjectReference       `json:"imagePullSecrets,omitempty" yaml:"image_pull_secrets,omitempty"`
	Labels                        map[string]string            `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                          string                       `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId                   string                       `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	NodeID                        string                       `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	OwnerReferences               []OwnerReference             `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Priority                      *int64                       `json:"priority,omitempty" yaml:"priority,omitempty"`
	PriorityClassName             string                       `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	ProjectID                     string                       `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	PublicEndpoints               []PublicEndpoint             `json:"publicEndpoints,omitempty" yaml:"public_endpoints,omitempty"`
	ReadinessGates                []PodReadinessGate           `json:"readinessGates,omitempty" yaml:"readiness_gates,omitempty"`
	Removed                       string                       `json:"removed,omitempty" yaml:"removed,omitempty"`
	ReplicationControllerConfig   *ReplicationControllerConfig `json:"replicationControllerConfig,omitempty" yaml:"replication_controller_config,omitempty"`
	ReplicationControllerStatus   *ReplicationControllerStatus `json:"replicationControllerStatus,omitempty" yaml:"replication_controller_status,omitempty"`
	RestartPolicy                 string                       `json:"restartPolicy,omitempty" yaml:"restart_policy,omitempty"`
	RunAsGroup                    *int64                       `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot                  *bool                        `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	RuntimeClassName              string                       `json:"runtimeClassName,omitempty" yaml:"runtime_class_name,omitempty"`
	Scale                         *int64                       `json:"scale,omitempty" yaml:"scale,omitempty"`
	SchedulerName                 string                       `json:"schedulerName,omitempty" yaml:"scheduler_name,omitempty"`
	Scheduling                    *Scheduling                  `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	Selector                      map[string]string            `json:"selector,omitempty" yaml:"selector,omitempty"`
	ServiceAccountName            string                       `json:"serviceAccountName,omitempty" yaml:"service_account_name,omitempty"`
	ShareProcessNamespace         *bool                        `json:"shareProcessNamespace,omitempty" yaml:"share_process_namespace,omitempty"`
	State                         string                       `json:"state,omitempty" yaml:"state,omitempty"`
	Subdomain                     string                       `json:"subdomain,omitempty" yaml:"subdomain,omitempty"`
	Sysctls                       []Sysctl                     `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	TerminationGracePeriodSeconds *int64                       `json:"terminationGracePeriodSeconds,omitempty" yaml:"termination_grace_period_seconds,omitempty"`
	Transitioning                 string                       `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage          string                       `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                          string                       `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Uid                           *int64                       `json:"uid,omitempty" yaml:"uid,omitempty"`
	Volumes                       []Volume                     `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	WorkloadAnnotations           map[string]string            `json:"workloadAnnotations,omitempty" yaml:"workload_annotations,omitempty"`
	WorkloadLabels                map[string]string            `json:"workloadLabels,omitempty" yaml:"workload_labels,omitempty"`
	WorkloadMetrics               []WorkloadMetric             `json:"workloadMetrics,omitempty" yaml:"workload_metrics,omitempty"`
}

type ReplicationControllerCollection struct {
	types.Collection
	Data   []ReplicationController `json:"data,omitempty"`
	client *ReplicationControllerClient
}

type ReplicationControllerClient struct {
	apiClient *Client
}

type ReplicationControllerOperations interface {
	List(opts *types.ListOpts) (*ReplicationControllerCollection, error)
	Create(opts *ReplicationController) (*ReplicationController, error)
	Update(existing *ReplicationController, updates interface{}) (*ReplicationController, error)
	Replace(existing *ReplicationController) (*ReplicationController, error)
	ByID(id string) (*ReplicationController, error)
	Delete(container *ReplicationController) error
}

func newReplicationControllerClient(apiClient *Client) *ReplicationControllerClient {
	return &ReplicationControllerClient{
		apiClient: apiClient,
	}
}

func (c *ReplicationControllerClient) Create(container *ReplicationController) (*ReplicationController, error) {
	resp := &ReplicationController{}
	err := c.apiClient.Ops.DoCreate(ReplicationControllerType, container, resp)
	return resp, err
}

func (c *ReplicationControllerClient) Update(existing *ReplicationController, updates interface{}) (*ReplicationController, error) {
	resp := &ReplicationController{}
	err := c.apiClient.Ops.DoUpdate(ReplicationControllerType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ReplicationControllerClient) Replace(obj *ReplicationController) (*ReplicationController, error) {
	resp := &ReplicationController{}
	err := c.apiClient.Ops.DoReplace(ReplicationControllerType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ReplicationControllerClient) List(opts *types.ListOpts) (*ReplicationControllerCollection, error) {
	resp := &ReplicationControllerCollection{}
	err := c.apiClient.Ops.DoList(ReplicationControllerType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ReplicationControllerCollection) Next() (*ReplicationControllerCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ReplicationControllerCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ReplicationControllerClient) ByID(id string) (*ReplicationController, error) {
	resp := &ReplicationController{}
	err := c.apiClient.Ops.DoByID(ReplicationControllerType, id, resp)
	return resp, err
}

func (c *ReplicationControllerClient) Delete(container *ReplicationController) error {
	return c.apiClient.Ops.DoResourceDelete(ReplicationControllerType, &container.Resource)
}
