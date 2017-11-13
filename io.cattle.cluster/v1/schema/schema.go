package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapping/mapper"
	"github.com/rancher/types/commonmappers"
	"github.com/rancher/types/io.cattle.cluster/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "io.cattle.cluster",
		Path:    "/v1-cluster",
		SubContexts: map[string]bool{
			"projects": true,
		},
	}

	Schemas = commonmappers.Add(&Version, types.NewSchemas()).
		AddMapperForType(&Version, v1.Cluster{}, m.NewObject(nil)).
		AddMapperForType(&Version, v1.ClusterNode{}, m.NewObject(nil)).
		MustImport(&Version, v1.Cluster{}).
		MustImport(&Version, v1.ClusterNode{})
)
