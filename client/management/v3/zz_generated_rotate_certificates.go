package client

const (
	RotateCertificatesType                = "rotateCertificates"
	RotateCertificatesFieldCACertificates = "caCertificates"
	RotateCertificatesFieldServices       = "services"
)

type RotateCertificates struct {
	CACertificates bool   `json:"caCertificates,omitempty" yaml:"ca_certificates,omitempty"`
	Services       string `json:"services,omitempty" yaml:"services,omitempty"`
}
