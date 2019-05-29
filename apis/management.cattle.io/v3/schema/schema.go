package schema

import (
	"net/http"

	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapper"
	v3 "github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/factory"
	"github.com/rancher/types/mapper"
	v1 "k8s.io/api/core/v1"
)

var (
	Version = types.APIVersion{
		Version: "v3",
		Group:   "management.cattle.io",
		Path:    "/v3",
	}

	Schemas = factory.Schemas(&Version).
		Init(nativeNodeTypes).
		Init(nodeTypes).
		Init(authzTypes).
		Init(clusterTypes).
		Init(catalogTypes).
		Init(authnTypes).
		Init(tokens).
		Init(schemaTypes).
		Init(userTypes).
		Init(projectNetworkPolicyTypes).
		Init(logTypes).
		Init(globalTypes).
		Init(rkeTypes).
		Init(alertTypes).
		Init(composeType).
		Init(ClusterRandomizerType).
		Init(projectCatalogTypes).
		Init(clusterCatalogTypes).
		Init(multiClusterAppTypes).
		Init(globalDNSTypes).
		Init(kontainerTypes).
		Init(etcdBackupTypes).
		Init(monitorTypes).
		Init(credTypes).
		Init(mgmtSecretTypes)

	TokenSchemas = factory.Schemas(&Version).
			Init(tokens)
)

func rkeTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.AddMapperForType(&Version, v3.BaseService{}, m.Drop{Field: "image"})
}

func schemaTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.DynamicSchema{})
}

func credTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.CloudCredential{},
			&mapper.CredentialMapper{},
			&m.Drop{Field: "namespaceId"}).
		MustImport(&Version, v3.CloudCredential{})
}

func mgmtSecretTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.MustImportAndCustomize(&Version, v1.Secret{}, func(schema *types.Schema) {
		schema.ID = "managementSecret"
		schema.PluralName = "managementSecrets"
		schema.CodeName = "ManagementSecret"
		schema.CodeNamePlural = "ManagementSecrets"
	})
}

func catalogTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.Catalog{},
			&m.Move{From: "catalogKind", To: "kind"},
			&m.Embed{Field: "status"},
			&m.Drop{Field: "helmVersionCommits"},
		).
		MustImportAndCustomize(&Version, v3.Catalog{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"refresh": {},
			}
			schema.CollectionActions = map[string]types.Action{
				"refresh": {},
			}
		}).
		AddMapperForType(&Version, v3.Template{},
			m.DisplayName{},
		).
		MustImport(&Version, v3.Template{}, struct {
			VersionLinks map[string]string
		}{}).
		AddMapperForType(&Version, v3.CatalogTemplate{},
			m.DisplayName{},
			m.Drop{Field: "namespaceId"},
		).
		MustImport(&Version, v3.CatalogTemplate{}, struct {
			VersionLinks map[string]string
		}{}).
		AddMapperForType(&Version, v3.CatalogTemplateVersion{},
			m.Drop{Field: "namespaceId"},
		).
		MustImport(&Version, v3.CatalogTemplateVersion{}).
		MustImport(&Version, v3.TemplateVersion{}).
		MustImport(&Version, v3.TemplateContent{})
}

func nativeNodeTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		TypeName("internalNodeStatus", v1.NodeStatus{}).
		TypeName("internalNodeSpec", v1.NodeSpec{}).
		AddMapperForType(&Version, v1.NodeStatus{},
			&mapper.NodeAddressMapper{},
			&mapper.OSInfo{},
			&m.Drop{Field: "addresses"},
			&m.Drop{Field: "daemonEndpoints"},
			&m.Drop{Field: "images"},
			&m.Drop{Field: "nodeInfo"},
			&m.Move{From: "conditions", To: "nodeConditions"},
			&m.Drop{Field: "phase"},
			&m.SliceToMap{Field: "volumesAttached", Key: "devicePath"},
		).
		AddMapperForType(&Version, v1.NodeSpec{},
			&m.Drop{Field: "externalID"},
			&m.Drop{Field: "configSource"},
			&m.Move{From: "providerID", To: "providerId"},
			&m.Move{From: "podCIDR", To: "podCidr"},
			m.Access{Fields: map[string]string{
				"podCidr":       "r",
				"providerId":    "r",
				"taints":        "ru",
				"unschedulable": "ru",
			}}).
		AddMapperForType(&Version, v1.Node{},
			&mapper.NodeAddressAnnotationMapper{}).
		MustImportAndCustomize(&Version, v1.NodeSpec{}, func(schema *types.Schema) {
			schema.CodeName = "InternalNodeSpec"
			schema.CodeNamePlural = "InternalNodeSpecs"
		}).
		MustImportAndCustomize(&Version, v1.NodeStatus{}, func(schema *types.Schema) {
			schema.CodeName = "InternalNodeStatus"
			schema.CodeNamePlural = "InternalNodeStatuses"
		}, struct {
			IPAddress         string
			ExternalIPAddress string `json:"externalIpAddress,omitempty"`
			Hostname          string
			Info              NodeInfo
		}{})
}

func clusterTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.Cluster{},
			&m.Embed{Field: "status"},
			mapper.NewDropFromSchema("genericEngineConfig"),
			mapper.NewDropFromSchema("googleKubernetesEngineConfig"),
			mapper.NewDropFromSchema("azureKubernetesServiceConfig"),
			mapper.NewDropFromSchema("amazonElasticContainerServiceConfig"),
			m.DisplayName{},
		).
		AddMapperForType(&Version, v3.ClusterStatus{},
			m.Drop{Field: "serviceAccountToken"},
		).
		AddMapperForType(&Version, v3.ClusterRegistrationToken{},
			&m.Embed{Field: "status"},
		).
		AddMapperForType(&Version, v3.RancherKubernetesEngineConfig{},
			m.Drop{Field: "systemImages"},
		).
		MustImport(&Version, v3.Cluster{}).
		MustImport(&Version, v3.ClusterRegistrationToken{}).
		MustImport(&Version, v3.GenerateKubeConfigOutput{}).
		MustImport(&Version, v3.ImportClusterYamlInput{}).
		MustImport(&Version, v3.RotateCertificateInput{}).
		MustImport(&Version, v3.RotateCertificateOutput{}).
		MustImport(&Version, v3.ImportYamlOutput{}).
		MustImport(&Version, v3.ExportOutput{}).
		MustImport(&Version, v3.MonitoringInput{}).
		MustImport(&Version, v3.MonitoringOutput{}).
		MustImport(&Version, v3.RestoreFromEtcdBackupInput{}).
		MustImportAndCustomize(&Version, v3.ETCDService{}, func(schema *types.Schema) {
			schema.MustCustomizeField("extraArgs", func(field types.Field) types.Field {
				field.Default = map[string]interface{}{
					"election-timeout":   "5000",
					"heartbeat-interval": "500"}
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.Cluster{}, func(schema *types.Schema) {
			schema.MustCustomizeField("name", func(field types.Field) types.Field {
				field.Type = "dnsLabel"
				field.Nullable = true
				field.Required = false
				return field
			})
			schema.ResourceActions[v3.ClusterActionGenerateKubeconfig] = types.Action{
				Output: "generateKubeConfigOutput",
			}
			schema.ResourceActions[v3.ClusterActionImportYaml] = types.Action{
				Input:  "importClusterYamlInput",
				Output: "importYamlOutput",
			}
			schema.ResourceActions[v3.ClusterActionExportYaml] = types.Action{
				Output: "exportOutput",
			}
			schema.ResourceActions[v3.ClusterActionEnableMonitoring] = types.Action{
				Input: "monitoringInput",
			}
			schema.ResourceActions[v3.ClusterActionDisableMonitoring] = types.Action{}
			schema.ResourceActions[v3.ClusterActionViewMonitoring] = types.Action{
				Output: "monitoringOutput",
			}
			schema.ResourceActions[v3.ClusterActionEditMonitoring] = types.Action{
				Input: "monitoringInput",
			}
			schema.ResourceActions[v3.ClusterActionBackupEtcd] = types.Action{}
			schema.ResourceActions[v3.ClusterActionRestoreFromEtcdBackup] = types.Action{
				Input: "restoreFromEtcdBackupInput",
			}
			schema.ResourceActions[v3.ClusterActionRotateCertificates] = types.Action{
				Input:  "rotateCertificateInput",
				Output: "rotateCertificateOutput",
			}
		})
}

func authzTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.ProjectStatus{}).
		AddMapperForType(&Version, v3.Project{},
			m.DisplayName{},
			&m.Embed{Field: "status"},
		).
		AddMapperForType(&Version, v3.GlobalRole{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.RoleTemplate{}, m.DisplayName{}).
		AddMapperForType(&Version,
			v3.PodSecurityPolicyTemplateProjectBinding{},
			&mapper.NamespaceIDMapper{}).
		AddMapperForType(&Version, v3.ProjectRoleTemplateBinding{},
			&mapper.NamespaceIDMapper{},
		).
		MustImport(&Version, v3.SetPodSecurityPolicyTemplateInput{}).
		MustImport(&Version, v3.ImportYamlOutput{}).
		MustImport(&Version, v3.MonitoringInput{}).
		MustImport(&Version, v3.MonitoringOutput{}).
		MustImportAndCustomize(&Version, v3.Project{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"setpodsecuritypolicytemplate": {
					Input:  "setPodSecurityPolicyTemplateInput",
					Output: "project",
				},
				"exportYaml": {},
				"enableMonitoring": {
					Input: "monitoringInput",
				},
				"disableMonitoring": {},
				"viewMonitoring": {
					Output: "monitoringOutput",
				},
				"editMonitoring": {
					Input: "monitoringInput",
				},
			}
		}).
		MustImportAndCustomize(&Version, v3.GlobalRole{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet}
			schema.ResourceMethods = []string{http.MethodGet, http.MethodPut}
		}).
		MustImport(&Version, v3.GlobalRoleBinding{}).
		MustImport(&Version, v3.RoleTemplate{}).
		MustImport(&Version, v3.PodSecurityPolicyTemplate{}).
		MustImportAndCustomize(&Version, v3.PodSecurityPolicyTemplateProjectBinding{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet, http.MethodPost}
			schema.ResourceMethods = []string{}
		}).
		MustImport(&Version, v3.ClusterRoleTemplateBinding{}).
		MustImport(&Version, v3.ProjectRoleTemplateBinding{}).
		MustImport(&Version, v3.GlobalRoleBinding{})
}

func nodeTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.NodeSpec{}, &m.Embed{Field: "internalNodeSpec"}).
		AddMapperForType(&Version, v3.NodeStatus{},
			&m.Drop{Field: "nodeTemplateSpec"},
			&m.Embed{Field: "internalNodeStatus"},
			&m.Drop{Field: "config"},
			&m.SliceMerge{From: []string{"conditions", "nodeConditions"}, To: "conditions"}).
		AddMapperForType(&Version, v3.Node{},
			&m.Embed{Field: "status"},
			&m.Move{From: "rkeNode/user", To: "sshUser"},
			&m.ReadOnly{Field: "sshUser"},
			&m.Drop{Field: "rkeNode"},
			&m.Drop{Field: "labels"},
			&m.Drop{Field: "annotations"},
			&m.Move{From: "nodeLabels", To: "labels"},
			&m.Move{From: "nodeAnnotations", To: "annotations"},
			&m.Drop{Field: "desiredNodeLabels"},
			&m.Drop{Field: "desiredNodeAnnotations"},
			&m.Drop{Field: "currentNodeLabels"},
			&m.Drop{Field: "currentNodeAnnotations"},
			&m.Drop{Field: "desiredNodeUnschedulable"},
			&m.Drop{Field: "nodeDrainInput"},
			&m.AnnotationField{Field: "publicEndpoints", List: true},
			m.Copy{From: "namespaceId", To: "clusterName"},
			m.DisplayName{}).
		AddMapperForType(&Version, v3.NodeDriver{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.NodeTemplate{}, m.DisplayName{}).
		MustImport(&Version, v3.PublicEndpoint{}).
		MustImport(&Version, v3.NodePool{}).
		MustImport(&Version, v3.NodeDrainInput{}).
		MustImportAndCustomize(&Version, v3.Node{}, func(schema *types.Schema) {
			labelField := schema.ResourceFields["labels"]
			labelField.Create = true
			labelField.Update = true
			schema.ResourceFields["labels"] = labelField
			unschedulable := schema.ResourceFields["unschedulable"]
			unschedulable.Create = false
			unschedulable.Update = false
			schema.ResourceFields["unschedulable"] = unschedulable
			clusterField := schema.ResourceFields["clusterId"]
			clusterField.Type = "reference[cluster]"
			schema.ResourceFields["clusterId"] = clusterField
			schema.ResourceActions["cordon"] = types.Action{}
			schema.ResourceActions["uncordon"] = types.Action{}
			schema.ResourceActions["stopDrain"] = types.Action{}
			schema.ResourceActions["drain"] = types.Action{
				Input: "nodeDrainInput",
			}
		}, struct {
			PublicEndpoints string `json:"publicEndpoints" norman:"type=array[publicEndpoint],nocreate,noupdate"`
		}{}).
		MustImportAndCustomize(&Version, v3.NodeDriver{}, func(schema *types.Schema) {
			schema.ResourceActions["activate"] = types.Action{
				Output: "nodeDriver",
			}
			schema.ResourceActions["deactivate"] = types.Action{
				Output: "nodeDriver",
			}
		}).
		MustImportAndCustomize(&Version, v3.NodeTemplate{}, func(schema *types.Schema) {
			delete(schema.ResourceFields, "namespaceId")
		})
}

