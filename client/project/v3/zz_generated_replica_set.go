package client

import (
	"github.com/rancher/norman/types"
)

const (
	ReplicaSetType                               = "replicaSet"
	ReplicaSetFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	ReplicaSetFieldAnnotations                   = "annotations"
	ReplicaSetFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	ReplicaSetFieldContainers                    = "containers"
	ReplicaSetFieldCreated                       = "created"
	ReplicaSetFieldCreatorID                     = "creatorId"
	ReplicaSetFieldDNSConfig                     = "dnsConfig"
	ReplicaSetFieldDNSPolicy                     = "dnsPolicy"
	ReplicaSetFieldFsgid                         = "fsgid"
	ReplicaSetFieldGids                          = "gids"
	ReplicaSetFieldHostAliases                   = "hostAliases"
	ReplicaSetFieldHostIPC                       = "hostIPC"
	ReplicaSetFieldHostNetwork                   = "hostNetwork"
	ReplicaSetFieldHostPID                       = "hostPID"
	ReplicaSetFieldHostname                      = "hostname"
	ReplicaSetFieldImagePullSecrets              = "imagePullSecrets"
	ReplicaSetFieldLabels                        = "labels"
	ReplicaSetFieldName                          = "name"
	ReplicaSetFieldNamespaceId                   = "namespaceId"
	ReplicaSetFieldNodeID                        = "nodeId"
	ReplicaSetFieldOwnerReferences               = "ownerReferences"
	ReplicaSetFieldPriority                      = "priority"
	ReplicaSetFieldPriorityClassName             = "priorityClassName"
	ReplicaSetFieldProjectID                     = "projectId"
	ReplicaSetFieldPublicEndpoints               = "publicEndpoints"
	ReplicaSetFieldReadinessGates                = "readinessGates"
	ReplicaSetFieldRemoved                       = "removed"
	ReplicaSetFieldReplicaSetConfig              = "replicaSetConfig"
	ReplicaSetFieldReplicaSetStatus              = "replicaSetStatus"
	ReplicaSetFieldRestartPolicy                 = "restartPolicy"
	ReplicaSetFieldRunAsGroup                    = "runAsGroup"
	ReplicaSetFieldRunAsNonRoot                  = "runAsNonRoot"
	ReplicaSetFieldRuntimeClassName              = "runtimeClassName"
	ReplicaSetFieldScale                         = "scale"
	ReplicaSetFieldSchedulerName                 = "schedulerName"
	ReplicaSetFieldScheduling                    = "scheduling"
	ReplicaSetFieldSelector                      = "selector"
	ReplicaSetFieldServiceAccountName            = "serviceAccountName"
	ReplicaSetFieldShareProcessNamespace         = "shareProcessNamespace"
	ReplicaSetFieldState                         = "state"
	ReplicaSetFieldSubdomain                     = "subdomain"
	ReplicaSetFieldSysctls                       = "sysctls"
	ReplicaSetFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	ReplicaSetFieldTransitioning                 = "transitioning"
	ReplicaSetFieldTransitioningMessage          = "transitioningMessage"
	ReplicaSetFieldUUID                          = "uuid"
	ReplicaSetFieldUid                           = "uid"
	ReplicaSetFieldVolumes                       = "volumes"
	ReplicaSetFieldWorkloadAnnotations           = "workloadAnnotations"
	ReplicaSetFieldWorkloadLabels                = "workloadLabels"
	ReplicaSetFieldWorkloadMetrics               = "workloadMetrics"
)

