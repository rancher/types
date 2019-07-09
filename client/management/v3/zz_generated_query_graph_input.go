package client

const (
	QueryGraphInputType              = "queryGraphInput"
	QueryGraphInputFieldFilters      = "filters"
	QueryGraphInputFieldFrom         = "from"
	QueryGraphInputFieldInterval     = "interval"
	QueryGraphInputFieldIsDetails    = "isDetails"
	QueryGraphInputFieldMetricParams = "metricParams"
	QueryGraphInputFieldTo           = "to"
)

type QueryGraphInput struct {
	Filters      map[string]string `json:"filters,omitempty" yaml:"filters,omitempty"`
	From         string            `json:"from,omitempty" yaml:"from,omitempty"`
	Interval     string            `json:"interval,omitempty" yaml:"interval,omitempty"`
	IsDetails    bool              `json:"isDetails,omitempty" yaml:"is_details,omitempty"`
	MetricParams map[string]string `json:"metricParams,omitempty" yaml:"metric_params,omitempty"`
	To           string            `json:"to,omitempty" yaml:"to,omitempty"`
}
