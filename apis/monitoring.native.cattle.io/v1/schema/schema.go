package schema

import (
	"github.com/rancher/norman/types"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "monitoring.native.cattle.io",
		Path:    "/v3/project",
	}
)
