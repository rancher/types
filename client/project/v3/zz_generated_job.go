package client

import (
	"github.com/rancher/norman/types"
)

const (
	JobType                               = "job"
	JobFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	JobFieldAnnotations                   = "annotations"
	JobFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	JobFieldContainers                    = "containers"
	JobFieldCreated                       = "created"
	JobFieldCreatorID                     = "creatorId"
	JobFieldDNSConfig                     = "dnsConfig"
	JobFieldDNSPolicy                     = "dnsPolicy"
	JobFieldFsgid                         = "fsgid"
	JobFieldGids                          = "gids"
	JobFieldHostAliases                   = "hostAliases"
	JobFieldHostIPC                       = "hostIPC"
	JobFieldHostNetwork                   = "hostNetwork"
	JobFieldHostPID                       = "hostPID"
	JobFieldHostname                      = "hostname"
	JobFieldImagePullSecrets              = "imagePullSecrets"
	JobFieldJobConfig                     = "jobConfig"
	JobFieldJobStatus                     = "jobStatus"
	JobFieldLabels                        = "labels"
	JobFieldName                          = "name"
	JobFieldNamespaceId                   = "namespaceId"
	JobFieldNodeID                        = "nodeId"
	JobFieldOwnerReferences               = "ownerReferences"
	JobFieldPriority                      = "priority"
	JobFieldPriorityClassName             = "priorityClassName"
	JobFieldProjectID                     = "projectId"
	JobFieldPublicEndpoints               = "publicEndpoints"
	JobFieldReadinessGates                = "readinessGates"
	JobFieldRemoved                       = "removed"
	JobFieldRestartPolicy                 = "restartPolicy"
	JobFieldRunAsGroup                    = "runAsGroup"
	JobFieldRunAsNonRoot                  = "runAsNonRoot"
	JobFieldRuntimeClassName              = "runtimeClassName"
	JobFieldSchedulerName                 = "schedulerName"
	JobFieldScheduling                    = "scheduling"
	JobFieldSelector                      = "selector"
	JobFieldServiceAccountName            = "serviceAccountName"
	JobFieldShareProcessNamespace         = "shareProcessNamespace"
	JobFieldState                         = "state"
	JobFieldSubdomain                     = "subdomain"
	JobFieldSysctls                       = "sysctls"
	JobFieldTTLSecondsAfterFinished       = "ttlSecondsAfterFinished"
	JobFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	JobFieldTransitioning                 = "transitioning"
	JobFieldTransitioningMessage          = "transitioningMessage"
	JobFieldUUID                          = "uuid"
	JobFieldUid                           = "uid"
	JobFieldVolumes                       = "volumes"
	JobFieldWorkloadAnnotations           = "workloadAnnotations"
	JobFieldWorkloadLabels                = "workloadLabels"
	JobFieldWorkloadMetrics               = "workloadMetrics"
)

type Job struct {
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
	JobConfig                     *JobConfig             `json:"jobConfig,omitempty" yaml:"job_config,omitempty"`
	JobStatus                     *JobStatus             `json:"jobStatus,omitempty" yaml:"job_status,omitempty"`
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
	TTLSecondsAfterFinished       *int64                 `json:"ttlSecondsAfterFinished,omitempty" yaml:"ttl_seconds_after_finished,omitempty"`
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

type JobCollection struct {
	types.Collection
	Data   []Job `json:"data,omitempty"`
	client *JobClient
}

type JobClient struct {
	apiClient *Client
}

type JobOperations interface {
	List(opts *types.ListOpts) (*JobCollection, error)
	Create(opts *Job) (*Job, error)
	Update(existing *Job, updates interface{}) (*Job, error)
	Replace(existing *Job) (*Job, error)
	ByID(id string) (*Job, error)
	Delete(container *Job) error
}

func newJobClient(apiClient *Client) *JobClient {
	return &JobClient{
		apiClient: apiClient,
	}
}

func (c *JobClient) Create(container *Job) (*Job, error) {
	resp := &Job{}
	err := c.apiClient.Ops.DoCreate(JobType, container, resp)
	return resp, err
}

func (c *JobClient) Update(existing *Job, updates interface{}) (*Job, error) {
	resp := &Job{}
	err := c.apiClient.Ops.DoUpdate(JobType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *JobClient) Replace(obj *Job) (*Job, error) {
	resp := &Job{}
	err := c.apiClient.Ops.DoReplace(JobType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *JobClient) List(opts *types.ListOpts) (*JobCollection, error) {
	resp := &JobCollection{}
	err := c.apiClient.Ops.DoList(JobType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *JobCollection) Next() (*JobCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &JobCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *JobClient) ByID(id string) (*Job, error) {
	resp := &Job{}
	err := c.apiClient.Ops.DoByID(JobType, id, resp)
	return resp, err
}

func (c *JobClient) Delete(container *Job) error {
	return c.apiClient.Ops.DoResourceDelete(JobType, &container.Resource)
}
