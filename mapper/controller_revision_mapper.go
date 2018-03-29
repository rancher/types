package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/values"
)

type ControllerRevisionMapper struct {
}

func (m ControllerRevisionMapper) FromInternal(data map[string]interface{}) {
	spec := convert.ToMapInterface(values.GetValueN(data, "data", "spec", "template", "spec"))
	data = spec
	data["type"] = "/v3/project/schemas/podTemplateSpec"
}

func (m ControllerRevisionMapper) ToInternal(data map[string]interface{}) {
}

func (m ControllerRevisionMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
