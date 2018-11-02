package client

const (
	TemplateReferenceType         = "templateReference"
	TemplateReferenceFieldName    = "name"
	TemplateReferenceFieldVersion = "version"
)

type TemplateReference struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
