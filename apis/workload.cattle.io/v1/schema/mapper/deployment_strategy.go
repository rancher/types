package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	m "github.com/rancher/norman/types/mapper"
)

type DeploymentStrategyMapper struct {
}

func (d DeploymentStrategyMapper) FromInternal(data map[string]interface{}) {
	if m.GetValueN(data, "strategy", "type") != "Recreate" {
		m.PutValue(data, "Parallel", "deploymentStrategy", "kind")
		maxUnavailable := m.GetValueN(data, "strategy", "rollingUpdate", "maxUnavailable")
		maxSurge := m.GetValueN(data, "strategy", "rollingUpdate", "maxSurge")
		if !convert.IsEmpty(maxSurge) {
			m.PutValue(data, true, "deploymentStrategy", "parallelConfig", "startFirst")
			m.PutValue(data, convert.ToString(maxSurge), "batchSize")
		} else if !convert.IsEmpty(maxUnavailable) {
			m.PutValue(data, convert.ToString(maxUnavailable), "batchSize")
		}
	}
}

func (d DeploymentStrategyMapper) ToInternal(data map[string]interface{}) {
	batchSize := m.GetValueN(data, "batchSize")
	if convert.IsEmpty(batchSize) {
		batchSize = 1
	}

	batchSize, _ = convert.ToNumber(batchSize)

	kind, _ := m.GetValueN(data, "deploymentStrategy", "kind").(string)
	if kind == "" || kind == "Parallel" {
		startFirst, _ := m.GetValueN(data, "deploymentStrategy", "startFirst").(bool)
		if startFirst {
			m.PutValue(data, batchSize, "strategy", "rollingUpdate", "maxSurge")
		} else {
			m.PutValue(data, batchSize, "strategy", "rollingUpdate", "maxUnavailable")
		}
	}
}

func (d DeploymentStrategyMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
