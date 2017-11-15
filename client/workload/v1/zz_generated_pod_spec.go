package client

const (
	PodSpecType                               = "podSpec"
	PodSpecFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	PodSpecFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	PodSpecFieldContainers                    = "containers"
	PodSpecFieldDNSPolicy                     = "dnsPolicy"
	PodSpecFieldDeprecatedServiceAccount      = "deprecatedServiceAccount"
	PodSpecFieldFSGroup                       = "fsGroup"
	PodSpecFieldHostAliases                   = "hostAliases"
	PodSpecFieldHostname                      = "hostname"
	PodSpecFieldIPC                           = "ipc"
	PodSpecFieldImagePullSecrets              = "imagePullSecrets"
	PodSpecFieldNet                           = "net"
	PodSpecFieldNodeName                      = "nodeName"
	PodSpecFieldPID                           = "pid"
	PodSpecFieldPriority                      = "priority"
	PodSpecFieldPriorityClassName             = "priorityClassName"
	PodSpecFieldRestartPolicy                 = "restartPolicy"
	PodSpecFieldRunAsNonRoot                  = "runAsNonRoot"
	PodSpecFieldRunAsUser                     = "runAsUser"
	PodSpecFieldSchedulerName                 = "schedulerName"
	PodSpecFieldServiceAccountName            = "serviceAccountName"
	PodSpecFieldSubdomain                     = "subdomain"
	PodSpecFieldSupplementalGroups            = "supplementalGroups"
	PodSpecFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	PodSpecFieldTolerations                   = "tolerations"
	PodSpecFieldVolumes                       = "volumes"
)

type PodSpec struct {
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty"`
	AutomountServiceAccountToken  *bool                  `json:"automountServiceAccountToken,omitempty"`
	Containers                    map[string]Container   `json:"containers,omitempty"`
	DNSPolicy                     string                 `json:"dnsPolicy,omitempty"`
	DeprecatedServiceAccount      string                 `json:"deprecatedServiceAccount,omitempty"`
	FSGroup                       *int64                 `json:"fsGroup,omitempty"`
	HostAliases                   map[string]HostAlias   `json:"hostAliases,omitempty"`
	Hostname                      string                 `json:"hostname,omitempty"`
	IPC                           string                 `json:"ipc,omitempty"`
	ImagePullSecrets              []LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Net                           string                 `json:"net,omitempty"`
	NodeName                      string                 `json:"nodeName,omitempty"`
	PID                           string                 `json:"pid,omitempty"`
	Priority                      *int64                 `json:"priority,omitempty"`
	PriorityClassName             string                 `json:"priorityClassName,omitempty"`
	RestartPolicy                 string                 `json:"restartPolicy,omitempty"`
	RunAsNonRoot                  *bool                  `json:"runAsNonRoot,omitempty"`
	RunAsUser                     *int64                 `json:"runAsUser,omitempty"`
	SchedulerName                 string                 `json:"schedulerName,omitempty"`
	ServiceAccountName            string                 `json:"serviceAccountName,omitempty"`
	Subdomain                     string                 `json:"subdomain,omitempty"`
	SupplementalGroups            []int64                `json:"supplementalGroups,omitempty"`
	TerminationGracePeriodSeconds *int64                 `json:"terminationGracePeriodSeconds,omitempty"`
	Tolerations                   []Toleration           `json:"tolerations,omitempty"`
	Volumes                       map[string]Volume      `json:"volumes,omitempty"`
}
