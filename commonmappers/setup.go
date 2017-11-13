package commonmappers

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapping/mapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Add(version *types.APIVersion, schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(version, metav1.ObjectMeta{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Drop{"generateName"},
				m.Drop{"selfLink"},
				m.Move{From: "uid", To: "uuid"},
				m.Drop{"resourceVersion"},
				m.Drop{"generation"},
				m.Move{From: "creationTimestamp", To: "created"},
				m.Move{From: "deletionTimestamp", To: "removed"},
				m.Drop{"deletionGracePeriodSeconds"},
				//DeletionGracePeriodSecondsMapper{},
				m.Drop{"initializers"},
				m.Drop{"finalizers"},
				m.Drop{"clusterName"},
				m.Drop{"ownerReferences"},
			},
		})
}
