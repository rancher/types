package client

const (
	CloudProviderType                        = "cloudProvider"
	CloudProviderFieldAWSCloudProvider       = "awsCloudProvider"
	CloudProviderFieldAzureCloudProvider     = "azureCloudProvider"
	CloudProviderFieldCustomCloudProvider    = "customCloudProvider"
	CloudProviderFieldName                   = "name"
	CloudProviderFieldOpenstackCloudProvider = "openstackCloudProvider"
	CloudProviderFieldVsphereCloudProvider   = "vsphereCloudProvider"
)

type CloudProvider struct {
	AWSCloudProvider       *AWSCloudProvider       `json:"awsCloudProvider,omitempty" yaml:"aws_cloud_provider,omitempty"`
	AzureCloudProvider     *AzureCloudProvider     `json:"azureCloudProvider,omitempty" yaml:"azure_cloud_provider,omitempty"`
	CustomCloudProvider    string                  `json:"customCloudProvider,omitempty" yaml:"custom_cloud_provider,omitempty"`
	Name                   string                  `json:"name,omitempty" yaml:"name,omitempty"`
	OpenstackCloudProvider *OpenstackCloudProvider `json:"openstackCloudProvider,omitempty" yaml:"openstack_cloud_provider,omitempty"`
	VsphereCloudProvider   *VsphereCloudProvider   `json:"vsphereCloudProvider,omitempty" yaml:"vsphere_cloud_provider,omitempty"`
}
