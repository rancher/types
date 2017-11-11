package schema

import (
	"github.com/rancher/norman/types"
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

	Schemas = types.NewSchemas().
		MustImport(&Version, v1.Cluster{}).
		MustImport(&Version, v1.ClusterNode{})
)
