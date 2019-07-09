package client

import (
	"github.com/rancher/norman/types"
)

const (
	WorkloadType                               = "workload"
	WorkloadFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	WorkloadFieldAnnotations                   = "annotations"
	WorkloadFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	WorkloadFieldContainers                    = "containers"
	WorkloadFieldCreated                       = "created"
	WorkloadFieldCreatorID                     = "creatorId"
	WorkloadFieldCronJobConfig                 = "cronJobConfig"
	WorkloadFieldCronJobStatus                 = "cronJobStatus"
	WorkloadFieldDNSConfig                     = "dnsConfig"
	WorkloadFieldDNSPolicy                     = "dnsPolicy"
	WorkloadFieldDaemonSetConfig               = "daemonSetConfig"
	WorkloadFieldDaemonSetStatus               = "daemonSetStatus"
	WorkloadFieldDeploymentConfig              = "deploymentConfig"
	WorkloadFieldDeploymentStatus              = "deploymentStatus"
	WorkloadFieldFsgid                         = "fsgid"
	WorkloadFieldGids                          = "gids"
	WorkloadFieldHostAliases                   = "hostAliases"
	WorkloadFieldHostIPC                       = "hostIPC"
	WorkloadFieldHostNetwork                   = "hostNetwork"
	WorkloadFieldHostPID                       = "hostPID"
	WorkloadFieldHostname                      = "hostname"
	WorkloadFieldImagePullSecrets              = "imagePullSecrets"
	WorkloadFieldJobConfig                     = "jobConfig"
	WorkloadFieldJobStatus                     = "jobStatus"
	WorkloadFieldLabels                        = "labels"
	WorkloadFieldName                          = "name"
	WorkloadFieldNamespaceId                   = "namespaceId"
	WorkloadFieldNodeID                        = "nodeId"
	WorkloadFieldOwnerReferences               = "ownerReferences"
	WorkloadFieldPaused                        = "paused"
	WorkloadFieldPriority                      = "priority"
	WorkloadFieldPriorityClassName             = "priorityClassName"
	WorkloadFieldProjectID                     = "projectId"
	WorkloadFieldPublicEndpoints               = "publicEndpoints"
	WorkloadFieldReadinessGates                = "readinessGates"
	WorkloadFieldRemoved                       = "removed"
	WorkloadFieldReplicaSetConfig              = "replicaSetConfig"
	WorkloadFieldReplicaSetStatus              = "replicaSetStatus"
	WorkloadFieldReplicationControllerConfig   = "replicationControllerConfig"
	WorkloadFieldReplicationControllerStatus   = "replicationControllerStatus"
	WorkloadFieldRestartPolicy                 = "restartPolicy"
	WorkloadFieldRunAsGroup                    = "runAsGroup"
	WorkloadFieldRunAsNonRoot                  = "runAsNonRoot"
	WorkloadFieldRuntimeClassName              = "runtimeClassName"
	WorkloadFieldScale                         = "scale"
	WorkloadFieldSchedulerName                 = "schedulerName"
	WorkloadFieldScheduling                    = "scheduling"
	WorkloadFieldSelector                      = "selector"
	WorkloadFieldServiceAccountName            = "serviceAccountName"
	WorkloadFieldShareProcessNamespace         = "shareProcessNamespace"
	WorkloadFieldState                         = "state"
	WorkloadFieldStatefulSetConfig             = "statefulSetConfig"
	WorkloadFieldStatefulSetStatus             = "statefulSetStatus"
	WorkloadFieldSubdomain                     = "subdomain"
	WorkloadFieldSysctls                       = "sysctls"
	WorkloadFieldTTLSecondsAfterFinished       = "ttlSecondsAfterFinished"
	WorkloadFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	WorkloadFieldTransitioning                 = "transitioning"
	WorkloadFieldTransitioningMessage          = "transitioningMessage"
	WorkloadFieldUUID                          = "uuid"
	WorkloadFieldUid                           = "uid"
	WorkloadFieldVolumes                       = "volumes"
	WorkloadFieldWorkloadAnnotations           = "workloadAnnotations"
	WorkloadFieldWorkloadLabels                = "workloadLabels"
	WorkloadFieldWorkloadMetrics               = "workloadMetrics"
)

