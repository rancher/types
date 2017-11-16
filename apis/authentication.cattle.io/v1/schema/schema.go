package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/authentication.cattle.io/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "authentication.cattle.io",
		Path:    "/v1-authn",
	}

	Schemas = types.NewSchemas().
		MustImport(&Version, v1.Token{})
)
