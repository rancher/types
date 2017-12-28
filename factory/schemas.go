package factory

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/factory"
	"github.com/rancher/types/mapper"
)

func Schemas(version *types.APIVersion) *types.Schemas {
	schemas := factory.Schemas(version)
	baseFunc := schemas.DefaultMappers
	schemas.DefaultMappers = func() []types.Mapper {
		mappers := append([]types.Mapper{
			&mapper.Status{},
		}, baseFunc()...)
		mappers = append(mappers, &m.Scope{
			If: types.NamespaceScope,
			Mappers: []types.Mapper{
				&mapper.NamespaceIDMapper{},
			},
		})
		return mappers
	}
	}
	return schemas
}
