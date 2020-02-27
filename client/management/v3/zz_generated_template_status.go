package client

const (
	TemplateStatusType             = "templateStatus"
	TemplateStatusFieldHelmVersion = "helmversion"
)

type TemplateStatus struct {
	HelmVersion string `json:"helmversion,omitempty" yaml:"helmversion,omitempty"`
}
