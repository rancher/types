package client

const (
	PersistentVolumeClaimVolumeSourceType                         = "persistentVolumeClaimVolumeSource"
	PersistentVolumeClaimVolumeSourceFieldPersistentVolumeClaimID = "persistentVolumeClaimId"
	PersistentVolumeClaimVolumeSourceFieldReadOnly                = "readOnly"
)

type PersistentVolumeClaimVolumeSource struct {
	PersistentVolumeClaimID string `json:"persistentVolumeClaimId,omitempty" yaml:"persistent_volume_claim_id,omitempty"`
	ReadOnly                bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
}
