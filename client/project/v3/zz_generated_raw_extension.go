package client

const (
	RawExtensionType     = "rawExtension"
	RawExtensionFieldRaw = "raw"
)

type RawExtension struct {
	Raw string `json:"raw,omitempty" yaml:"raw,omitempty"`
}
