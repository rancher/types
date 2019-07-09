package client

const (
	CertExpirationType                = "certExpiration"
	CertExpirationFieldExpirationDate = "expirationDate"
)

type CertExpiration struct {
	ExpirationDate string `json:"expirationDate,omitempty" yaml:"expiration_date,omitempty"`
}
