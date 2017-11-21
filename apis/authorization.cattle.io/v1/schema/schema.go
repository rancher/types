package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/factory"
	"github.com/rancher/norman/types/mapper"
	"github.com/rancher/types/apis/authorization.cattle.io/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Path:    "/v1-authz",
	}

	Schemas = factory.Schemas(&Version).
		AddMapperForType(&Version, v1.Project{},
			mapper.DisplayName{},
		).
		AddMapperForType(&Version, v1.ProjectRoleTemplateBinding{},
			&mapper.Move{From: "subject/name", To: "subjectName"},
			&mapper.Move{From: "subject/kind", To: "subjectKind"},
			&mapper.Move{From: "subject/namespace", To: "subjectNamespace"},
			&mapper.Drop{Field: "subject"},
		).
		MustImportAndCustomize(&Version, v1.Project{}, func(schema *types.Schema) {
			schema.SubContext = "projects"
		}).
		MustImport(&Version, v1.ProjectRoleTemplate{}).
		MustImport(&Version, v1.ProjectRoleTemplateBinding{}).
		MustImport(&Version, v1.PodSecurityPolicyTemplate{}).
		MustImport(&Version, v1.ClusterRoleTemplate{}).
		MustImport(&Version, v1.ClusterRoleTemplateBinding{}).
		MustImportAndCustomize(&Version, v1.ProjectRoleTemplateBinding{}, func(schema *types.Schema) {
			schema.MustCustomizeField("subjectKind", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"User", "Group", "ServiceAccount"}
				field.Nullable = false
				return field
			})
		})
)
