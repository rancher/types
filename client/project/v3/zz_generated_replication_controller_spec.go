package client

const (
	ReplicationControllerSpecType                               = "replicationControllerSpec"
	ReplicationControllerSpecFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	ReplicationControllerSpecFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	ReplicationControllerSpecFieldContainers                    = "containers"
	ReplicationControllerSpecFieldDNSConfig                     = "dnsConfig"
	ReplicationControllerSpecFieldDNSPolicy                     = "dnsPolicy"
	ReplicationControllerSpecFieldFsgid                         = "fsgid"
	ReplicationControllerSpecFieldGids                          = "gids"
	ReplicationControllerSpecFieldHostAliases                   = "hostAliases"
	ReplicationControllerSpecFieldHostIPC                       = "hostIPC"
	ReplicationControllerSpecFieldHostNetwork                   = "hostNetwork"
	ReplicationControllerSpecFieldHostPID                       = "hostPID"
	ReplicationControllerSpecFieldHostname                      = "hostname"
	ReplicationControllerSpecFieldImagePullSecrets              = "imagePullSecrets"
	ReplicationControllerSpecFieldNodeID                        = "nodeId"
	ReplicationControllerSpecFieldObjectMeta                    = "metadata"
	ReplicationControllerSpecFieldPriority                      = "priority"
	ReplicationControllerSpecFieldPriorityClassName             = "priorityClassName"
	ReplicationControllerSpecFieldReadinessGates                = "readinessGates"
	ReplicationControllerSpecFieldReplicationControllerConfig   = "replicationControllerConfig"
	ReplicationControllerSpecFieldRestartPolicy                 = "restartPolicy"
	ReplicationControllerSpecFieldRunAsGroup                    = "runAsGroup"
	ReplicationControllerSpecFieldRunAsNonRoot                  = "runAsNonRoot"
	ReplicationControllerSpecFieldRuntimeClassName              = "runtimeClassName"
	ReplicationControllerSpecFieldScale                         = "scale"
	ReplicationControllerSpecFieldSchedulerName                 = "schedulerName"
	ReplicationControllerSpecFieldScheduling                    = "scheduling"
	ReplicationControllerSpecFieldSelector                      = "selector"
	ReplicationControllerSpecFieldServiceAccountName            = "serviceAccountName"
	ReplicationControllerSpecFieldShareProcessNamespace         = "shareProcessNamespace"
	ReplicationControllerSpecFieldSubdomain                     = "subdomain"
	ReplicationControllerSpecFieldSysctls                       = "sysctls"
	ReplicationControllerSpecFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	ReplicationControllerSpecFieldUid                           = "uid"
	ReplicationControllerSpecFieldVolumes                       = "volumes"
)

type ReplicationControllerSpec struct {
	ActiveDeadlineSeconds         *int64                       `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	AutomountServiceAccountToken  *bool                        `json:"automountServiceAccountToken,omitempty" yaml:"automount_service_account_token,omitempty"`
	Containers                    []Container                  `json:"containers,omitempty" yaml:"containers,omitempty"`
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
	NodeID                        string                       `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	ObjectMeta                    *ObjectMeta                  `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Priority                      *int64                       `json:"priority,omitempty" yaml:"priority,omitempty"`
	PriorityClassName             string                       `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	ReadinessGates                []PodReadinessGate           `json:"readinessGates,omitempty" yaml:"readiness_gates,omitempty"`
	ReplicationControllerConfig   *ReplicationControllerConfig `json:"replicationControllerConfig,omitempty" yaml:"replication_controller_config,omitempty"`
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
	Subdomain                     string                       `json:"subdomain,omitempty" yaml:"subdomain,omitempty"`
	Sysctls                       []Sysctl                     `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	TerminationGracePeriodSeconds *int64                       `json:"terminationGracePeriodSeconds,omitempty" yaml:"termination_grace_period_seconds,omitempty"`
	Uid                           *int64                       `json:"uid,omitempty" yaml:"uid,omitempty"`
	Volumes                       []Volume                     `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
