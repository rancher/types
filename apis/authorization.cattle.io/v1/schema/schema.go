package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/authorization.cattle.io/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Path:    "/v1-authz",
	}

	Schemas = types.NewSchemas().
		MustImport(&Version, v1.Project{}).
		MustImport(&Version, v1.ProjectRoleTemplate{}).
		MustImport(&Version, v1.PodSecurityPolicyTemplate{}).
		MustImport(&Version, v1.ProjectRoleTemplateBinding{}).
		MustImport(&Version, v1.ClusterRoleTemplate{}).
		MustImport(&Version, v1.ClusterRoleTemplateBinding{})
)