type Workload struct {
	types.Resource
	ActiveDeadlineSeconds         *int64                       `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	Annotations                   map[string]string            `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AutomountServiceAccountToken  *bool                        `json:"automountServiceAccountToken,omitempty" yaml:"automount_service_account_token,omitempty"`
	Containers                    []Container                  `json:"containers,omitempty" yaml:"containers,omitempty"`
	Created                       string                       `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                     string                       `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	CronJobConfig                 *CronJobConfig               `json:"cronJobConfig,omitempty" yaml:"cron_job_config,omitempty"`
	CronJobStatus                 *CronJobStatus               `json:"cronJobStatus,omitempty" yaml:"cron_job_status,omitempty"`
	DNSConfig                     *PodDNSConfig                `json:"dnsConfig,omitempty" yaml:"dns_config,omitempty"`
	DNSPolicy                     string                       `json:"dnsPolicy,omitempty" yaml:"dns_policy,omitempty"`
	DaemonSetConfig               *DaemonSetConfig             `json:"daemonSetConfig,omitempty" yaml:"daemon_set_config,omitempty"`
	DaemonSetStatus               *DaemonSetStatus             `json:"daemonSetStatus,omitempty" yaml:"daemon_set_status,omitempty"`
	DeploymentConfig              *DeploymentConfig            `json:"deploymentConfig,omitempty" yaml:"deployment_config,omitempty"`
	DeploymentStatus              *DeploymentStatus            `json:"deploymentStatus,omitempty" yaml:"deployment_status,omitempty"`
	Fsgid                         *int64                       `json:"fsgid,omitempty" yaml:"fsgid,omitempty"`
	Gids                          []int64                      `json:"gids,omitempty" yaml:"gids,omitempty"`
	HostAliases                   []HostAlias                  `json:"hostAliases,omitempty" yaml:"host_aliases,omitempty"`
	HostIPC                       bool                         `json:"hostIPC,omitempty" yaml:"host_ipc,omitempty"`
	HostNetwork                   bool                         `json:"hostNetwork,omitempty" yaml:"host_network,omitempty"`
	HostPID                       bool                         `json:"hostPID,omitempty" yaml:"host_pid,omitempty"`
	Hostname                      string                       `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	ImagePullSecrets              []LocalObjectReference       `json:"imagePullSecrets,omitempty" yaml:"image_pull_secrets,omitempty"`
	JobConfig                     *JobConfig                   `json:"jobConfig,omitempty" yaml:"job_config,omitempty"`
	JobStatus                     *JobStatus                   `json:"jobStatus,omitempty" yaml:"job_status,omitempty"`
	Labels                        map[string]string            `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                          string                       `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId                   string                       `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	NodeID                        string                       `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	OwnerReferences               []OwnerReference             `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Paused                        bool                         `json:"paused,omitempty" yaml:"paused,omitempty"`
	Priority                      *int64                       `json:"priority,omitempty" yaml:"priority,omitempty"`
	PriorityClassName             string                       `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	ProjectID                     string                       `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	PublicEndpoints               []PublicEndpoint             `json:"publicEndpoints,omitempty" yaml:"public_endpoints,omitempty"`
	ReadinessGates                []PodReadinessGate           `json:"readinessGates,omitempty" yaml:"readiness_gates,omitempty"`
	Removed                       string                       `json:"removed,omitempty" yaml:"removed,omitempty"`
	ReplicaSetConfig              *ReplicaSetConfig            `json:"replicaSetConfig,omitempty" yaml:"replica_set_config,omitempty"`
	ReplicaSetStatus              *ReplicaSetStatus            `json:"replicaSetStatus,omitempty" yaml:"replica_set_status,omitempty"`
	ReplicationControllerConfig   *ReplicationControllerConfig `json:"replicationControllerConfig,omitempty" yaml:"replication_controller_config,omitempty"`
	ReplicationControllerStatus   *ReplicationControllerStatus `json:"replicationControllerStatus,omitempty" yaml:"replication_controller_status,omitempty"`
	RestartPolicy                 string                       `json:"restartPolicy,omitempty" yaml:"restart_policy,omitempty"`
	RunAsGroup                    *int64                       `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot                  *bool                        `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	RuntimeClassName              string                       `json:"runtimeClassName,omitempty" yaml:"runtime_class_name,omitempty"`
	Scale                         *int64                       `json:"scale,omitempty" yaml:"scale,omitempty"`
	SchedulerName                 string                       `json:"schedulerName,omitempty" yaml:"scheduler_name,omitempty"`
	Scheduling                    *Scheduling                  `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	Selector                      *LabelSelector               `json:"selector,omitempty" yaml:"selector,omitempty"`
	ServiceAccountName            string                       `json:"serviceAccountName,omitempty" yaml:"service_account_name,omitempty"`
	ShareProcessNamespace         *bool                        `json:"shareProcessNamespace,omitempty" yaml:"share_process_namespace,omitempty"`
	State                         string                       `json:"state,omitempty" yaml:"state,omitempty"`
	StatefulSetConfig             *StatefulSetConfig           `json:"statefulSetConfig,omitempty" yaml:"stateful_set_config,omitempty"`
	StatefulSetStatus             *StatefulSetStatus           `json:"statefulSetStatus,omitempty" yaml:"stateful_set_status,omitempty"`
	Subdomain                     string                       `json:"subdomain,omitempty" yaml:"subdomain,omitempty"`
	Sysctls                       []Sysctl                     `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	TTLSecondsAfterFinished       *int64                       `json:"ttlSecondsAfterFinished,omitempty" yaml:"ttl_seconds_after_finished,omitempty"`
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

type WorkloadCollection struct {
	types.Collection
	Data   []Workload `json:"data,omitempty"`
	client *WorkloadClient
}

type WorkloadClient struct {
	apiClient *Client
}

type WorkloadOperations interface {
	List(opts *types.ListOpts) (*WorkloadCollection, error)
	Create(opts *Workload) (*Workload, error)
	Update(existing *Workload, updates interface{}) (*Workload, error)
	Replace(existing *Workload) (*Workload, error)
	ByID(id string) (*Workload, error)
	Delete(container *Workload) error

	ActionPause(resource *Workload) error

	ActionResume(resource *Workload) error

	ActionRollback(resource *Workload, input *RollbackRevision) error
}

func newWorkloadClient(apiClient *Client) *WorkloadClient {
	return &WorkloadClient{
		apiClient: apiClient,
	}
}

func (c *WorkloadClient) Create(container *Workload) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoCreate(WorkloadType, container, resp)
	return resp, err
}

