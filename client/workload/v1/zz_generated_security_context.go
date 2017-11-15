package client

const (
	SecurityContextType                          = "securityContext"
	SecurityContextFieldAdd                      = "add"
	SecurityContextFieldAllowPrivilegeEscalation = "allowPrivilegeEscalation"
	SecurityContextFieldDrop                     = "drop"
	SecurityContextFieldPrivileged               = "privileged"
	SecurityContextFieldReadOnlyRootFilesystem   = "readOnlyRootFilesystem"
	SecurityContextFieldRunAsNonRoot             = "runAsNonRoot"
	SecurityContextFieldRunAsUser                = "runAsUser"
)

type SecurityContext struct {
	Add                      []string `json:"add,omitempty"`
	AllowPrivilegeEscalation *bool    `json:"allowPrivilegeEscalation,omitempty"`
	Drop                     []string `json:"drop,omitempty"`
	Privileged               *bool    `json:"privileged,omitempty"`
	ReadOnlyRootFilesystem   *bool    `json:"readOnlyRootFilesystem,omitempty"`
	RunAsNonRoot             *bool    `json:"runAsNonRoot,omitempty"`
	RunAsUser                *int64   `json:"runAsUser,omitempty"`
}
