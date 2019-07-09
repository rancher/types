package client

const (
	PodSecurityPolicySpecType                                 = "podSecurityPolicySpec"
	PodSecurityPolicySpecFieldAllowPrivilegeEscalation        = "allowPrivilegeEscalation"
	PodSecurityPolicySpecFieldAllowedCapabilities             = "allowedCapabilities"
	PodSecurityPolicySpecFieldAllowedFlexVolumes              = "allowedFlexVolumes"
	PodSecurityPolicySpecFieldAllowedHostPaths                = "allowedHostPaths"
	PodSecurityPolicySpecFieldAllowedProcMountTypes           = "allowedProcMountTypes"
	PodSecurityPolicySpecFieldAllowedUnsafeSysctls            = "allowedUnsafeSysctls"
	PodSecurityPolicySpecFieldDefaultAddCapabilities          = "defaultAddCapabilities"
	PodSecurityPolicySpecFieldDefaultAllowPrivilegeEscalation = "defaultAllowPrivilegeEscalation"
	PodSecurityPolicySpecFieldFSGroup                         = "fsGroup"
	PodSecurityPolicySpecFieldForbiddenSysctls                = "forbiddenSysctls"
	PodSecurityPolicySpecFieldHostIPC                         = "hostIPC"
	PodSecurityPolicySpecFieldHostNetwork                     = "hostNetwork"
	PodSecurityPolicySpecFieldHostPID                         = "hostPID"
	PodSecurityPolicySpecFieldHostPorts                       = "hostPorts"
	PodSecurityPolicySpecFieldPrivileged                      = "privileged"
	PodSecurityPolicySpecFieldReadOnlyRootFilesystem          = "readOnlyRootFilesystem"
	PodSecurityPolicySpecFieldRequiredDropCapabilities        = "requiredDropCapabilities"
	PodSecurityPolicySpecFieldRunAsUser                       = "runAsUser"
	PodSecurityPolicySpecFieldSELinux                         = "seLinux"
	PodSecurityPolicySpecFieldSupplementalGroups              = "supplementalGroups"
	PodSecurityPolicySpecFieldVolumes                         = "volumes"
)

type PodSecurityPolicySpec struct {
	AllowPrivilegeEscalation        *bool                              `json:"allowPrivilegeEscalation,omitempty" yaml:"allow_privilege_escalation,omitempty"`
	AllowedCapabilities             []string                           `json:"allowedCapabilities,omitempty" yaml:"allowed_capabilities,omitempty"`
	AllowedFlexVolumes              []AllowedFlexVolume                `json:"allowedFlexVolumes,omitempty" yaml:"allowed_flex_volumes,omitempty"`
	AllowedHostPaths                []AllowedHostPath                  `json:"allowedHostPaths,omitempty" yaml:"allowed_host_paths,omitempty"`
	AllowedProcMountTypes           []string                           `json:"allowedProcMountTypes,omitempty" yaml:"allowed_proc_mount_types,omitempty"`
	AllowedUnsafeSysctls            []string                           `json:"allowedUnsafeSysctls,omitempty" yaml:"allowed_unsafe_sysctls,omitempty"`
	DefaultAddCapabilities          []string                           `json:"defaultAddCapabilities,omitempty" yaml:"default_add_capabilities,omitempty"`
	DefaultAllowPrivilegeEscalation *bool                              `json:"defaultAllowPrivilegeEscalation,omitempty" yaml:"default_allow_privilege_escalation,omitempty"`
	FSGroup                         *FSGroupStrategyOptions            `json:"fsGroup,omitempty" yaml:"fs_group,omitempty"`
	ForbiddenSysctls                []string                           `json:"forbiddenSysctls,omitempty" yaml:"forbidden_sysctls,omitempty"`
	HostIPC                         bool                               `json:"hostIPC,omitempty" yaml:"host_ipc,omitempty"`
	HostNetwork                     bool                               `json:"hostNetwork,omitempty" yaml:"host_network,omitempty"`
	HostPID                         bool                               `json:"hostPID,omitempty" yaml:"host_pid,omitempty"`
	HostPorts                       []HostPortRange                    `json:"hostPorts,omitempty" yaml:"host_ports,omitempty"`
	Privileged                      bool                               `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	ReadOnlyRootFilesystem          bool                               `json:"readOnlyRootFilesystem,omitempty" yaml:"read_only_root_filesystem,omitempty"`
	RequiredDropCapabilities        []string                           `json:"requiredDropCapabilities,omitempty" yaml:"required_drop_capabilities,omitempty"`
	RunAsUser                       *RunAsUserStrategyOptions          `json:"runAsUser,omitempty" yaml:"run_as_user,omitempty"`
	SELinux                         *SELinuxStrategyOptions            `json:"seLinux,omitempty" yaml:"se_linux,omitempty"`
	SupplementalGroups              *SupplementalGroupsStrategyOptions `json:"supplementalGroups,omitempty" yaml:"supplemental_groups,omitempty"`
	Volumes                         []string                           `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
