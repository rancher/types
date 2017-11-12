package client

const (
	AzureDiskVolumeSourceType             = "azureDiskVolumeSource"
	AzureDiskVolumeSourceFieldCachingMode = "cachingMode"
	AzureDiskVolumeSourceFieldDataDiskURI = "dataDiskURI"
	AzureDiskVolumeSourceFieldDiskName    = "diskName"
	AzureDiskVolumeSourceFieldFSType      = "fsType"
	AzureDiskVolumeSourceFieldKind        = "kind"
	AzureDiskVolumeSourceFieldReadOnly    = "readOnly"
)

type AzureDiskVolumeSource struct {
	CachingMode string `json:"cachingMode,omitempty"`
	DataDiskURI string `json:"dataDiskURI,omitempty"`
	DiskName    string `json:"diskName,omitempty"`
	FSType      string `json:"fsType,omitempty"`
	Kind        string `json:"kind,omitempty"`
	ReadOnly    *bool  `json:"readOnly,omitempty"`
}
