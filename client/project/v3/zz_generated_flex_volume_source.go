package client

const (
	FlexVolumeSourceType           = "flexVolumeSource"
	FlexVolumeSourceFieldDriver    = "driver"
	FlexVolumeSourceFieldFSType    = "fsType"
	FlexVolumeSourceFieldOptions   = "options"
	FlexVolumeSourceFieldReadOnly  = "readOnly"
	FlexVolumeSourceFieldSecretRef = "secretRef"
)

type FlexVolumeSource struct {
	Driver    string                `json:"driver,omitempty" yaml:"driver,omitempty"`
	FSType    string                `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	Options   map[string]string     `json:"options,omitempty" yaml:"options,omitempty"`
	ReadOnly  bool                  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
}
