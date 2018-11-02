package client

const (
	ClusterConfigType                          = "clusterConfig"
	ClusterConfigFieldCertificateAuthorityData = "certificateAuthorityData"
	ClusterConfigFieldCertificateAuthorityPath = "certificateAuthorityPath"
	ClusterConfigFieldClientCertificateData    = "clientCertificateData"
	ClusterConfigFieldClientCertificatePath    = "clientCertificatePath"
	ClusterConfigFieldClientKeyData            = "clientKeyData"
	ClusterConfigFieldClientKeyPath            = "clientKeyPath"
	ClusterConfigFieldServer                   = "server"
	ClusterConfigFieldTokenFile                = "tokenFile"
)

type ClusterConfig struct {
	CertificateAuthorityData string `json:"certificateAuthorityData,omitempty" yaml:"certificateAuthorityData,omitempty"`
	CertificateAuthorityPath string `json:"certificateAuthorityPath,omitempty" yaml:"certificateAuthorityPath,omitempty"`
	ClientCertificateData    string `json:"clientCertificateData,omitempty" yaml:"clientCertificateData,omitempty"`
	ClientCertificatePath    string `json:"clientCertificatePath,omitempty" yaml:"clientCertificatePath,omitempty"`
	ClientKeyData            string `json:"clientKeyData,omitempty" yaml:"clientKeyData,omitempty"`
	ClientKeyPath            string `json:"clientKeyPath,omitempty" yaml:"clientKeyPath,omitempty"`
	Server                   string `json:"server,omitempty" yaml:"server,omitempty"`
	TokenFile                string `json:"tokenFile,omitempty" yaml:"tokenFile,omitempty"`
}
