package client

const (
	QueryClusterGraphType           = "queryClusterGraph"
	QueryClusterGraphFieldGraphName = "graphID"
	QueryClusterGraphFieldSeries    = "series"
)

type QueryClusterGraph struct {
	GraphName string   `json:"graphID,omitempty" yaml:"graph_id,omitempty"`
	Series    []string `json:"series,omitempty" yaml:"series,omitempty"`
}
