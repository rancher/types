package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/factory"
	"github.com/rancher/types/apis/catalog.cattle.io/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "catalog.cattle.io",
		Path:    "/v1-catalog",
	}

	Schemas = factory.Schemas(&Version).
		MustImport(&Version, v1.Catalog{}).
		MustImport(&Version, v1.Template{}).
		MustImport(&Version, v1.TemplateVersion{})
)
