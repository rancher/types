package client

const (
	ISCSIPersistentVolumeSourceType                   = "iscsiPersistentVolumeSource"
	ISCSIPersistentVolumeSourceFieldDiscoveryCHAPAuth = "chapAuthDiscovery"
	ISCSIPersistentVolumeSourceFieldFSType            = "fsType"
	ISCSIPersistentVolumeSourceFieldIQN               = "iqn"
	ISCSIPersistentVolumeSourceFieldISCSIInterface    = "iscsiInterface"
	ISCSIPersistentVolumeSourceFieldInitiatorName     = "initiatorName"
	ISCSIPersistentVolumeSourceFieldLun               = "lun"
	ISCSIPersistentVolumeSourceFieldPortals           = "portals"
	ISCSIPersistentVolumeSourceFieldReadOnly          = "readOnly"
	ISCSIPersistentVolumeSourceFieldSecretRef         = "secretRef"
	ISCSIPersistentVolumeSourceFieldSessionCHAPAuth   = "chapAuthSession"
	ISCSIPersistentVolumeSourceFieldTargetPortal      = "targetPortal"
)

type ISCSIPersistentVolumeSource struct {
	DiscoveryCHAPAuth bool             `json:"chapAuthDiscovery,omitempty" yaml:"chap_auth_discovery,omitempty"`
	FSType            string           `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	IQN               string           `json:"iqn,omitempty" yaml:"iqn,omitempty"`
	ISCSIInterface    string           `json:"iscsiInterface,omitempty" yaml:"iscsi_interface,omitempty"`
	InitiatorName     string           `json:"initiatorName,omitempty" yaml:"initiator_name,omitempty"`
	Lun               int64            `json:"lun,omitempty" yaml:"lun,omitempty"`
	Portals           []string         `json:"portals,omitempty" yaml:"portals,omitempty"`
	ReadOnly          bool             `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef         *SecretReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	SessionCHAPAuth   bool             `json:"chapAuthSession,omitempty" yaml:"chap_auth_session,omitempty"`
	TargetPortal      string           `json:"targetPortal,omitempty" yaml:"target_portal,omitempty"`
}
