package client

const (
	GlusterfsVolumeSourceType               = "glusterfsVolumeSource"
	GlusterfsVolumeSourceFieldEndpointsName = "endpointsName"
	GlusterfsVolumeSourceFieldPath          = "path"
	GlusterfsVolumeSourceFieldReadOnly      = "readOnly"
)

type GlusterfsVolumeSource struct {
	EndpointsName string `json:"endpointsName,omitempty"`
	Path          string `json:"path,omitempty"`
	ReadOnly      bool   `json:"readOnly,omitempty"`
}
