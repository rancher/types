package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapping/mapper"
)

var (
	handlerMapper = &types.TypeMapper{
		Mappers: []types.Mapper{
			&m.UnionEmbed{
				Fields: []m.UnionMapping{
					{
						FieldName:   "exec",
						CheckFields: []string{"command"},
					},
					{
						FieldName:   "tcpSocket",
						CheckFields: []string{"tcp", "port"},
					},
					{
						FieldName:   "httpGet",
						CheckFields: []string{"port"},
					},
				},
			},
		},
	}
)

type handlerOverride struct {
	TCP bool
}

type EnvironmentFrom struct {
	Source     string
	SourceName string
	SourceKey  string
	Prefix     string
	Optional   bool
	TargetKey  string
}

type Resources struct {
	CPU       *ResourceRequest
	Memory    *ResourceRequest
	NvidiaGPU *ResourceRequest
}

type ResourceRequest struct {
	Request string
	Limit   string
}

type Scheduling struct {
	AntiAffinity      string
	Node              *NodeScheduling
	Tolerate          []string
	Scheduler         string
	Priority          *int64
	PriorityClassName string
}

type NodeScheduling struct {
	Name       string
	RequireAll []string
	RequireAny []string
	Preferred  []string
}

type deployParams struct {
	BatchSize  int64
	Scale      int64
	Global     bool
	Cron       string
	Job        bool
	Ordered    bool
	QuorumSize int64
}
