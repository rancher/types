package client

const (
	RKEConfigHostType                    = "rkeConfigHost"
	RKEConfigHostFieldAdvertiseAddress   = "advertiseAddress"
	RKEConfigHostFieldAdvertisedHostname = "advertisedHostname"
	RKEConfigHostFieldDockerSocket       = "dockerSocket"
	RKEConfigHostFieldIP                 = "ip"
	RKEConfigHostFieldRole               = "role"
	RKEConfigHostFieldUser               = "user"
)

type RKEConfigHost struct {
	AdvertiseAddress   string   `json:"advertiseAddress,omitempty"`
	AdvertisedHostname string   `json:"advertisedHostname,omitempty"`
	DockerSocket       string   `json:"dockerSocket,omitempty"`
	IP                 string   `json:"ip,omitempty"`
	Role               []string `json:"role,omitempty"`
	User               string   `json:"user,omitempty"`
}