func (c *WorkloadClient) Update(existing *Workload, updates interface{}) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoUpdate(WorkloadType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *WorkloadClient) Replace(obj *Workload) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoReplace(WorkloadType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *WorkloadClient) List(opts *types.ListOpts) (*WorkloadCollection, error) {
	resp := &WorkloadCollection{}
	err := c.apiClient.Ops.DoList(WorkloadType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *WorkloadCollection) Next() (*WorkloadCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &WorkloadCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *WorkloadClient) ByID(id string) (*Workload, error) {
	resp := &Workload{}
	err := c.apiClient.Ops.DoByID(WorkloadType, id, resp)
	return resp, err
}

func (c *WorkloadClient) Delete(container *Workload) error {
	return c.apiClient.Ops.DoResourceDelete(WorkloadType, &container.Resource)
}

func (c *WorkloadClient) ActionPause(resource *Workload) error {
	err := c.apiClient.Ops.DoAction(WorkloadType, "pause", &resource.Resource, nil, nil)
	return err
}

func (c *WorkloadClient) ActionResume(resource *Workload) error {
	err := c.apiClient.Ops.DoAction(WorkloadType, "resume", &resource.Resource, nil, nil)
	return err
}

func (c *WorkloadClient) ActionRollback(resource *Workload, input *RollbackRevision) error {
	err := c.apiClient.Ops.DoAction(WorkloadType, "rollback", &resource.Resource, input, nil)
	return err
}
