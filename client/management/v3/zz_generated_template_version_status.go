package client

const (
	TemplateVersionStatusType             = "templateVersionStatus"
	TemplateVersionStatusFieldHelmVersion = "helmversion"
)

type TemplateVersionStatus struct {
	HelmVersion string `json:"helmversion,omitempty" yaml:"helmversion,omitempty"`
}
