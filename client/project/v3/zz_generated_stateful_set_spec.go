package client

const (
	StatefulSetSpecType                               = "statefulSetSpec"
	StatefulSetSpecFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	StatefulSetSpecFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	StatefulSetSpecFieldContainers                    = "containers"
	StatefulSetSpecFieldDNSConfig                     = "dnsConfig"
	StatefulSetSpecFieldDNSPolicy                     = "dnsPolicy"
	StatefulSetSpecFieldFsgid                         = "fsgid"
	StatefulSetSpecFieldGids                          = "gids"
	StatefulSetSpecFieldHostAliases                   = "hostAliases"
	StatefulSetSpecFieldHostIPC                       = "hostIPC"
	StatefulSetSpecFieldHostNetwork                   = "hostNetwork"
	StatefulSetSpecFieldHostPID                       = "hostPID"
	StatefulSetSpecFieldHostname                      = "hostname"
	StatefulSetSpecFieldImagePullSecrets              = "imagePullSecrets"
	StatefulSetSpecFieldNodeID                        = "nodeId"
	StatefulSetSpecFieldObjectMeta                    = "metadata"
	StatefulSetSpecFieldPriority                      = "priority"
	StatefulSetSpecFieldPriorityClassName             = "priorityClassName"
	StatefulSetSpecFieldReadinessGates                = "readinessGates"
	StatefulSetSpecFieldRestartPolicy                 = "restartPolicy"
	StatefulSetSpecFieldRunAsGroup                    = "runAsGroup"
	StatefulSetSpecFieldRunAsNonRoot                  = "runAsNonRoot"
	StatefulSetSpecFieldRuntimeClassName              = "runtimeClassName"
	StatefulSetSpecFieldScale                         = "scale"
	StatefulSetSpecFieldSchedulerName                 = "schedulerName"
	StatefulSetSpecFieldScheduling                    = "scheduling"
	StatefulSetSpecFieldSelector                      = "selector"
	StatefulSetSpecFieldServiceAccountName            = "serviceAccountName"
	StatefulSetSpecFieldShareProcessNamespace         = "shareProcessNamespace"
	StatefulSetSpecFieldStatefulSetConfig             = "statefulSetConfig"
	StatefulSetSpecFieldSubdomain                     = "subdomain"
	StatefulSetSpecFieldSysctls                       = "sysctls"
	StatefulSetSpecFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	StatefulSetSpecFieldUid                           = "uid"
	StatefulSetSpecFieldVolumes                       = "volumes"
)

type StatefulSetSpec struct {
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
	NodeID                        string                 `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	ObjectMeta                    *ObjectMeta            `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Priority                      *int64                 `json:"priority,omitempty" yaml:"priority,omitempty"`
	PriorityClassName             string                 `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	ReadinessGates                []PodReadinessGate     `json:"readinessGates,omitempty" yaml:"readiness_gates,omitempty"`
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
	StatefulSetConfig             *StatefulSetConfig     `json:"statefulSetConfig,omitempty" yaml:"stateful_set_config,omitempty"`
	Subdomain                     string                 `json:"subdomain,omitempty" yaml:"subdomain,omitempty"`
	Sysctls                       []Sysctl               `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	TerminationGracePeriodSeconds *int64                 `json:"terminationGracePeriodSeconds,omitempty" yaml:"termination_grace_period_seconds,omitempty"`
	Uid                           *int64                 `json:"uid,omitempty" yaml:"uid,omitempty"`
	Volumes                       []Volume               `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
