package client

const (
	AzureDiskVolumeSourceType             = "azureDiskVolumeSource"
	AzureDiskVolumeSourceFieldCachingMode = "cachingMode"
	AzureDiskVolumeSourceFieldDataDiskURI = "diskURI"
	AzureDiskVolumeSourceFieldDiskName    = "diskName"
	AzureDiskVolumeSourceFieldFSType      = "fsType"
	AzureDiskVolumeSourceFieldKind        = "kind"
	AzureDiskVolumeSourceFieldReadOnly    = "readOnly"
)

type AzureDiskVolumeSource struct {
	CachingMode string `json:"cachingMode,omitempty" yaml:"caching_mode,omitempty"`
	DataDiskURI string `json:"diskURI,omitempty" yaml:"disk_uri,omitempty"`
	DiskName    string `json:"diskName,omitempty" yaml:"disk_name,omitempty"`
	FSType      string `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	Kind        string `json:"kind,omitempty" yaml:"kind,omitempty"`
	ReadOnly    *bool  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
}
