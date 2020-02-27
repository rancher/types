package client

const (
	AppStatusType                      = "appStatus"
	AppStatusFieldAppliedFiles         = "appliedFiles"
	AppStatusFieldConditions           = "conditions"
	AppStatusFieldHelmVersion          = "helmversion"
	AppStatusFieldLastAppliedTemplates = "lastAppliedTemplate"
	AppStatusFieldNotes                = "notes"
)

type AppStatus struct {
	AppliedFiles         map[string]string `json:"appliedFiles,omitempty" yaml:"appliedFiles,omitempty"`
	Conditions           []AppCondition    `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	HelmVersion          string            `json:"helmversion,omitempty" yaml:"helmversion,omitempty"`
	LastAppliedTemplates string            `json:"lastAppliedTemplate,omitempty" yaml:"lastAppliedTemplate,omitempty"`
	Notes                string            `json:"notes,omitempty" yaml:"notes,omitempty"`
}
