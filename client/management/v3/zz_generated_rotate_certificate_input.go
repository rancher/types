package client

const (
	RotateCertificateInputType                = "rotateCertificateInput"
	RotateCertificateInputFieldCACertificates = "caCertificates"
	RotateCertificateInputFieldServices       = "services"
)

type RotateCertificateInput struct {
	CACertificates bool   `json:"caCertificates,omitempty" yaml:"ca_certificates,omitempty"`
	Services       string `json:"services,omitempty" yaml:"services,omitempty"`
}
