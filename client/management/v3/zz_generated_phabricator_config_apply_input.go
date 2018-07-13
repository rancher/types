package client

const (
	PhabricatorConfigApplyInputType                   = "phabricatorConfigApplyInput"
	PhabricatorConfigApplyInputFieldCode              = "code"
	PhabricatorConfigApplyInputFieldEnabled           = "enabled"
	PhabricatorConfigApplyInputFieldPhabricatorConfig = "phabricatorConfig"
)

type PhabricatorConfigApplyInput struct {
	Code              string             `json:"code,omitempty" yaml:"code,omitempty"`
	Enabled           bool               `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	PhabricatorConfig *PhabricatorConfig `json:"phabricatorConfig,omitempty" yaml:"phabricatorConfig,omitempty"`
}
