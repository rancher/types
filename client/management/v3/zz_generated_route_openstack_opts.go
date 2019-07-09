package client

const (
	RouteOpenstackOptsType          = "routeOpenstackOpts"
	RouteOpenstackOptsFieldRouterID = "router-id"
)

type RouteOpenstackOpts struct {
	RouterID string `json:"router-id,omitempty" yaml:"router_id,omitempty"`
}
