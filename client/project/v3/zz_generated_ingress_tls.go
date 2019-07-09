package client

const (
	IngressTLSType               = "ingressTLS"
	IngressTLSFieldCertificateID = "certificateId"
	IngressTLSFieldHosts         = "hosts"
)

type IngressTLS struct {
	CertificateID string   `json:"certificateId,omitempty" yaml:"certificate_id,omitempty"`
	Hosts         []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`
}
