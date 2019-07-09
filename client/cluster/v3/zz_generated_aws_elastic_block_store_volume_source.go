package client

const (
	AWSElasticBlockStoreVolumeSourceType           = "awsElasticBlockStoreVolumeSource"
	AWSElasticBlockStoreVolumeSourceFieldFSType    = "fsType"
	AWSElasticBlockStoreVolumeSourceFieldPartition = "partition"
	AWSElasticBlockStoreVolumeSourceFieldReadOnly  = "readOnly"
	AWSElasticBlockStoreVolumeSourceFieldVolumeID  = "volumeID"
)

type AWSElasticBlockStoreVolumeSource struct {
	FSType    string `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	Partition int64  `json:"partition,omitempty" yaml:"partition,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	VolumeID  string `json:"volumeID,omitempty" yaml:"volume_id,omitempty"`
}
