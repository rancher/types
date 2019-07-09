package client

const (
	FCVolumeSourceType            = "fcVolumeSource"
	FCVolumeSourceFieldFSType     = "fsType"
	FCVolumeSourceFieldLun        = "lun"
	FCVolumeSourceFieldReadOnly   = "readOnly"
	FCVolumeSourceFieldTargetWWNs = "targetWWNs"
	FCVolumeSourceFieldWWIDs      = "wwids"
)

type FCVolumeSource struct {
	FSType     string   `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	Lun        *int64   `json:"lun,omitempty" yaml:"lun,omitempty"`
	ReadOnly   bool     `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	TargetWWNs []string `json:"targetWWNs,omitempty" yaml:"target_ww_ns,omitempty"`
	WWIDs      []string `json:"wwids,omitempty" yaml:"wwids,omitempty"`
}
