package client

const (
	CSIPersistentVolumeSourceType                            = "csiPersistentVolumeSource"
	CSIPersistentVolumeSourceFieldControllerPublishSecretRef = "controllerPublishSecretRef"
	CSIPersistentVolumeSourceFieldDriver                     = "driver"
	CSIPersistentVolumeSourceFieldFSType                     = "fsType"
	CSIPersistentVolumeSourceFieldNodePublishSecretRef       = "nodePublishSecretRef"
	CSIPersistentVolumeSourceFieldNodeStageSecretRef         = "nodeStageSecretRef"
	CSIPersistentVolumeSourceFieldReadOnly                   = "readOnly"
	CSIPersistentVolumeSourceFieldVolumeAttributes           = "volumeAttributes"
	CSIPersistentVolumeSourceFieldVolumeHandle               = "volumeHandle"
)

type CSIPersistentVolumeSource struct {
	ControllerPublishSecretRef *SecretReference  `json:"controllerPublishSecretRef,omitempty" yaml:"controller_publish_secret_ref,omitempty"`
	Driver                     string            `json:"driver,omitempty" yaml:"driver,omitempty"`
	FSType                     string            `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	NodePublishSecretRef       *SecretReference  `json:"nodePublishSecretRef,omitempty" yaml:"node_publish_secret_ref,omitempty"`
	NodeStageSecretRef         *SecretReference  `json:"nodeStageSecretRef,omitempty" yaml:"node_stage_secret_ref,omitempty"`
	ReadOnly                   bool              `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	VolumeAttributes           map[string]string `json:"volumeAttributes,omitempty" yaml:"volume_attributes,omitempty"`
	VolumeHandle               string            `json:"volumeHandle,omitempty" yaml:"volume_handle,omitempty"`
}
