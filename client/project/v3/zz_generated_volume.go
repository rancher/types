package client

const (
	VolumeType                       = "volume"
	VolumeFieldAWSElasticBlockStore  = "awsElasticBlockStore"
	VolumeFieldAzureDisk             = "azureDisk"
	VolumeFieldAzureFile             = "azureFile"
	VolumeFieldCephFS                = "cephfs"
	VolumeFieldCinder                = "cinder"
	VolumeFieldConfigMap             = "configMap"
	VolumeFieldDownwardAPI           = "downwardAPI"
	VolumeFieldEmptyDir              = "emptyDir"
	VolumeFieldFC                    = "fc"
	VolumeFieldFlexVolume            = "flexVolume"
	VolumeFieldFlocker               = "flocker"
	VolumeFieldGCEPersistentDisk     = "gcePersistentDisk"
	VolumeFieldGitRepo               = "gitRepo"
	VolumeFieldGlusterfs             = "glusterfs"
	VolumeFieldHostPath              = "hostPath"
	VolumeFieldISCSI                 = "iscsi"
	VolumeFieldNFS                   = "nfs"
	VolumeFieldName                  = "name"
	VolumeFieldPersistentVolumeClaim = "persistentVolumeClaim"
	VolumeFieldPhotonPersistentDisk  = "photonPersistentDisk"
	VolumeFieldPortworxVolume        = "portworxVolume"
	VolumeFieldProjected             = "projected"
	VolumeFieldQuobyte               = "quobyte"
	VolumeFieldRBD                   = "rbd"
	VolumeFieldScaleIO               = "scaleIO"
	VolumeFieldSecret                = "secret"
	VolumeFieldStorageOS             = "storageos"
	VolumeFieldVsphereVolume         = "vsphereVolume"
)

type Volume struct {
	AWSElasticBlockStore  *AWSElasticBlockStoreVolumeSource  `json:"awsElasticBlockStore,omitempty" yaml:"aws_elastic_block_store,omitempty"`
	AzureDisk             *AzureDiskVolumeSource             `json:"azureDisk,omitempty" yaml:"azure_disk,omitempty"`
	AzureFile             *AzureFileVolumeSource             `json:"azureFile,omitempty" yaml:"azure_file,omitempty"`
	CephFS                *CephFSVolumeSource                `json:"cephfs,omitempty" yaml:"cephfs,omitempty"`
	Cinder                *CinderVolumeSource                `json:"cinder,omitempty" yaml:"cinder,omitempty"`
	ConfigMap             *ConfigMapVolumeSource             `json:"configMap,omitempty" yaml:"config_map,omitempty"`
	DownwardAPI           *DownwardAPIVolumeSource           `json:"downwardAPI,omitempty" yaml:"downward_api,omitempty"`
	EmptyDir              *EmptyDirVolumeSource              `json:"emptyDir,omitempty" yaml:"empty_dir,omitempty"`
	FC                    *FCVolumeSource                    `json:"fc,omitempty" yaml:"fc,omitempty"`
	FlexVolume            *FlexVolumeSource                  `json:"flexVolume,omitempty" yaml:"flex_volume,omitempty"`
	Flocker               *FlockerVolumeSource               `json:"flocker,omitempty" yaml:"flocker,omitempty"`
	GCEPersistentDisk     *GCEPersistentDiskVolumeSource     `json:"gcePersistentDisk,omitempty" yaml:"gce_persistent_disk,omitempty"`
	GitRepo               *GitRepoVolumeSource               `json:"gitRepo,omitempty" yaml:"git_repo,omitempty"`
	Glusterfs             *GlusterfsVolumeSource             `json:"glusterfs,omitempty" yaml:"glusterfs,omitempty"`
	HostPath              *HostPathVolumeSource              `json:"hostPath,omitempty" yaml:"host_path,omitempty"`
	ISCSI                 *ISCSIVolumeSource                 `json:"iscsi,omitempty" yaml:"iscsi,omitempty"`
	NFS                   *NFSVolumeSource                   `json:"nfs,omitempty" yaml:"nfs,omitempty"`
	Name                  string                             `json:"name,omitempty" yaml:"name,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty" yaml:"persistent_volume_claim,omitempty"`
	PhotonPersistentDisk  *PhotonPersistentDiskVolumeSource  `json:"photonPersistentDisk,omitempty" yaml:"photon_persistent_disk,omitempty"`
	PortworxVolume        *PortworxVolumeSource              `json:"portworxVolume,omitempty" yaml:"portworx_volume,omitempty"`
	Projected             *ProjectedVolumeSource             `json:"projected,omitempty" yaml:"projected,omitempty"`
	Quobyte               *QuobyteVolumeSource               `json:"quobyte,omitempty" yaml:"quobyte,omitempty"`
	RBD                   *RBDVolumeSource                   `json:"rbd,omitempty" yaml:"rbd,omitempty"`
	ScaleIO               *ScaleIOVolumeSource               `json:"scaleIO,omitempty" yaml:"scale_io,omitempty"`
	Secret                *SecretVolumeSource                `json:"secret,omitempty" yaml:"secret,omitempty"`
	StorageOS             *StorageOSVolumeSource             `json:"storageos,omitempty" yaml:"storageos,omitempty"`
	VsphereVolume         *VsphereVirtualDiskVolumeSource    `json:"vsphereVolume,omitempty" yaml:"vsphere_volume,omitempty"`
}
