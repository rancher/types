package client

const (
	Route53ProviderConfigType                 = "route53ProviderConfig"
	Route53ProviderConfigFieldAccessKey       = "accessKey"
	Route53ProviderConfigFieldCredentialsPath = "credentialsPath"
	Route53ProviderConfigFieldRegion          = "region"
	Route53ProviderConfigFieldRoleArn         = "roleArn"
	Route53ProviderConfigFieldSecretKey       = "secretKey"
	Route53ProviderConfigFieldZoneType        = "zoneType"
)

type Route53ProviderConfig struct {
	AccessKey       string `json:"accessKey,omitempty" yaml:"access_key,omitempty"`
	CredentialsPath string `json:"credentialsPath,omitempty" yaml:"credentials_path,omitempty"`
	Region          string `json:"region,omitempty" yaml:"region,omitempty"`
	RoleArn         string `json:"roleArn,omitempty" yaml:"role_arn,omitempty"`
	SecretKey       string `json:"secretKey,omitempty" yaml:"secret_key,omitempty"`
	ZoneType        string `json:"zoneType,omitempty" yaml:"zone_type,omitempty"`
}
