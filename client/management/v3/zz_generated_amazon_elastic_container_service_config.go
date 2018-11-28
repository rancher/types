package client

const (
	AmazonElasticContainerServiceConfigType                             = "amazonElasticContainerServiceConfig"
	AmazonElasticContainerServiceConfigFieldAMI                         = "ami"
	AmazonElasticContainerServiceConfigFieldAccessKey                   = "accessKey"
	AmazonElasticContainerServiceConfigFieldAssociateWorkerNodePublicIP = "associateWorkerNodePublicIp"
	AmazonElasticContainerServiceConfigFieldInstanceType                = "instanceType"
	AmazonElasticContainerServiceConfigFieldMaximumNodes                = "maximumNodes"
	AmazonElasticContainerServiceConfigFieldMinimumNodes                = "minimumNodes"
	AmazonElasticContainerServiceConfigFieldRegion                      = "region"
	AmazonElasticContainerServiceConfigFieldSecretKey                   = "secretKey"
	AmazonElasticContainerServiceConfigFieldSecurityGroups              = "securityGroups"
	AmazonElasticContainerServiceConfigFieldServiceRole                 = "serviceRole"
	AmazonElasticContainerServiceConfigFieldSessionToken                = "sessionToken"
	AmazonElasticContainerServiceConfigFieldSubnets                     = "subnets"
	AmazonElasticContainerServiceConfigFieldVirtualNetwork              = "virtualNetwork"
)

type AmazonElasticContainerServiceConfig struct {
	AMI                         string   `json:"ami,omitempty" yaml:"ami,omitempty"`
	AccessKey                   string   `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	AssociateWorkerNodePublicIP *bool    `json:"associateWorkerNodePublicIp,omitempty" yaml:"associateWorkerNodePublicIp,omitempty"`
	InstanceType                string   `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	MaximumNodes                int64    `json:"maximumNodes,omitempty" yaml:"maximumNodes,omitempty"`
	MinimumNodes                int64    `json:"minimumNodes,omitempty" yaml:"minimumNodes,omitempty"`
	Region                      string   `json:"region,omitempty" yaml:"region,omitempty"`
	SecretKey                   string   `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
	SecurityGroups              []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
	ServiceRole                 string   `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`
	SessionToken                string   `json:"sessionToken,omitempty" yaml:"sessionToken,omitempty"`
	Subnets                     []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
	VirtualNetwork              string   `json:"virtualNetwork,omitempty" yaml:"virtualNetwork,omitempty"`
}
