package client

const (
	ISCSIVolumeSourceType                   = "iscsiVolumeSource"
	ISCSIVolumeSourceFieldDiscoveryCHAPAuth = "chapAuthDiscovery"
	ISCSIVolumeSourceFieldFSType            = "fsType"
	ISCSIVolumeSourceFieldIQN               = "iqn"
	ISCSIVolumeSourceFieldISCSIInterface    = "iscsiInterface"
	ISCSIVolumeSourceFieldInitiatorName     = "initiatorName"
	ISCSIVolumeSourceFieldLun               = "lun"
	ISCSIVolumeSourceFieldPortals           = "portals"
	ISCSIVolumeSourceFieldReadOnly          = "readOnly"
	ISCSIVolumeSourceFieldSecretRef         = "secretRef"
	ISCSIVolumeSourceFieldSessionCHAPAuth   = "chapAuthSession"
	ISCSIVolumeSourceFieldTargetPortal      = "targetPortal"
)

type ISCSIVolumeSource struct {
	DiscoveryCHAPAuth bool                  `json:"chapAuthDiscovery,omitempty" yaml:"chap_auth_discovery,omitempty"`
	FSType            string                `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	IQN               string                `json:"iqn,omitempty" yaml:"iqn,omitempty"`
	ISCSIInterface    string                `json:"iscsiInterface,omitempty" yaml:"iscsi_interface,omitempty"`
	InitiatorName     string                `json:"initiatorName,omitempty" yaml:"initiator_name,omitempty"`
	Lun               int64                 `json:"lun,omitempty" yaml:"lun,omitempty"`
	Portals           []string              `json:"portals,omitempty" yaml:"portals,omitempty"`
	ReadOnly          bool                  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef         *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	SessionCHAPAuth   bool                  `json:"chapAuthSession,omitempty" yaml:"chap_auth_session,omitempty"`
	TargetPortal      string                `json:"targetPortal,omitempty" yaml:"target_portal,omitempty"`
}
