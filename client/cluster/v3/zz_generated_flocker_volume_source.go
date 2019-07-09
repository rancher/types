package client

const (
	FlockerVolumeSourceType             = "flockerVolumeSource"
	FlockerVolumeSourceFieldDatasetName = "datasetName"
	FlockerVolumeSourceFieldDatasetUUID = "datasetUUID"
)

type FlockerVolumeSource struct {
	DatasetName string `json:"datasetName,omitempty" yaml:"dataset_name,omitempty"`
	DatasetUUID string `json:"datasetUUID,omitempty" yaml:"dataset_uuid,omitempty"`
}
