package client

const (
	PhotonPersistentDiskVolumeSourceType        = "photonPersistentDiskVolumeSource"
	PhotonPersistentDiskVolumeSourceFieldFSType = "fsType"
	PhotonPersistentDiskVolumeSourceFieldPdID   = "pdID"
)

type PhotonPersistentDiskVolumeSource struct {
	FSType string `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	PdID   string `json:"pdID,omitempty" yaml:"pd_id,omitempty"`
}
