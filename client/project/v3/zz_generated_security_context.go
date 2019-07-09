package client

const (
	SecurityContextType                          = "securityContext"
	SecurityContextFieldAllowPrivilegeEscalation = "allowPrivilegeEscalation"
	SecurityContextFieldCapAdd                   = "capAdd"
	SecurityContextFieldCapDrop                  = "capDrop"
	SecurityContextFieldPrivileged               = "privileged"
	SecurityContextFieldProcMount                = "procMount"
	SecurityContextFieldReadOnly                 = "readOnly"
	SecurityContextFieldRunAsGroup               = "runAsGroup"
	SecurityContextFieldRunAsNonRoot             = "runAsNonRoot"
	SecurityContextFieldUid                      = "uid"
)

type SecurityContext struct {
	AllowPrivilegeEscalation *bool    `json:"allowPrivilegeEscalation,omitempty" yaml:"allow_privilege_escalation,omitempty"`
	CapAdd                   []string `json:"capAdd,omitempty" yaml:"cap_add,omitempty"`
	CapDrop                  []string `json:"capDrop,omitempty" yaml:"cap_drop,omitempty"`
	Privileged               *bool    `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	ProcMount                string   `json:"procMount,omitempty" yaml:"proc_mount,omitempty"`
	ReadOnly                 *bool    `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	RunAsGroup               *int64   `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot             *bool    `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	Uid                      *int64   `json:"uid,omitempty" yaml:"uid,omitempty"`
}
