package client

const (
	VolumeDeviceType            = "volumeDevice"
	VolumeDeviceFieldDevicePath = "devicePath"
	VolumeDeviceFieldName       = "name"
)

type VolumeDevice struct {
	DevicePath string `json:"devicePath,omitempty" yaml:"device_path,omitempty"`
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
}
