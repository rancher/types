package client

const (
	PhabricatorLoginType              = "phabricatorLogin"
	PhabricatorLoginFieldCode         = "code"
	PhabricatorLoginFieldDescription  = "description"
	PhabricatorLoginFieldResponseType = "responseType"
	PhabricatorLoginFieldTTLMillis    = "ttl"
)

type PhabricatorLogin struct {
	Code         string `json:"code,omitempty" yaml:"code,omitempty"`
	Description  string `json:"description,omitempty" yaml:"description,omitempty"`
	ResponseType string `json:"responseType,omitempty" yaml:"responseType,omitempty"`
	TTLMillis    int64  `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
