package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapping/mapper"
	"github.com/rancher/types/apis/cluster.cattle.io/v1"
	"github.com/rancher/types/commonmappers"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "cluster.cattle.io",
		Path:    "/v1-cluster",
		SubContexts: map[string]bool{
			"clusters": true,
		},
	}

	Schemas = commonmappers.Add(&Version, types.NewSchemas()).
		AddMapperForType(&Version, v1.Cluster{}, m.NewObject()).
		AddMapperForType(&Version, v1.ClusterNode{}, m.NewObject()).
		MustImport(&Version, v1.Cluster{}).
		MustImport(&Version, v1.ClusterNode{})
)
