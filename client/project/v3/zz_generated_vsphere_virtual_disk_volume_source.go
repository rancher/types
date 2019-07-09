package client

const (
	VsphereVirtualDiskVolumeSourceType                   = "vsphereVirtualDiskVolumeSource"
	VsphereVirtualDiskVolumeSourceFieldFSType            = "fsType"
	VsphereVirtualDiskVolumeSourceFieldStoragePolicyID   = "storagePolicyID"
	VsphereVirtualDiskVolumeSourceFieldStoragePolicyName = "storagePolicyName"
	VsphereVirtualDiskVolumeSourceFieldVolumePath        = "volumePath"
)

type VsphereVirtualDiskVolumeSource struct {
	FSType            string `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	StoragePolicyID   string `json:"storagePolicyID,omitempty" yaml:"storage_policy_id,omitempty"`
	StoragePolicyName string `json:"storagePolicyName,omitempty" yaml:"storage_policy_name,omitempty"`
	VolumePath        string `json:"volumePath,omitempty" yaml:"volume_path,omitempty"`
}