func tokens(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImportAndCustomize(&Version, v3.Token{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"logout": {},
			}
		})
}

func authnTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.User{}, m.DisplayName{},
			&m.Embed{Field: "status"}).
		AddMapperForType(&Version, v3.Group{}, m.DisplayName{}).
		MustImport(&Version, v3.Group{}).
		MustImport(&Version, v3.GroupMember{}).
		AddMapperForType(&Version, v3.Principal{}, m.DisplayName{}).
		MustImportAndCustomize(&Version, v3.Principal{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet}
			schema.ResourceMethods = []string{http.MethodGet}
			schema.CollectionActions = map[string]types.Action{
				"search": {
					Input:  "searchPrincipalsInput",
					Output: "collection",
				},
			}
		}).
		MustImport(&Version, v3.SearchPrincipalsInput{}).
		MustImport(&Version, v3.ChangePasswordInput{}).
		MustImport(&Version, v3.SetPasswordInput{}).
		MustImportAndCustomize(&Version, v3.User{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"setpassword": {
					Input:  "setPasswordInput",
					Output: "user",
				},
				"refreshauthprovideraccess": {},
			}
			schema.CollectionActions = map[string]types.Action{
				"changepassword": {
					Input: "changePasswordInput",
				},
				"refreshauthprovideraccess": {},
			}
		}).
		MustImportAndCustomize(&Version, v3.AuthConfig{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet}
		}).
		// Local Config
		MustImportAndCustomize(&Version, v3.LocalConfig{}, func(schema *types.Schema) {
			schema.BaseType = "authConfig"
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		//Github Config
		MustImportAndCustomize(&Version, v3.GithubConfig{}, func(schema *types.Schema) {
			schema.BaseType = "authConfig"
			schema.ResourceActions = map[string]types.Action{
				"disable": {},
				"configureTest": {
					Input:  "githubConfig",
					Output: "githubConfigTestOutput",
				},
				"testAndApply": {
					Input: "githubConfigApplyInput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet, http.MethodPut}
		}).
		MustImport(&Version, v3.GithubConfigTestOutput{}).
		MustImport(&Version, v3.GithubConfigApplyInput{}).
		//AzureAD Config
		MustImportAndCustomize(&Version, v3.AzureADConfig{}, func(schema *types.Schema) {
			schema.BaseType = "authConfig"
			schema.ResourceActions = map[string]types.Action{
				"disable": {},
				"configureTest": {
					Input:  "azureADConfig",
					Output: "azureADConfigTestOutput",
				},
				"testAndApply": {
					Input: "azureADConfigApplyInput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet, http.MethodPut}
		}).
		MustImport(&Version, v3.AzureADConfigTestOutput{}).
		MustImport(&Version, v3.AzureADConfigApplyInput{}).
		// Active Directory Config
		MustImportAndCustomize(&Version, v3.ActiveDirectoryConfig{}, func(schema *types.Schema) {
			schema.BaseType = "authConfig"
			schema.ResourceActions = map[string]types.Action{
				"disable": {},
				"testAndApply": {
					Input: "activeDirectoryTestAndApplyInput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet, http.MethodPut}
		}).
		MustImport(&Version, v3.ActiveDirectoryTestAndApplyInput{}).
		// OpenLdap Config
		MustImportAndCustomize(&Version, v3.OpenLdapConfig{}, func(schema *types.Schema) {
			schema.BaseType = "authConfig"
			schema.ResourceActions = map[string]types.Action{
				"disable": {},
				"testAndApply": {
					Input: "openLdapTestAndApplyInput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet, http.MethodPut}
		}).
		MustImport(&Version, v3.OpenLdapTestAndApplyInput{}).
		// FreeIpa Config
		AddMapperForType(&Version, v3.FreeIpaConfig{}, m.Drop{Field: "nestedGroupMembershipEnabled"}).
		MustImportAndCustomize(&Version, v3.FreeIpaConfig{}, func(schema *types.Schema) {
			schema.BaseType = "authConfig"
			schema.ResourceActions = map[string]types.Action{
				"disable": {},
				"testAndApply": {
					Input: "freeIpaTestAndApplyInput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet, http.MethodPut}
			schema.MustCustomizeField("groupObjectClass", func(f types.Field) types.Field {
				f.Default = "groupofnames"
				return f
			})
			schema.MustCustomizeField("userNameAttribute", func(f types.Field) types.Field {
				f.Default = "givenName"
				return f
			})
			schema.MustCustomizeField("userObjectClass", func(f types.Field) types.Field {
				f.Default = "inetorgperson"
				return f
			})
			schema.MustCustomizeField("groupDNAttribute", func(f types.Field) types.Field {
				f.Default = "entrydn"
				return f
			})
			schema.MustCustomizeField("groupMemberUserAttribute", func(f types.Field) types.Field {
				f.Default = "entrydn"
				return f
			})
		}).
		MustImport(&Version, v3.FreeIpaTestAndApplyInput{}).
		// Saml Config
		// Ping-Saml Config
		// KeyCloak-Saml Configs
		MustImportAndCustomize(&Version, v3.PingConfig{}, configSchema).
		MustImportAndCustomize(&Version, v3.ADFSConfig{}, configSchema).
		MustImportAndCustomize(&Version, v3.KeyCloakConfig{}, configSchema).
		MustImportAndCustomize(&Version, v3.OKTAConfig{}, configSchema).
		MustImport(&Version, v3.SamlConfigTestInput{}).
		MustImport(&Version, v3.SamlConfigTestOutput{})
}

func configSchema(schema *types.Schema) {
	schema.BaseType = "authConfig"
	schema.ResourceActions = map[string]types.Action{
		"disable": {},
		"testAndEnable": {
			Input:  "samlConfigTestInput",
			Output: "samlConfigTestOutput",
		},
	}
	schema.CollectionMethods = []string{}
	schema.ResourceMethods = []string{http.MethodGet, http.MethodPut}
}

func userTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		MustImportAndCustomize(&Version, v3.Preference{}, func(schema *types.Schema) {
			schema.MustCustomizeField("name", func(f types.Field) types.Field {
				f.Required = true
				return f
			})
			schema.MustCustomizeField("namespaceId", func(f types.Field) types.Field {
				f.Required = false
				return f
			})
		}).
		MustImportAndCustomize(&Version, v3.UserAttribute{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{}
		})
}

func projectNetworkPolicyTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		MustImportAndCustomize(&Version, v3.ProjectNetworkPolicy{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet}
			schema.ResourceMethods = []string{http.MethodGet}
		})
}

func logTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		AddMapperForType(&Version, v3.ClusterLogging{},
			&m.Embed{Field: "status"},
			m.DisplayName{}).
		AddMapperForType(&Version, v3.ProjectLogging{},
			m.DisplayName{}).
		MustImport(&Version, v3.ClusterTestInput{}).
		MustImport(&Version, v3.ProjectTestInput{}).
		MustImportAndCustomize(&Version, v3.ClusterLogging{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"test": {
					Input: "clusterTestInput",
				},
				"dryRun": {
					Input: "clusterTestInput",
				},
			}
		}).
		MustImportAndCustomize(&Version, v3.ProjectLogging{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"test": {
					Input: "projectTestInput",
				},
				"dryRun": {
					Input: "projectTestInput",
				},
			}
		})
}

func globalTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		AddMapperForType(&Version, v3.ListenConfig{},
			m.DisplayName{},
			m.Drop{Field: "caKey"},
			m.Drop{Field: "caCert"},
		).
		MustImport(&Version, v3.ListenConfig{}).
		MustImportAndCustomize(&Version, v3.Setting{}, func(schema *types.Schema) {
			schema.MustCustomizeField("name", func(f types.Field) types.Field {
				f.Required = true
				return f
			})
		})
}

func alertTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		AddMapperForType(&Version, v3.Notifier{},
			m.DisplayName{}).
		MustImport(&Version, v3.ClusterAlert{}).
		MustImport(&Version, v3.ProjectAlert{}).
		MustImport(&Version, v3.Notification{}).
		MustImportAndCustomize(&Version, v3.Notifier{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"send": {
					Input: "notification",
				},
			}
			schema.ResourceActions = map[string]types.Action{
				"send": {
					Input: "notification",
				},
			}
		}).
		MustImport(&Version, v3.AlertStatus{}).
		AddMapperForType(&Version, v3.ClusterAlertGroup{},
			&m.Embed{Field: "status"},
			m.DisplayName{}).
		AddMapperForType(&Version, v3.ProjectAlertGroup{},
			&m.Embed{Field: "status"},
			m.DisplayName{}).
		AddMapperForType(&Version, v3.ClusterAlertRule{},
			&m.Embed{Field: "status"},
			m.DisplayName{}).
		AddMapperForType(&Version, v3.ProjectAlertRule{},
			&m.Embed{Field: "status"},
			m.DisplayName{}).
		MustImport(&Version, v3.ClusterAlertGroup{}).
		MustImport(&Version, v3.ProjectAlertGroup{}).
		MustImportAndCustomize(&Version, v3.ClusterAlertRule{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"activate":   {},
				"deactivate": {},
				"mute":       {},
				"unmute":     {},
			}
		}).
		MustImportAndCustomize(&Version, v3.ProjectAlertRule{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"activate":   {},
				"deactivate": {},
				"mute":       {},
				"unmute":     {},
			}
		})

}

