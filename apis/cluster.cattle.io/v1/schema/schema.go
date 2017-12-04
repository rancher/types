package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/mapper"
	"github.com/rancher/types/apis/cluster.cattle.io/v1"
	"github.com/rancher/types/factory"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "cluster.cattle.io",
		Path:    "/v1-cluster",
	}

	Schemas = factory.Schemas(&Version).
		AddMapperForType(&Version, v1.Cluster{}, mapper.DisplayName{}).
		AddMapperForType(&Version, v1.Machine{}, mapper.DisplayName{}).
		AddMapperForType(&Version, v1.MachineDriver{}, mapper.DisplayName{}).
		AddMapperForType(&Version, v1.MachineTemplate{}, mapper.DisplayName{}).
		MustImport(&Version, v1.Cluster{}).
		MustImport(&Version, v1.ClusterNode{}).
		MustImport(&Version, v1.Machine{}).
		MustImport(&Version, v1.MachineDriver{}).
		MustImport(&Version, v1.MachineTemplate{}).
		MustImport(&Version, v1.ClusterEvent{})
)
