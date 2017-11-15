package mapper

import (
	"strings"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/mapping/mapper"
)

type OSInfo struct {
}

func (o OSInfo) FromInternal(data map[string]interface{}) {
	cpuInfo := map[string]interface{}{
		"count": mapper.GetValueN(data, "capacity", "cpu"),
	}

	kib := strings.TrimSuffix(convert.ToString(mapper.GetValueN(data, "capacity", "memory")), "Ki")
	memoryInfo := map[string]interface{}{}

	kibNum, err := convert.ToNumber(kib)
	if err == nil {
		memoryInfo["memTotalKiB"] = kibNum
	}

	osInfo := map[string]interface{}{
		"dockerVersion":   strings.TrimPrefix(convert.ToString(mapper.GetValueN(data, "nodeInfo", "containerRuntimeVersion")), "docker://"),
		"kernelVersion":   mapper.GetValueN(data, "nodeInfo", "kernelVersion"),
		"operatingSystem": mapper.GetValueN(data, "nodeInfo", "osImage"),
	}

	data["info"] = map[string]interface{}{
		"cpu":    cpuInfo,
		"memory": memoryInfo,
		"os":     osInfo,
		"kubernetes": map[string]interface{}{
			"kubeletVersion":   mapper.GetValueN(data, "nodeInfo", "kubeletVersion"),
			"kubeProxyVersion": mapper.GetValueN(data, "nodeInfo", "kubeletVersion"),
		},
	}
}

func (o OSInfo) ToInternal(data map[string]interface{}) {
}

func (o OSInfo) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