type ReplicaSet struct {
	types.Resource
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	Annotations                   map[string]string      `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AutomountServiceAccountToken  *bool                  `json:"automountServiceAccountToken,omitempty" yaml:"automount_service_account_token,omitempty"`
	Containers                    []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	Created                       string                 `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                     string                 `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DNSConfig                     *PodDNSConfig          `json:"dnsConfig,omitempty" yaml:"dns_config,omitempty"`
	DNSPolicy                     string                 `json:"dnsPolicy,omitempty" yaml:"dns_policy,omitempty"`
	Fsgid                         *int64                 `json:"fsgid,omitempty" yaml:"fsgid,omitempty"`
	Gids                          []int64                `json:"gids,omitempty" yaml:"gids,omitempty"`
	HostAliases                   []HostAlias            `json:"hostAliases,omitempty" yaml:"host_aliases,omitempty"`
	HostIPC                       bool                   `json:"hostIPC,omitempty" yaml:"host_ipc,omitempty"`
	HostNetwork                   bool                   `json:"hostNetwork,omitempty" yaml:"host_network,omitempty"`
	HostPID                       bool                   `json:"hostPID,omitempty" yaml:"host_pid,omitempty"`
	Hostname                      string                 `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	ImagePullSecrets              []LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"image_pull_secrets,omitempty"`
	Labels                        map[string]string      `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                          string                 `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId                   string                 `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	NodeID                        string                 `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	OwnerReferences               []OwnerReference       `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Priority                      *int64                 `json:"priority,omitempty" yaml:"priority,omitempty"`
	PriorityClassName             string                 `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	ProjectID                     string                 `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	PublicEndpoints               []PublicEndpoint       `json:"publicEndpoints,omitempty" yaml:"public_endpoints,omitempty"`
	ReadinessGates                []PodReadinessGate     `json:"readinessGates,omitempty" yaml:"readiness_gates,omitempty"`
	Removed                       string                 `json:"removed,omitempty" yaml:"removed,omitempty"`
	ReplicaSetConfig              *ReplicaSetConfig      `json:"replicaSetConfig,omitempty" yaml:"replica_set_config,omitempty"`
	ReplicaSetStatus              *ReplicaSetStatus      `json:"replicaSetStatus,omitempty" yaml:"replica_set_status,omitempty"`
	RestartPolicy                 string                 `json:"restartPolicy,omitempty" yaml:"restart_policy,omitempty"`
	RunAsGroup                    *int64                 `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot                  *bool                  `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	RuntimeClassName              string                 `json:"runtimeClassName,omitempty" yaml:"runtime_class_name,omitempty"`
	Scale                         *int64                 `json:"scale,omitempty" yaml:"scale,omitempty"`
	SchedulerName                 string                 `json:"schedulerName,omitempty" yaml:"scheduler_name,omitempty"`
	Scheduling                    *Scheduling            `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	Selector                      *LabelSelector         `json:"selector,omitempty" yaml:"selector,omitempty"`
	ServiceAccountName            string                 `json:"serviceAccountName,omitempty" yaml:"service_account_name,omitempty"`
	ShareProcessNamespace         *bool                  `json:"shareProcessNamespace,omitempty" yaml:"share_process_namespace,omitempty"`
	State                         string                 `json:"state,omitempty" yaml:"state,omitempty"`
	Subdomain                     string                 `json:"subdomain,omitempty" yaml:"subdomain,omitempty"`
	Sysctls                       []Sysctl               `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	TerminationGracePeriodSeconds *int64                 `json:"terminationGracePeriodSeconds,omitempty" yaml:"termination_grace_period_seconds,omitempty"`
	Transitioning                 string                 `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage          string                 `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                          string                 `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Uid                           *int64                 `json:"uid,omitempty" yaml:"uid,omitempty"`
	Volumes                       []Volume               `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	WorkloadAnnotations           map[string]string      `json:"workloadAnnotations,omitempty" yaml:"workload_annotations,omitempty"`
	WorkloadLabels                map[string]string      `json:"workloadLabels,omitempty" yaml:"workload_labels,omitempty"`
	WorkloadMetrics               []WorkloadMetric       `json:"workloadMetrics,omitempty" yaml:"workload_metrics,omitempty"`
}

type ReplicaSetCollection struct {
	types.Collection
	Data   []ReplicaSet `json:"data,omitempty"`
	client *ReplicaSetClient
}

type ReplicaSetClient struct {
	apiClient *Client
}

type ReplicaSetOperations interface {
	List(opts *types.ListOpts) (*ReplicaSetCollection, error)
	Create(opts *ReplicaSet) (*ReplicaSet, error)
	Update(existing *ReplicaSet, updates interface{}) (*ReplicaSet, error)
	Replace(existing *ReplicaSet) (*ReplicaSet, error)
	ByID(id string) (*ReplicaSet, error)
	Delete(container *ReplicaSet) error
}

func newReplicaSetClient(apiClient *Client) *ReplicaSetClient {
	return &ReplicaSetClient{
		apiClient: apiClient,
	}
}

func (c *ReplicaSetClient) Create(container *ReplicaSet) (*ReplicaSet, error) {
	resp := &ReplicaSet{}
	err := c.apiClient.Ops.DoCreate(ReplicaSetType, container, resp)
	return resp, err
}

func (c *ReplicaSetClient) Update(existing *ReplicaSet, updates interface{}) (*ReplicaSet, error) {
	resp := &ReplicaSet{}
	err := c.apiClient.Ops.DoUpdate(ReplicaSetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ReplicaSetClient) Replace(obj *ReplicaSet) (*ReplicaSet, error) {
	resp := &ReplicaSet{}
	err := c.apiClient.Ops.DoReplace(ReplicaSetType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ReplicaSetClient) List(opts *types.ListOpts) (*ReplicaSetCollection, error) {
	resp := &ReplicaSetCollection{}
	err := c.apiClient.Ops.DoList(ReplicaSetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ReplicaSetCollection) Next() (*ReplicaSetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ReplicaSetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ReplicaSetClient) ByID(id string) (*ReplicaSet, error) {
	resp := &ReplicaSet{}
	err := c.apiClient.Ops.DoByID(ReplicaSetType, id, resp)
	return resp, err
}

func (c *ReplicaSetClient) Delete(container *ReplicaSet) error {
	return c.apiClient.Ops.DoResourceDelete(ReplicaSetType, &container.Resource)
}
