package client

const (
	PhabricatorConfigTestOutputType             = "phabricatorConfigTestOutput"
	PhabricatorConfigTestOutputFieldRedirectURL = "redirectUrl"
)

type PhabricatorConfigTestOutput struct {
	RedirectURL string `json:"redirectUrl,omitempty" yaml:"redirectUrl,omitempty"`
}
