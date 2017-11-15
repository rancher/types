package client

const (
	PodSecurityContextType                    = "podSecurityContext"
	PodSecurityContextFieldFSGroup            = "fsGroup"
	PodSecurityContextFieldRunAsNonRoot       = "runAsNonRoot"
	PodSecurityContextFieldRunAsUser          = "runAsUser"
	PodSecurityContextFieldSupplementalGroups = "supplementalGroups"
)

type PodSecurityContext struct {
	FSGroup            *int64  `json:"fsGroup,omitempty"`
	RunAsNonRoot       *bool   `json:"runAsNonRoot,omitempty"`
	RunAsUser          *int64  `json:"runAsUser,omitempty"`
	SupplementalGroups []int64 `json:"supplementalGroups,omitempty"`
}
