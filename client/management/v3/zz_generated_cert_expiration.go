package client

const (
	CertExpirationType                = "certExpiration"
	CertExpirationFieldDaysRemaining  = "daysRemaining"
	CertExpirationFieldExpirationDate = "expirationDate"
)

type CertExpiration struct {
	DaysRemaining  int64  `json:"daysRemaining,omitempty" yaml:"daysRemaining,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty" yaml:"expirationDate,omitempty"`
}
