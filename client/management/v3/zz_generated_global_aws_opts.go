package client

const (
	GlobalAwsOptsType                             = "globalAwsOpts"
	GlobalAwsOptsFieldDisableSecurityGroupIngress = "disable-security-group-ingress"
	GlobalAwsOptsFieldDisableStrictZoneCheck      = "disable-strict-zone-check"
	GlobalAwsOptsFieldElbSecurityGroup            = "elb-security-group"
	GlobalAwsOptsFieldKubernetesClusterID         = "kubernetes-cluster-id"
	GlobalAwsOptsFieldKubernetesClusterTag        = "kubernetes-cluster-tag"
	GlobalAwsOptsFieldRoleARN                     = "role-arn"
	GlobalAwsOptsFieldRouteTableID                = "routetable-id"
	GlobalAwsOptsFieldSubnetID                    = "subnet-id"
	GlobalAwsOptsFieldVPC                         = "vpc"
	GlobalAwsOptsFieldZone                        = "zone"
)

type GlobalAwsOpts struct {
	DisableSecurityGroupIngress bool   `json:"disable-security-group-ingress,omitempty" yaml:"disable_security_group_ingress,omitempty"`
	DisableStrictZoneCheck      bool   `json:"disable-strict-zone-check,omitempty" yaml:"disable_strict_zone_check,omitempty"`
	ElbSecurityGroup            string `json:"elb-security-group,omitempty" yaml:"elb_security_group,omitempty"`
	KubernetesClusterID         string `json:"kubernetes-cluster-id,omitempty" yaml:"kubernetes_cluster_id,omitempty"`
	KubernetesClusterTag        string `json:"kubernetes-cluster-tag,omitempty" yaml:"kubernetes_cluster_tag,omitempty"`
	RoleARN                     string `json:"role-arn,omitempty" yaml:"role_arn,omitempty"`
	RouteTableID                string `json:"routetable-id,omitempty" yaml:"routetable_id,omitempty"`
	SubnetID                    string `json:"subnet-id,omitempty" yaml:"subnet_id,omitempty"`
	VPC                         string `json:"vpc,omitempty" yaml:"vpc,omitempty"`
	Zone                        string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
