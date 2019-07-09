package client

const (
	JobTemplateSpecType                               = "jobTemplateSpec"
	JobTemplateSpecFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	JobTemplateSpecFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	JobTemplateSpecFieldContainers                    = "containers"
	JobTemplateSpecFieldDNSConfig                     = "dnsConfig"
	JobTemplateSpecFieldDNSPolicy                     = "dnsPolicy"
	JobTemplateSpecFieldFsgid                         = "fsgid"
	JobTemplateSpecFieldGids                          = "gids"
	JobTemplateSpecFieldHostAliases                   = "hostAliases"
	JobTemplateSpecFieldHostIPC                       = "hostIPC"
	JobTemplateSpecFieldHostNetwork                   = "hostNetwork"
	JobTemplateSpecFieldHostPID                       = "hostPID"
	JobTemplateSpecFieldHostname                      = "hostname"
	JobTemplateSpecFieldImagePullSecrets              = "imagePullSecrets"
	JobTemplateSpecFieldJobConfig                     = "jobConfig"
	JobTemplateSpecFieldJobMetadata                   = "jobMetadata"
	JobTemplateSpecFieldNodeID                        = "nodeId"
	JobTemplateSpecFieldObjectMeta                    = "metadata"
	JobTemplateSpecFieldPriority                      = "priority"
	JobTemplateSpecFieldPriorityClassName             = "priorityClassName"
	JobTemplateSpecFieldReadinessGates                = "readinessGates"
	JobTemplateSpecFieldRestartPolicy                 = "restartPolicy"
	JobTemplateSpecFieldRunAsGroup                    = "runAsGroup"
	JobTemplateSpecFieldRunAsNonRoot                  = "runAsNonRoot"
	JobTemplateSpecFieldRuntimeClassName              = "runtimeClassName"
	JobTemplateSpecFieldSchedulerName                 = "schedulerName"
	JobTemplateSpecFieldScheduling                    = "scheduling"
	JobTemplateSpecFieldSelector                      = "selector"
	JobTemplateSpecFieldServiceAccountName            = "serviceAccountName"
	JobTemplateSpecFieldShareProcessNamespace         = "shareProcessNamespace"
	JobTemplateSpecFieldSubdomain                     = "subdomain"
	JobTemplateSpecFieldSysctls                       = "sysctls"
	JobTemplateSpecFieldTTLSecondsAfterFinished       = "ttlSecondsAfterFinished"
	JobTemplateSpecFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	JobTemplateSpecFieldUid                           = "uid"
	JobTemplateSpecFieldVolumes                       = "volumes"
)

type JobTemplateSpec struct {
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty" yaml:"active_deadline_seconds,omitempty"`
	AutomountServiceAccountToken  *bool                  `json:"automountServiceAccountToken,omitempty" yaml:"automount_service_account_token,omitempty"`
	Containers                    []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
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
	JobMetadata                   *ObjectMeta            `json:"jobMetadata,omitempty" yaml:"job_metadata,omitempty"`
	NodeID                        string                 `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	ObjectMeta                    *ObjectMeta            `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Priority                      *int64                 `json:"priority,omitempty" yaml:"priority,omitempty"`
	PriorityClassName             string                 `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	ReadinessGates                []PodReadinessGate     `json:"readinessGates,omitempty" yaml:"readiness_gates,omitempty"`
	RestartPolicy                 string                 `json:"restartPolicy,omitempty" yaml:"restart_policy,omitempty"`
	RunAsGroup                    *int64                 `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot                  *bool                  `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	RuntimeClassName              string                 `json:"runtimeClassName,omitempty" yaml:"runtime_class_name,omitempty"`
	SchedulerName                 string                 `json:"schedulerName,omitempty" yaml:"scheduler_name,omitempty"`
	Scheduling                    *Scheduling            `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	Selector                      *LabelSelector         `json:"selector,omitempty" yaml:"selector,omitempty"`
	ServiceAccountName            string                 `json:"serviceAccountName,omitempty" yaml:"service_account_name,omitempty"`
	ShareProcessNamespace         *bool                  `json:"shareProcessNamespace,omitempty" yaml:"share_process_namespace,omitempty"`
	Subdomain                     string                 `json:"subdomain,omitempty" yaml:"subdomain,omitempty"`
	Sysctls                       []Sysctl               `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	TTLSecondsAfterFinished       *int64                 `json:"ttlSecondsAfterFinished,omitempty" yaml:"ttl_seconds_after_finished,omitempty"`
	TerminationGracePeriodSeconds *int64                 `json:"terminationGracePeriodSeconds,omitempty" yaml:"termination_grace_period_seconds,omitempty"`
	Uid                           *int64                 `json:"uid,omitempty" yaml:"uid,omitempty"`
	Volumes                       []Volume               `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
