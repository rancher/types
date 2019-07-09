package client

import (
	"github.com/rancher/norman/types"
)

const (
	CronJobType                               = "cronJob"
	CronJobFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	CronJobFieldAnnotations                   = "annotations"
	CronJobFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	CronJobFieldContainers                    = "containers"
	CronJobFieldCreated                       = "created"
	CronJobFieldCreatorID                     = "creatorId"
	CronJobFieldCronJobConfig                 = "cronJobConfig"
	CronJobFieldCronJobStatus                 = "cronJobStatus"
	CronJobFieldDNSConfig                     = "dnsConfig"
	CronJobFieldDNSPolicy                     = "dnsPolicy"
	CronJobFieldFsgid                         = "fsgid"
	CronJobFieldGids                          = "gids"
	CronJobFieldHostAliases                   = "hostAliases"
	CronJobFieldHostIPC                       = "hostIPC"
	CronJobFieldHostNetwork                   = "hostNetwork"
	CronJobFieldHostPID                       = "hostPID"
	CronJobFieldHostname                      = "hostname"
	CronJobFieldImagePullSecrets              = "imagePullSecrets"
	CronJobFieldLabels                        = "labels"
	CronJobFieldName                          = "name"
	CronJobFieldNamespaceId                   = "namespaceId"
	CronJobFieldNodeID                        = "nodeId"
	CronJobFieldOwnerReferences               = "ownerReferences"
	CronJobFieldPriority                      = "priority"
	CronJobFieldPriorityClassName             = "priorityClassName"
	CronJobFieldProjectID                     = "projectId"
	CronJobFieldPublicEndpoints               = "publicEndpoints"
	CronJobFieldReadinessGates                = "readinessGates"
	CronJobFieldRemoved                       = "removed"
	CronJobFieldRestartPolicy                 = "restartPolicy"
	CronJobFieldRunAsGroup                    = "runAsGroup"
	CronJobFieldRunAsNonRoot                  = "runAsNonRoot"
	CronJobFieldRuntimeClassName              = "runtimeClassName"
	CronJobFieldSchedulerName                 = "schedulerName"
	CronJobFieldScheduling                    = "scheduling"
	CronJobFieldSelector                      = "selector"
	CronJobFieldServiceAccountName            = "serviceAccountName"
	CronJobFieldShareProcessNamespace         = "shareProcessNamespace"
	CronJobFieldState                         = "state"
	CronJobFieldSubdomain                     = "subdomain"
	CronJobFieldSysctls                       = "sysctls"
	CronJobFieldTTLSecondsAfterFinished       = "ttlSecondsAfterFinished"
	CronJobFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	CronJobFieldTransitioning                 = "transitioning"
	CronJobFieldTransitioningMessage          = "transitioningMessage"
	CronJobFieldUUID                          = "uuid"
	CronJobFieldUid                           = "uid"
	CronJobFieldVolumes                       = "volumes"
	CronJobFieldWorkloadAnnotations           = "workloadAnnotations"
	CronJobFieldWorkloadLabels                = "workloadLabels"
	CronJobFieldWorkloadMetrics               = "workloadMetrics"
)

type CronJob struct {
	types.Resource
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	Annotations                   map[string]string      `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AutomountServiceAccountToken  *bool                  `json:"automountServiceAccountToken,omitempty" yaml:"automount_service_account_token,omitempty"`
	Containers                    []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	Created                       string                 `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                     string                 `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	CronJobConfig                 *CronJobConfig         `json:"cronJobConfig,omitempty" yaml:"cron_job_config,omitempty"`
	CronJobStatus                 *CronJobStatus         `json:"cronJobStatus,omitempty" yaml:"cron_job_status,omitempty"`
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

type CronJobCollection struct {
	types.Collection
	Data   []CronJob `json:"data,omitempty"`
	client *CronJobClient
}

type CronJobClient struct {
	apiClient *Client
}

type CronJobOperations interface {
	List(opts *types.ListOpts) (*CronJobCollection, error)
	Create(opts *CronJob) (*CronJob, error)
	Update(existing *CronJob, updates interface{}) (*CronJob, error)
	Replace(existing *CronJob) (*CronJob, error)
	ByID(id string) (*CronJob, error)
	Delete(container *CronJob) error
}

func newCronJobClient(apiClient *Client) *CronJobClient {
	return &CronJobClient{
		apiClient: apiClient,
	}
}

func (c *CronJobClient) Create(container *CronJob) (*CronJob, error) {
	resp := &CronJob{}
	err := c.apiClient.Ops.DoCreate(CronJobType, container, resp)
	return resp, err
}

func (c *CronJobClient) Update(existing *CronJob, updates interface{}) (*CronJob, error) {
	resp := &CronJob{}
	err := c.apiClient.Ops.DoUpdate(CronJobType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *CronJobClient) Replace(obj *CronJob) (*CronJob, error) {
	resp := &CronJob{}
	err := c.apiClient.Ops.DoReplace(CronJobType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *CronJobClient) List(opts *types.ListOpts) (*CronJobCollection, error) {
	resp := &CronJobCollection{}
	err := c.apiClient.Ops.DoList(CronJobType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *CronJobCollection) Next() (*CronJobCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &CronJobCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *CronJobClient) ByID(id string) (*CronJob, error) {
	resp := &CronJob{}
	err := c.apiClient.Ops.DoByID(CronJobType, id, resp)
	return resp, err
}

func (c *CronJobClient) Delete(container *CronJob) error {
	return c.apiClient.Ops.DoResourceDelete(CronJobType, &container.Resource)
}
