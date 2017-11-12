package client

const (
	ISCSIVolumeSourceType                   = "iscsiVolumeSource"
	ISCSIVolumeSourceFieldDiscoveryCHAPAuth = "discoveryCHAPAuth"
	ISCSIVolumeSourceFieldFSType            = "fsType"
	ISCSIVolumeSourceFieldIQN               = "iqn"
	ISCSIVolumeSourceFieldISCSIInterface    = "iscsiInterface"
	ISCSIVolumeSourceFieldInitiatorName     = "initiatorName"
	ISCSIVolumeSourceFieldLun               = "lun"
	ISCSIVolumeSourceFieldPortals           = "portals"
	ISCSIVolumeSourceFieldReadOnly          = "readOnly"
	ISCSIVolumeSourceFieldSecretRef         = "secretRef"
	ISCSIVolumeSourceFieldSessionCHAPAuth   = "sessionCHAPAuth"
	ISCSIVolumeSourceFieldTargetPortal      = "targetPortal"
)

type ISCSIVolumeSource struct {
	DiscoveryCHAPAuth bool                  `json:"discoveryCHAPAuth,omitempty"`
	FSType            string                `json:"fsType,omitempty"`
	IQN               string                `json:"iqn,omitempty"`
	ISCSIInterface    string                `json:"iscsiInterface,omitempty"`
	InitiatorName     string                `json:"initiatorName,omitempty"`
	Lun               int64                 `json:"lun,omitempty"`
	Portals           []string              `json:"portals,omitempty"`
	ReadOnly          bool                  `json:"readOnly,omitempty"`
	SecretRef         *LocalObjectReference `json:"secretRef,omitempty"`
	SessionCHAPAuth   bool                  `json:"sessionCHAPAuth,omitempty"`
	TargetPortal      string                `json:"targetPortal,omitempty"`
}
