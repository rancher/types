package client

const (
	PersistentVolumeClaimSpecType                = "persistentVolumeClaimSpec"
	PersistentVolumeClaimSpecFieldAccessModes    = "accessModes"
	PersistentVolumeClaimSpecFieldDataSource     = "dataSource"
	PersistentVolumeClaimSpecFieldResources      = "resources"
	PersistentVolumeClaimSpecFieldSelector       = "selector"
	PersistentVolumeClaimSpecFieldStorageClassID = "storageClassId"
	PersistentVolumeClaimSpecFieldVolumeID       = "volumeId"
	PersistentVolumeClaimSpecFieldVolumeMode     = "volumeMode"
)

type PersistentVolumeClaimSpec struct {
	AccessModes    []string                   `json:"accessModes,omitempty" yaml:"access_modes,omitempty"`
	DataSource     *TypedLocalObjectReference `json:"dataSource,omitempty" yaml:"data_source,omitempty"`
	Resources      *ResourceRequirements      `json:"resources,omitempty" yaml:"resources,omitempty"`
	Selector       *LabelSelector             `json:"selector,omitempty" yaml:"selector,omitempty"`
	StorageClassID string                     `json:"storageClassId,omitempty" yaml:"storage_class_id,omitempty"`
	VolumeID       string                     `json:"volumeId,omitempty" yaml:"volume_id,omitempty"`
	VolumeMode     string                     `json:"volumeMode,omitempty" yaml:"volume_mode,omitempty"`
}
