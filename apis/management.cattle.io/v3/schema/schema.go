package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapper"
	"github.com/rancher/types/apis/cluster.cattle.io/v3/schema"
	"github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/factory"
	"github.com/rancher/types/mapper"
)

var (
	Version = types.APIVersion{
		Version: "v3",
		Group:   "management.cattle.io",
		Path:    "/v3",
		SubContexts: map[string]bool{
			"clusters": true,
		},
	}

	Schemas = factory.Schemas(&Version).
		Init(nodeTypes).
		Init(machineTypes).
		Init(authzTypes).
		Init(clusterTypes).
		Init(catalogTypes).
		Init(authnTypes).
		Init(schemaTypes).
		Init(stackTypes).
		Init(userTypes)
)

func schemaTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.DynamicSchema{})
}

func catalogTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.Catalog{}).
		MustImport(&Version, v3.Template{}).
		MustImport(&Version, v3.TemplateVersion{})
}

func nodeTypes(schemas *types.Schemas) *types.Schemas {
	return schema.NodeTypes(&Version, schemas)
}

func clusterTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.Cluster{},
			&m.Embed{Field: "status"},
			m.DisplayName{},
		).
		AddMapperForType(&Version, v3.ClusterStatus{},
			m.Drop{Field: "serviceAccountToken"},
			m.Drop{Field: "appliedSpec"},
			m.Drop{Field: "clusterName"},
		).
		AddMapperForType(&Version, v3.ClusterEvent{}, &m.Move{
			From: "type",
			To:   "eventType",
		}).
		MustImportAndCustomize(&Version, v3.Cluster{}, func(schema *types.Schema) {
			schema.SubContext = "clusters"
		}).
		MustImport(&Version, v3.ClusterEvent{}).
		MustImport(&Version, v3.ClusterRegistrationToken{})
}

func authzTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.ProjectStatus{}).
		AddMapperForType(&Version, v3.Project{}, m.DisplayName{},
			&m.Embed{Field: "status"}).
		AddMapperForType(&Version, v3.GlobalRole{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.RoleTemplate{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.ProjectRoleTemplateBinding{},
			&m.Move{From: "subject/name", To: "subjectName"},
			&m.Move{From: "subject/kind", To: "subjectKind"},
			&m.Move{From: "subject/namespace", To: "subjectNamespace"},
			&m.Drop{Field: "subject"},
			&mapper.NamespaceIDMapper{},
		).
		AddMapperForType(&Version, v3.ClusterRoleTemplateBinding{},
			&m.Move{From: "subject/name", To: "subjectName"},
			&m.Move{From: "subject/kind", To: "subjectKind"},
			&m.Move{From: "subject/namespace", To: "subjectNamespace"},
			&m.Drop{Field: "subject"},
		).
		AddMapperForType(&Version, v3.GlobalRoleBinding{},
			&m.Move{From: "subject/name", To: "subjectName"},
			&m.Move{From: "subject/kind", To: "subjectKind"},
			&m.Drop{Field: "subject"},
		).
		MustImportAndCustomize(&Version, v3.Project{}, func(schema *types.Schema) {
			schema.SubContext = "projects"
		}).
		MustImport(&Version, v3.GlobalRole{}).
		MustImport(&Version, v3.GlobalRoleBinding{}).
		MustImportAndCustomize(&Version, v3.RoleTemplate{}, func(schema *types.Schema) {
			schema.MustCustomizeField("context", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"cluster", "project"}
				field.Nullable = false
				return field
			})
		}).
		MustImport(&Version, v3.PodSecurityPolicyTemplate{}).
		MustImportAndCustomize(&Version, v3.ClusterRoleTemplateBinding{}, func(schema *types.Schema) {
			schema.MustCustomizeField("subjectKind", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"User", "Group", "ServiceAccount", "Principal"}
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("subjectName", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("roleTemplateId", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.ProjectRoleTemplateBinding{}, func(schema *types.Schema) {
			schema.MustCustomizeField("subjectKind", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"User", "Group", "ServiceAccount", "Principal"}
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("subjectName", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("roleTemplateId", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.GlobalRoleBinding{}, func(schema *types.Schema) {
			schema.MustCustomizeField("subjectKind", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"User", "Group", "Principal"}
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("subjectName", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("globalRoleId", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
		})
}

func machineTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.MachineSpec{}, &m.Embed{Field: "nodeSpec"}).
		AddMapperForType(&Version, v3.MachineStatus{},
			&m.Drop{Field: "token"},
			&m.Drop{Field: "rkeNode"},
			&m.Drop{Field: "machineTemplateSpec"},
			&m.Drop{Field: "machineDriverConfig"},
			&m.Embed{Field: "nodeStatus"},
			&m.SliceMerge{From: []string{"conditions", "nodeConditions"}, To: "conditions"}).
		AddMapperForType(&Version, v3.MachineConfig{},
			&m.Drop{Field: "clusterName"}).
		AddMapperForType(&Version, v3.Machine{},
			&m.Embed{Field: "status"},
			m.DisplayName{}).
		AddMapperForType(&Version, v3.MachineDriver{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.MachineTemplate{}, m.DisplayName{}).
		MustImport(&Version, v3.Machine{}).
		MustImportAndCustomize(&Version, v3.MachineDriver{}, func(schema *types.Schema) {
			schema.ResourceActions["activate"] = types.Action{
				Output: "machineDriver",
			}
			schema.ResourceActions["deactivate"] = types.Action{
				Output: "machineDriver",
			}
		}).
		MustImport(&Version, v3.MachineTemplate{})
}

func authnTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.User{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.Group{}, m.DisplayName{}).
		MustImport(&Version, v3.Group{}).
		MustImport(&Version, v3.GroupMember{}).
		MustImport(&Version, v3.Principal{}).
		MustImport(&Version, v3.LoginInput{}).
		MustImport(&Version, v3.LocalCredential{}).
		MustImport(&Version, v3.GithubCredential{}).
		MustImport(&Version, v3.ChangePasswordInput{}).
		MustImportAndCustomize(&Version, v3.Token{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"login": {
					Input:  "loginInput",
					Output: "token",
				},
				"logout": {},
			}
		}).
		MustImportAndCustomize(&Version, v3.User{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"changepassword": {
					Input:  "changePasswordInput",
					Output: "user",
				},
			}
		})
}

func stackTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		MustImportAndCustomize(&Version, v3.Stack{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"upgrade": {
					Input: "templateVersionId",
				},
				"rollback": {
					Input: "revision",
				},
			}
		})
}

func userTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		AddMapperForType(&Version, v3.Preference{}).
		MustImportAndCustomize(&Version, v3.Preference{}, func(schema *types.Schema) {
			schema.MustCustomizeField("name", func(f types.Field) types.Field {
				f.Required = true
				return f
			})
			schema.MustCustomizeField("namespaceId", func(f types.Field) types.Field {
				f.Required = false
				return f
			})
		})
}
