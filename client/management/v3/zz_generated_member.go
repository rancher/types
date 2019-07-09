package client

const (
	MemberType                  = "member"
	MemberFieldAccessType       = "accessType"
	MemberFieldGroupPrincipalID = "groupPrincipalId"
	MemberFieldUserPrincipalID  = "userPrincipalId"
)

type Member struct {
	AccessType       string `json:"accessType,omitempty" yaml:"access_type,omitempty"`
	GroupPrincipalID string `json:"groupPrincipalId,omitempty" yaml:"group_principal_id,omitempty"`
	UserPrincipalID  string `json:"userPrincipalId,omitempty" yaml:"user_principal_id,omitempty"`
}
