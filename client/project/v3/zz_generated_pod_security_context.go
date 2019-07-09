package client

const (
	PodSecurityContextType              = "podSecurityContext"
	PodSecurityContextFieldFsgid        = "fsgid"
	PodSecurityContextFieldGids         = "gids"
	PodSecurityContextFieldRunAsGroup   = "runAsGroup"
	PodSecurityContextFieldRunAsNonRoot = "runAsNonRoot"
	PodSecurityContextFieldSysctls      = "sysctls"
	PodSecurityContextFieldUid          = "uid"
)

type PodSecurityContext struct {
	Fsgid        *int64   `json:"fsgid,omitempty" yaml:"fsgid,omitempty"`
	Gids         []int64  `json:"gids,omitempty" yaml:"gids,omitempty"`
	RunAsGroup   *int64   `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot *bool    `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	Sysctls      []Sysctl `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	Uid          *int64   `json:"uid,omitempty" yaml:"uid,omitempty"`
}
