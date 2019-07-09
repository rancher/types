package client

const (
	ServiceAccountTokenProjectionType                   = "serviceAccountTokenProjection"
	ServiceAccountTokenProjectionFieldAudience          = "audience"
	ServiceAccountTokenProjectionFieldExpirationSeconds = "expirationSeconds"
	ServiceAccountTokenProjectionFieldPath              = "path"
)

type ServiceAccountTokenProjection struct {
	Audience          string `json:"audience,omitempty" yaml:"audience,omitempty"`
	ExpirationSeconds *int64 `json:"expirationSeconds,omitempty" yaml:"expiration_seconds,omitempty"`
	Path              string `json:"path,omitempty" yaml:"path,omitempty"`
}
