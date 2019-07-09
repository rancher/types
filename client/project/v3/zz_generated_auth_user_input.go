package client

const (
	AuthUserInputType                = "authUserInput"
	AuthUserInputFieldCode           = "code"
	AuthUserInputFieldRedirectURL    = "redirectUrl"
	AuthUserInputFieldSourceCodeType = "sourceCodeType"
)

type AuthUserInput struct {
	Code           string `json:"code,omitempty" yaml:"code,omitempty"`
	RedirectURL    string `json:"redirectUrl,omitempty" yaml:"redirect_url,omitempty"`
	SourceCodeType string `json:"sourceCodeType,omitempty" yaml:"source_code_type,omitempty"`
}
