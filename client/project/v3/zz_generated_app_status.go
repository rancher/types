package client

const (
	AppStatusType                      = "appStatus"
	AppStatusFieldAppliedFiles         = "appliedFiles"
	AppStatusFieldConditions           = "conditions"
	AppStatusFieldLastAppliedTemplates = "lastAppliedTemplate"
	AppStatusFieldNotes                = "notes"
)

type AppStatus struct {
	AppliedFiles         map[string]string `json:"appliedFiles,omitempty" yaml:"applied_files,omitempty"`
	Conditions           []AppCondition    `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	LastAppliedTemplates string            `json:"lastAppliedTemplate,omitempty" yaml:"last_applied_template,omitempty"`
	Notes                string            `json:"notes,omitempty" yaml:"notes,omitempty"`
}
