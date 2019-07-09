package client

const (
	PersistentVolumeSpecType                               = "persistentVolumeSpec"
	PersistentVolumeSpecFieldAWSElasticBlockStore          = "awsElasticBlockStore"
	PersistentVolumeSpecFieldAccessModes                   = "accessModes"
	PersistentVolumeSpecFieldAzureDisk                     = "azureDisk"
	PersistentVolumeSpecFieldAzureFile                     = "azureFile"
	PersistentVolumeSpecFieldCSI                           = "csi"
	PersistentVolumeSpecFieldCapacity                      = "capacity"
	PersistentVolumeSpecFieldCephFS                        = "cephfs"
	PersistentVolumeSpecFieldCinder                        = "cinder"
	PersistentVolumeSpecFieldClaimRef                      = "claimRef"
	PersistentVolumeSpecFieldFC                            = "fc"
	PersistentVolumeSpecFieldFlexVolume                    = "flexVolume"
	PersistentVolumeSpecFieldFlocker                       = "flocker"
	PersistentVolumeSpecFieldGCEPersistentDisk             = "gcePersistentDisk"
	PersistentVolumeSpecFieldGlusterfs                     = "glusterfs"
	PersistentVolumeSpecFieldHostPath                      = "hostPath"
	PersistentVolumeSpecFieldISCSI                         = "iscsi"
	PersistentVolumeSpecFieldLocal                         = "local"
	PersistentVolumeSpecFieldMountOptions                  = "mountOptions"
	PersistentVolumeSpecFieldNFS                           = "nfs"
	PersistentVolumeSpecFieldNodeAffinity                  = "nodeAffinity"
	PersistentVolumeSpecFieldPersistentVolumeReclaimPolicy = "persistentVolumeReclaimPolicy"
	PersistentVolumeSpecFieldPhotonPersistentDisk          = "photonPersistentDisk"
	PersistentVolumeSpecFieldPortworxVolume                = "portworxVolume"
	PersistentVolumeSpecFieldQuobyte                       = "quobyte"
	PersistentVolumeSpecFieldRBD                           = "rbd"
	PersistentVolumeSpecFieldScaleIO                       = "scaleIO"
	PersistentVolumeSpecFieldStorageClassID                = "storageClassId"
	PersistentVolumeSpecFieldStorageOS                     = "storageos"
	PersistentVolumeSpecFieldVolumeMode                    = "volumeMode"
	PersistentVolumeSpecFieldVsphereVolume                 = "vsphereVolume"
)

type PersistentVolumeSpec struct {
	AWSElasticBlockStore          *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" yaml:"aws_elastic_block_store,omitempty"`
	AccessModes                   []string                          `json:"accessModes,omitempty" yaml:"access_modes,omitempty"`
	AzureDisk                     *AzureDiskVolumeSource            `json:"azureDisk,omitempty" yaml:"azure_disk,omitempty"`
	AzureFile                     *AzureFilePersistentVolumeSource  `json:"azureFile,omitempty" yaml:"azure_file,omitempty"`
	CSI                           *CSIPersistentVolumeSource        `json:"csi,omitempty" yaml:"csi,omitempty"`
	Capacity                      map[string]string                 `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	CephFS                        *CephFSPersistentVolumeSource     `json:"cephfs,omitempty" yaml:"cephfs,omitempty"`
	Cinder                        *CinderPersistentVolumeSource     `json:"cinder,omitempty" yaml:"cinder,omitempty"`
	ClaimRef                      *ObjectReference                  `json:"claimRef,omitempty" yaml:"claim_ref,omitempty"`
	FC                            *FCVolumeSource                   `json:"fc,omitempty" yaml:"fc,omitempty"`
	FlexVolume                    *FlexPersistentVolumeSource       `json:"flexVolume,omitempty" yaml:"flex_volume,omitempty"`
	Flocker                       *FlockerVolumeSource              `json:"flocker,omitempty" yaml:"flocker,omitempty"`
	GCEPersistentDisk             *GCEPersistentDiskVolumeSource    `json:"gcePersistentDisk,omitempty" yaml:"gce_persistent_disk,omitempty"`
	Glusterfs                     *GlusterfsVolumeSource            `json:"glusterfs,omitempty" yaml:"glusterfs,omitempty"`
	HostPath                      *HostPathVolumeSource             `json:"hostPath,omitempty" yaml:"host_path,omitempty"`
	ISCSI                         *ISCSIPersistentVolumeSource      `json:"iscsi,omitempty" yaml:"iscsi,omitempty"`
	Local                         *LocalVolumeSource                `json:"local,omitempty" yaml:"local,omitempty"`
	MountOptions                  []string                          `json:"mountOptions,omitempty" yaml:"mount_options,omitempty"`
	NFS                           *NFSVolumeSource                  `json:"nfs,omitempty" yaml:"nfs,omitempty"`
	NodeAffinity                  *VolumeNodeAffinity               `json:"nodeAffinity,omitempty" yaml:"node_affinity,omitempty"`
	PersistentVolumeReclaimPolicy string                            `json:"persistentVolumeReclaimPolicy,omitempty" yaml:"persistent_volume_reclaim_policy,omitempty"`
	PhotonPersistentDisk          *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty" yaml:"photon_persistent_disk,omitempty"`
	PortworxVolume                *PortworxVolumeSource             `json:"portworxVolume,omitempty" yaml:"portworx_volume,omitempty"`
	Quobyte                       *QuobyteVolumeSource              `json:"quobyte,omitempty" yaml:"quobyte,omitempty"`
	RBD                           *RBDPersistentVolumeSource        `json:"rbd,omitempty" yaml:"rbd,omitempty"`
	ScaleIO                       *ScaleIOPersistentVolumeSource    `json:"scaleIO,omitempty" yaml:"scale_io,omitempty"`
	StorageClassID                string                            `json:"storageClassId,omitempty" yaml:"storage_class_id,omitempty"`
	StorageOS                     *StorageOSPersistentVolumeSource  `json:"storageos,omitempty" yaml:"storageos,omitempty"`
	VolumeMode                    string                            `json:"volumeMode,omitempty" yaml:"volume_mode,omitempty"`
	VsphereVolume                 *VsphereVirtualDiskVolumeSource   `json:"vsphereVolume,omitempty" yaml:"vsphere_volume,omitempty"`
}
