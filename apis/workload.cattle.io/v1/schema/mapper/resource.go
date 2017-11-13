package mapper

import (
	"strings"

	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapping/mapper"
)

type ResourceRequirementsMapper struct {
}

func (r ResourceRequirementsMapper) FromInternal(data map[string]interface{}) {
	for key, value := range data {
		mapValue, ok := value.(map[string]interface{})
		if !ok {
			continue
		}
		for subKey, subValue := range mapValue {
			m.PutValue(data, subValue, subKey, strings.TrimSuffix(key, "s"))

		}
		delete(data, key)
	}
}

func (r ResourceRequirementsMapper) ToInternal(data map[string]interface{}) {
	for key, value := range data {
		mapValue, ok := value.(map[string]interface{})
		if !ok {
			continue
		}
		for subKey, subValue := range mapValue {
			m.PutValue(data, subValue, subKey, key+"s")

		}
		delete(data, key)
	}
}

func (r ResourceRequirementsMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
