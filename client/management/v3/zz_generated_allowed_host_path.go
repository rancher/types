package client

const (
	AllowedHostPathType            = "allowedHostPath"
	AllowedHostPathFieldPathPrefix = "pathPrefix"
	AllowedHostPathFieldReadOnly   = "readOnly"
)

type AllowedHostPath struct {
	PathPrefix string `json:"pathPrefix,omitempty" yaml:"path_prefix,omitempty"`
	ReadOnly   bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
}
