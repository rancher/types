package client

const (
	NodeSchedulingType            = "nodeScheduling"
	NodeSchedulingFieldNodeID     = "nodeId"
	NodeSchedulingFieldPreferred  = "preferred"
	NodeSchedulingFieldRequireAll = "requireAll"
	NodeSchedulingFieldRequireAny = "requireAny"
)

type NodeScheduling struct {
	NodeID     string   `json:"nodeId,omitempty" yaml:"node_id,omitempty"`
	Preferred  []string `json:"preferred,omitempty" yaml:"preferred,omitempty"`
	RequireAll []string `json:"requireAll,omitempty" yaml:"require_all,omitempty"`
	RequireAny []string `json:"requireAny,omitempty" yaml:"require_any,omitempty"`
}
