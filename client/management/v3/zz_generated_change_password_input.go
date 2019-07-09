package client

const (
	ChangePasswordInputType                 = "changePasswordInput"
	ChangePasswordInputFieldCurrentPassword = "currentPassword"
	ChangePasswordInputFieldNewPassword     = "newPassword"
)

type ChangePasswordInput struct {
	CurrentPassword string `json:"currentPassword,omitempty" yaml:"current_password,omitempty"`
	NewPassword     string `json:"newPassword,omitempty" yaml:"new_password,omitempty"`
}
