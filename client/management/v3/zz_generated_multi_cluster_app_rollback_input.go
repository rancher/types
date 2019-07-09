package client

const (
	MultiClusterAppRollbackInputType            = "multiClusterAppRollbackInput"
	MultiClusterAppRollbackInputFieldRevisionID = "revisionId"
)

type MultiClusterAppRollbackInput struct {
	RevisionID string `json:"revisionId,omitempty" yaml:"revision_id,omitempty"`
}
