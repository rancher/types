package client

const (
	AzureCloudProviderType                              = "azureCloudProvider"
	AzureCloudProviderFieldAADClientCertPassword        = "aadClientCertPassword"
	AzureCloudProviderFieldAADClientCertPath            = "aadClientCertPath"
	AzureCloudProviderFieldAADClientID                  = "aadClientId"
	AzureCloudProviderFieldAADClientSecret              = "aadClientSecret"
	AzureCloudProviderFieldCloud                        = "cloud"
	AzureCloudProviderFieldCloudProviderBackoff         = "cloudProviderBackoff"
	AzureCloudProviderFieldCloudProviderBackoffDuration = "cloudProviderBackoffDuration"
	AzureCloudProviderFieldCloudProviderBackoffExponent = "cloudProviderBackoffExponent"
	AzureCloudProviderFieldCloudProviderBackoffJitter   = "cloudProviderBackoffJitter"
	AzureCloudProviderFieldCloudProviderBackoffRetries  = "cloudProviderBackoffRetries"
	AzureCloudProviderFieldCloudProviderRateLimit       = "cloudProviderRateLimit"
	AzureCloudProviderFieldCloudProviderRateLimitBucket = "cloudProviderRateLimitBucket"
	AzureCloudProviderFieldCloudProviderRateLimitQPS    = "cloudProviderRateLimitQPS"
	AzureCloudProviderFieldExcludeMasterFromStandardLB  = "excludeMasterFromStandardLB"
	AzureCloudProviderFieldLoadBalancerSku              = "loadBalancerSku"
	AzureCloudProviderFieldLocation                     = "location"
	AzureCloudProviderFieldMaximumLoadBalancerRuleCount = "maximumLoadBalancerRuleCount"
	AzureCloudProviderFieldPrimaryAvailabilitySetName   = "primaryAvailabilitySetName"
	AzureCloudProviderFieldPrimaryScaleSetName          = "primaryScaleSetName"
	AzureCloudProviderFieldResourceGroup                = "resourceGroup"
	AzureCloudProviderFieldRouteTableName               = "routeTableName"
	AzureCloudProviderFieldSecurityGroupName            = "securityGroupName"
	AzureCloudProviderFieldSubnetName                   = "subnetName"
	AzureCloudProviderFieldSubscriptionID               = "subscriptionId"
	AzureCloudProviderFieldTenantID                     = "tenantId"
	AzureCloudProviderFieldUseInstanceMetadata          = "useInstanceMetadata"
	AzureCloudProviderFieldUseManagedIdentityExtension  = "useManagedIdentityExtension"
	AzureCloudProviderFieldUserAssignedIdentityID       = "userAssignedIdentityID"
	AzureCloudProviderFieldVMType                       = "vmType"
	AzureCloudProviderFieldVnetName                     = "vnetName"
	AzureCloudProviderFieldVnetResourceGroup            = "vnetResourceGroup"
)

type AzureCloudProvider struct {
	AADClientCertPassword        string `json:"aadClientCertPassword,omitempty" yaml:"aad_client_cert_password,omitempty"`
	AADClientCertPath            string `json:"aadClientCertPath,omitempty" yaml:"aad_client_cert_path,omitempty"`
	AADClientID                  string `json:"aadClientId,omitempty" yaml:"aad_client_id,omitempty"`
	AADClientSecret              string `json:"aadClientSecret,omitempty" yaml:"aad_client_secret,omitempty"`
	Cloud                        string `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	CloudProviderBackoff         bool   `json:"cloudProviderBackoff,omitempty" yaml:"cloud_provider_backoff,omitempty"`
	CloudProviderBackoffDuration int64  `json:"cloudProviderBackoffDuration,omitempty" yaml:"cloud_provider_backoff_duration,omitempty"`
	CloudProviderBackoffExponent int64  `json:"cloudProviderBackoffExponent,omitempty" yaml:"cloud_provider_backoff_exponent,omitempty"`
	CloudProviderBackoffJitter   int64  `json:"cloudProviderBackoffJitter,omitempty" yaml:"cloud_provider_backoff_jitter,omitempty"`
	CloudProviderBackoffRetries  int64  `json:"cloudProviderBackoffRetries,omitempty" yaml:"cloud_provider_backoff_retries,omitempty"`
	CloudProviderRateLimit       bool   `json:"cloudProviderRateLimit,omitempty" yaml:"cloud_provider_rate_limit,omitempty"`
	CloudProviderRateLimitBucket int64  `json:"cloudProviderRateLimitBucket,omitempty" yaml:"cloud_provider_rate_limit_bucket,omitempty"`
	CloudProviderRateLimitQPS    int64  `json:"cloudProviderRateLimitQPS,omitempty" yaml:"cloud_provider_rate_limit_qps,omitempty"`
	ExcludeMasterFromStandardLB  *bool  `json:"excludeMasterFromStandardLB,omitempty" yaml:"exclude_master_from_standard_lb,omitempty"`
	LoadBalancerSku              string `json:"loadBalancerSku,omitempty" yaml:"load_balancer_sku,omitempty"`
	Location                     string `json:"location,omitempty" yaml:"location,omitempty"`
	MaximumLoadBalancerRuleCount int64  `json:"maximumLoadBalancerRuleCount,omitempty" yaml:"maximum_load_balancer_rule_count,omitempty"`
	PrimaryAvailabilitySetName   string `json:"primaryAvailabilitySetName,omitempty" yaml:"primary_availability_set_name,omitempty"`
	PrimaryScaleSetName          string `json:"primaryScaleSetName,omitempty" yaml:"primary_scale_set_name,omitempty"`
	ResourceGroup                string `json:"resourceGroup,omitempty" yaml:"resource_group,omitempty"`
	RouteTableName               string `json:"routeTableName,omitempty" yaml:"route_table_name,omitempty"`
	SecurityGroupName            string `json:"securityGroupName,omitempty" yaml:"security_group_name,omitempty"`
	SubnetName                   string `json:"subnetName,omitempty" yaml:"subnet_name,omitempty"`
	SubscriptionID               string `json:"subscriptionId,omitempty" yaml:"subscription_id,omitempty"`
	TenantID                     string `json:"tenantId,omitempty" yaml:"tenant_id,omitempty"`
	UseInstanceMetadata          bool   `json:"useInstanceMetadata,omitempty" yaml:"use_instance_metadata,omitempty"`
	UseManagedIdentityExtension  bool   `json:"useManagedIdentityExtension,omitempty" yaml:"use_managed_identity_extension,omitempty"`
	UserAssignedIdentityID       string `json:"userAssignedIdentityID,omitempty" yaml:"user_assigned_identity_id,omitempty"`
	VMType                       string `json:"vmType,omitempty" yaml:"vm_type,omitempty"`
	VnetName                     string `json:"vnetName,omitempty" yaml:"vnet_name,omitempty"`
	VnetResourceGroup            string `json:"vnetResourceGroup,omitempty" yaml:"vnet_resource_group,omitempty"`
}
