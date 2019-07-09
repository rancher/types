package client

const (
	BlockStorageOpenstackOptsType                 = "blockStorageOpenstackOpts"
	BlockStorageOpenstackOptsFieldBSVersion       = "bs-version"
	BlockStorageOpenstackOptsFieldIgnoreVolumeAZ  = "ignore-volume-az"
	BlockStorageOpenstackOptsFieldTrustDevicePath = "trust-device-path"
)

type BlockStorageOpenstackOpts struct {
	BSVersion       string `json:"bs-version,omitempty" yaml:"bs_version,omitempty"`
	IgnoreVolumeAZ  bool   `json:"ignore-volume-az,omitempty" yaml:"ignore_volume_az,omitempty"`
	TrustDevicePath bool   `json:"trust-device-path,omitempty" yaml:"trust_device_path,omitempty"`
}
