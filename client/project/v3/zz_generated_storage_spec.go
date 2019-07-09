package client

const (
	StorageSpecType                     = "storageSpec"
	StorageSpecFieldEmptyDir            = "emptyDir"
	StorageSpecFieldVolumeClaimTemplate = "volumeClaimTemplate"
)

type StorageSpec struct {
	EmptyDir            *EmptyDirVolumeSource  `json:"emptyDir,omitempty" yaml:"empty_dir,omitempty"`
	VolumeClaimTemplate *PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty" yaml:"volume_claim_template,omitempty"`
}