func composeType(schemas *types.Schemas) *types.Schemas {
	return schemas.MustImport(&Version, v3.ComposeConfig{})
}

func ClusterRandomizerType(schemas *types.Schemas) *types.Schemas {
	return schemas.MustImport(&Version, v3.ClusterRandomizer{})
}

func projectCatalogTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.ProjectCatalog{},
			&m.Move{From: "catalogKind", To: "kind"},
			&m.Embed{Field: "status"},
			&m.Drop{Field: "helmVersionCommits"},
			&mapper.NamespaceIDMapper{}).
		MustImportAndCustomize(&Version, v3.ProjectCatalog{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"refresh": {},
			}
			schema.CollectionActions = map[string]types.Action{
				"refresh": {},
			}
		})
}

func clusterCatalogTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.ClusterCatalog{},
			&m.Move{From: "catalogKind", To: "kind"},
			&m.Embed{Field: "status"},
			&m.Drop{Field: "helmVersionCommits"},
			&mapper.NamespaceIDMapper{}).
		MustImportAndCustomize(&Version, v3.ClusterCatalog{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"refresh": {},
			}
			schema.CollectionActions = map[string]types.Action{
				"refresh": {},
			}
		})
}

func multiClusterAppTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.MultiClusterApp{}, m.Drop{Field: "namespaceId"}).
		AddMapperForType(&Version, v3.MultiClusterAppRevision{}, m.Drop{Field: "namespaceId"}).
		AddMapperForType(&Version, v3.Member{}, m.Drop{Field: "userName"}, m.Drop{Field: "displayName"}).
		MustImport(&Version, v3.MultiClusterApp{}).
		MustImport(&Version, v3.Target{}).
		MustImport(&Version, v3.UpgradeStrategy{}).
		MustImport(&Version, v3.MultiClusterAppRollbackInput{}).
		MustImport(&Version, v3.MultiClusterAppRevision{}).
		MustImport(&Version, v3.UpdateMultiClusterAppTargetsInput{}).
		MustImportAndCustomize(&Version, v3.MultiClusterApp{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"rollback": {
					Input: "multiClusterAppRollbackInput",
				},
				"addProjects": {
					Input: "updateMultiClusterAppTargetsInput",
				},
				"removeProjects": {
					Input: "updateMultiClusterAppTargetsInput",
				},
			}
		})
}

func globalDNSTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		TypeName("globalDns", v3.GlobalDNS{}).
		TypeName("globalDnsProvider", v3.GlobalDNSProvider{}).
		TypeName("globalDnsSpec", v3.GlobalDNSSpec{}).
		TypeName("globalDnsStatus", v3.GlobalDNSStatus{}).
		TypeName("globalDnsProviderSpec", v3.GlobalDNSProviderSpec{}).
		MustImport(&Version, v3.UpdateGlobalDNSTargetsInput{}).
		AddMapperForType(&Version, v3.GlobalDNS{}, m.Drop{Field: "namespaceId"}).
		AddMapperForType(&Version, v3.GlobalDNSProvider{}, m.Drop{Field: "namespaceId"}).
		MustImportAndCustomize(&Version, v3.GlobalDNS{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"addProjects": {
					Input: "updateGlobalDNSTargetsInput",
				},
				"removeProjects": {
					Input: "updateGlobalDNSTargetsInput",
				},
			}
		}).
		MustImportAndCustomize(&Version, v3.GlobalDNSProvider{}, func(schema *types.Schema) {
		})
}

func kontainerTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.KontainerDriver{},
			&m.Embed{Field: "status"},
			m.DisplayName{},
		).
		MustImportAndCustomize(&Version, v3.KontainerDriver{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"activate":   {},
				"deactivate": {},
			}
		})
}

func monitorTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.QueryGraphInput{}).
		MustImport(&Version, v3.QueryClusterGraphOutput{}).
		MustImport(&Version, v3.QueryProjectGraphOutput{}).
		MustImport(&Version, v3.QueryClusterMetricInput{}).
		MustImport(&Version, v3.QueryProjectMetricInput{}).
		MustImport(&Version, v3.QueryMetricOutput{}).
		MustImport(&Version, v3.ClusterMetricNamesInput{}).
		MustImport(&Version, v3.ProjectMetricNamesInput{}).
		MustImport(&Version, v3.MetricNamesOutput{}).
		MustImport(&Version, v3.TimeSeries{}).
		MustImportAndCustomize(&Version, v3.MonitorMetric{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"querycluster": {
					Input:  "queryClusterMetricInput",
					Output: "queryMetricOutput",
				},
				"listclustermetricname": {
					Input:  "clusterMetricNamesInput",
					Output: "metricNamesOutput",
				},
				"queryproject": {
					Input:  "queryProjectMetricInput",
					Output: "queryMetricOutput",
				},
				"listprojectmetricname": {
					Input:  "projectMetricNamesInput",
					Output: "metricNamesOutput",
				},
			}
		}).
		MustImportAndCustomize(&Version, v3.ClusterMonitorGraph{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"query": {
					Input:  "queryGraphInput",
					Output: "queryClusterGraphOutput",
				},
			}
		}).
		MustImportAndCustomize(&Version, v3.ProjectMonitorGraph{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"query": {
					Input:  "queryGraphInput",
					Output: "queryProjectGraphOutput",
				},
			}
		})

}

func etcdBackupTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.MustImport(&Version, v3.EtcdBackup{})
}
