package client

const (
	RollbackRevisionType              = "rollbackRevision"
	RollbackRevisionFieldForceUpgrade = "forceUpgrade"
	RollbackRevisionFieldRevisionID   = "revisionId"
)

type RollbackRevision struct {
	ForceUpgrade bool   `json:"forceUpgrade,omitempty" yaml:"force_upgrade,omitempty"`
	RevisionID   string `json:"revisionId,omitempty" yaml:"revision_id,omitempty"`
}
