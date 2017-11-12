package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/types/io.cattle.authorization/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "io.cattle.authorization",
		Path:    "/v1-authz",
	}

	Schemas = types.NewSchemas().
		MustImport(&Version, v1.Project{}).
		MustImport(&Version, v1.RoleTemplate{}).
		MustImport(&Version, v1.PodSecurityPolicyTemplate{}).
		MustImport(&Version, v1.ProjectRoleBinding{})
)
