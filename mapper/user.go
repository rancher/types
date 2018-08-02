package mapper

import (
	"github.com/rancher/norman/types"
)

type UserMapper struct {
}

func (u UserMapper) FromInternal(data map[string]interface{}) {}

func (u UserMapper) ToInternal(data map[string]interface{}) error {
	if _, ok := data["conditions"]; !ok {
		data["conditions"] = []map[string]interface{}{{"status": "True", "type": "InitialRolesPopulated"}}
	}
	return nil
}

func (u UserMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
