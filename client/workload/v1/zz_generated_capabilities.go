package client

const (
	CapabilitiesType      = "capabilities"
	CapabilitiesFieldAdd  = "add"
	CapabilitiesFieldDrop = "drop"
)

type Capabilities struct {
	Add  []string `json:"add,omitempty"`
	Drop []string `json:"drop,omitempty"`
}
