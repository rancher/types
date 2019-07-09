package client

import (
	"github.com/rancher/norman/types"
)

const (
	DaemonSetType                               = "daemonSet"
	DaemonSetFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	DaemonSetFieldAnnotations                   = "annotations"
	DaemonSetFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	DaemonSetFieldContainers                    = "containers"
	DaemonSetFieldCreated                       = "created"
	DaemonSetFieldCreatorID                     = "creatorId"
	DaemonSetFieldDNSConfig                     = "dnsConfig"
	DaemonSetFieldDNSPolicy                     = "dnsPolicy"
	DaemonSetFieldDaemonSetConfig               = "daemonSetConfig"
	DaemonSetFieldDaemonSetStatus               = "daemonSetStatus"
	DaemonSetFieldFsgid                         = "fsgid"
	DaemonSetFieldGids                          = "gids"
	DaemonSetFieldHostAliases                   = "hostAliases"
	DaemonSetFieldHostIPC                       = "hostIPC"
	DaemonSetFieldHostNetwork                   = "hostNetwork"
	DaemonSetFieldHostPID                       = "hostPID"
	DaemonSetFieldHostname                      = "hostname"
	DaemonSetFieldImagePullSecrets              = "imagePullSecrets"
	DaemonSetFieldLabels                        = "labels"
	DaemonSetFieldName                          = "name"
	DaemonSetFieldNamespaceId                   = "namespaceId"
	DaemonSetFieldNodeID                        = "nodeId"
	DaemonSetFieldOwnerReferences               = "ownerReferences"
	DaemonSetFieldPriority                      = "priority"
	DaemonSetFieldPriorityClassName             = "priorityClassName"
	DaemonSetFieldProjectID                     = "projectId"
	DaemonSetFieldPublicEndpoints               = "publicEndpoints"
	DaemonSetFieldReadinessGates                = "readinessGates"
	DaemonSetFieldRemoved                       = "removed"
	DaemonSetFieldRestartPolicy                 = "restartPolicy"
	DaemonSetFieldRunAsGroup                    = "runAsGroup"
	DaemonSetFieldRunAsNonRoot                  = "runAsNonRoot"
	DaemonSetFieldRuntimeClassName              = "runtimeClassName"
	DaemonSetFieldSchedulerName                 = "schedulerName"
	DaemonSetFieldScheduling                    = "scheduling"
	DaemonSetFieldSelector                      = "selector"
	DaemonSetFieldServiceAccountName            = "serviceAccountName"
	DaemonSetFieldShareProcessNamespace         = "shareProcessNamespace"
	DaemonSetFieldState                         = "state"
	DaemonSetFieldSubdomain                     = "subdomain"
	DaemonSetFieldSysctls                       = "sysctls"
	DaemonSetFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	DaemonSetFieldTransitioning                 = "transitioning"
	DaemonSetFieldTransitioningMessage          = "transitioningMessage"
	DaemonSetFieldUUID                          = "uuid"
	DaemonSetFieldUid                           = "uid"
	DaemonSetFieldVolumes                       = "volumes"
	DaemonSetFieldWorkloadAnnotations           = "workloadAnnotations"
	DaemonSetFieldWorkloadLabels                = "workloadLabels"
	DaemonSetFieldWorkloadMetrics               = "workloadMetrics"
)

type DaemonSet struct {
	types.Resource
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	Annotations                   map[string]string      `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AutomountServiceAccountToken  *bool                  `json:"automountServiceAccountToken,omitempty" yaml:"automount_service_account_token,omitempty"`
	Containers                    []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	Created                       string                 `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                     string                 `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DNSConfig                     *PodDNSConfig          `json:"dnsConfig,omitempty" yaml:"dns_config,omitempty"`
	DNSPolicy                     string                 `json:"dnsPolicy,omitempty" yaml:"dns_policy,omitempty"`
	DaemonSetConfig               *DaemonSetConfig       `json:"daemonSetConfig,omitempty" yaml:"daemon_set_config,omitempty"`
	DaemonSetStatus               *DaemonSetStatus       `json:"daemonSetStatus,omitempty" yaml:"daemon_set_status,omitempty"`
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
	RestartPolicy                 string                 `json:"restartPolicy,omitempty" yaml:"restart_policy,omitempty"`
	RunAsGroup                    *int64                 `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot                  *bool                  `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	RuntimeClassName              string                 `json:"runtimeClassName,omitempty" yaml:"runtime_class_name,omitempty"`
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

type DaemonSetCollection struct {
	types.Collection
	Data   []DaemonSet `json:"data,omitempty"`
	client *DaemonSetClient
}

type DaemonSetClient struct {
	apiClient *Client
}

type DaemonSetOperations interface {
	List(opts *types.ListOpts) (*DaemonSetCollection, error)
	Create(opts *DaemonSet) (*DaemonSet, error)
	Update(existing *DaemonSet, updates interface{}) (*DaemonSet, error)
	Replace(existing *DaemonSet) (*DaemonSet, error)
	ByID(id string) (*DaemonSet, error)
	Delete(container *DaemonSet) error
}

func newDaemonSetClient(apiClient *Client) *DaemonSetClient {
	return &DaemonSetClient{
		apiClient: apiClient,
	}
}

func (c *DaemonSetClient) Create(container *DaemonSet) (*DaemonSet, error) {
	resp := &DaemonSet{}
	err := c.apiClient.Ops.DoCreate(DaemonSetType, container, resp)
	return resp, err
}

func (c *DaemonSetClient) Update(existing *DaemonSet, updates interface{}) (*DaemonSet, error) {
	resp := &DaemonSet{}
	err := c.apiClient.Ops.DoUpdate(DaemonSetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DaemonSetClient) Replace(obj *DaemonSet) (*DaemonSet, error) {
	resp := &DaemonSet{}
	err := c.apiClient.Ops.DoReplace(DaemonSetType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *DaemonSetClient) List(opts *types.ListOpts) (*DaemonSetCollection, error) {
	resp := &DaemonSetCollection{}
	err := c.apiClient.Ops.DoList(DaemonSetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DaemonSetCollection) Next() (*DaemonSetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DaemonSetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DaemonSetClient) ByID(id string) (*DaemonSet, error) {
	resp := &DaemonSet{}
	err := c.apiClient.Ops.DoByID(DaemonSetType, id, resp)
	return resp, err
}

func (c *DaemonSetClient) Delete(container *DaemonSet) error {
	return c.apiClient.Ops.DoResourceDelete(DaemonSetType, &container.Resource)
}
