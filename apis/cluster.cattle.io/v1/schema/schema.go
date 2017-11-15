package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/factory"
	"github.com/rancher/types/apis/cluster.cattle.io/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "cluster.cattle.io",
		Path:    "/v1-cluster",
	}

	Schemas = factory.Schemas(&Version).
		MustImport(&Version, v1.Cluster{}).
		MustImport(&Version, v1.ClusterNode{})
)
