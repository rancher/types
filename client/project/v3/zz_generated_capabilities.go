package client

const (
	CapabilitiesType         = "capabilities"
	CapabilitiesFieldCapAdd  = "capAdd"
	CapabilitiesFieldCapDrop = "capDrop"
)

type Capabilities struct {
	CapAdd  []string `json:"capAdd,omitempty" yaml:"cap_add,omitempty"`
	CapDrop []string `json:"capDrop,omitempty" yaml:"cap_drop,omitempty"`
}
