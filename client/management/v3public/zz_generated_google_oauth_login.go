package client

const (
	GoogleOauthLoginType              = "googleOauthLogin"
	GoogleOauthLoginFieldCode         = "code"
	GoogleOauthLoginFieldDescription  = "description"
	GoogleOauthLoginFieldResponseType = "responseType"
	GoogleOauthLoginFieldTTLMillis    = "ttl"
)

type GoogleOauthLogin struct {
	Code         string `json:"code,omitempty" yaml:"code,omitempty"`
	Description  string `json:"description,omitempty" yaml:"description,omitempty"`
	ResponseType string `json:"responseType,omitempty" yaml:"response_type,omitempty"`
	TTLMillis    int64  `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
