package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapping/mapper"
	"github.com/rancher/types/apis/workload.cattle.io/v1/schema/mapper"
	"github.com/rancher/types/commonmappers"
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/staging/src/k8s.io/api/apps/v1beta2"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "workload.cattle.io",
		Path:    "/v1-workload",
		SubContexts: map[string]bool{
			"projects":   true,
			"namespaces": true,
		},
	}

	Schemas = commonmappers.Add(&Version, types.NewSchemas()).
		Init(podTypes).
		Init(namespaceTypes).
		Init(nodeTypes).
		Init(replicaSetTypes).
		Init(deploymentTypes).
		Init(statefulSetTypes)
)

func statefulSetTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta2.StatefulSetUpdateStrategy{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Enum{Field: "type",
					Values: map[string][]string{
						"RollingUpdate": {"rollingUpdate"},
						"OnDelete":      {"onDelete"},
					},
				},
				m.Move{From: "type", To: "kind"},
				m.Move{From: "rollingUpdate", To: "rollingUpdateConfig"},
			},
		}).
		AddMapperForType(&Version, v1beta2.StatefulSetSpec{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&m.Move{
					From: "replicas",
					To:   "deploy/scale",
				},
				m.Enum{Field: "podManagementPolicy",
					Values: map[string][]string{
						"OrderedReady": {"ordered"},
						"Parallel":     {"parallel"},
					},
				},
				&m.Move{
					From: "podManagementPolicy",
					To:   "deploy/strategy",
				},
				&m.Move{
					From: "revisionHistoryLimit",
					To:   "deploy/maxRevisions",
				},
				m.Drop{"selector"},
				m.SliceToMap{
					Field: "volumeClaimTemplates",
					Key:   "name",
				},
			},
		}).
		AddMapperForType(&Version, v1beta2.StatefulSet{}, m.NewObject()).
		MustImport(&Version, v1beta2.StatefulSetSpec{}, deployOverride{}).
		MustImport(&Version, v1beta2.StatefulSet{})
}

func deploymentTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta2.DeploymentStrategy{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Enum{Field: "type",
					Values: map[string][]string{
						"RollingUpdate": {"rollingUpdate"},
						"Recreate":      {"recreate"},
					},
				},
				m.Move{From: "type", To: "kind"},
				m.Move{From: "rollingUpdate", To: "rollingUpdateConfig"},
			},
		}).
		AddMapperForType(&Version, v1beta2.DeploymentSpec{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&m.Move{
					From: "replicas",
					To:   "deploy/scale",
				},
				&m.Move{
					From: "minReadySeconds",
					To:   "deploy/minReadySeconds",
				},
				&m.Move{
					From: "progressDeadlineSeconds",
					To:   "deploy/progressDeadlineSeconds",
				},
				&m.Move{
					From: "revisionHistoryLimit",
					To:   "deploy/maxRevisions",
				},
				m.Drop{"selector"},
				m.Move{From: "strategy", To: "updateStrategy"},
			},
		}).
		AddMapperForType(&Version, v1beta2.Deployment{}, m.NewObject()).
		MustImport(&Version, v1beta2.DeploymentSpec{}, deployOverride{}).
		MustImport(&Version, v1beta2.Deployment{})
}

func replicaSetTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta2.ReplicaSetSpec{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&m.Move{
					From: "replicas",
					To:   "deploy/scale",
				},
				&m.Move{
					From: "minReadySeconds",
					To:   "deploy/minReadySeconds",
				},
				m.Drop{"selector"},
			},
		}).
		AddMapperForType(&Version, v1beta2.ReplicaSet{}, m.NewObject()).
		MustImport(&Version, v1beta2.ReplicaSetSpec{}, deployOverride{}).
		MustImport(&Version, v1beta2.ReplicaSet{})
}

func nodeTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.NodeStatus{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&mapper.NodeAddressMapper{},
				&mapper.OSInfo{},
				&m.Drop{"addresses"},
				&m.Drop{"conditions"},
				&m.Drop{"daemonEndpoints"},
				&m.Drop{"images"},
				&m.Drop{"nodeInfo"},
				&m.SliceToMap{Field: "volumesAttached", Key: "devicePath"},
			},
		}).
		AddMapperForType(&Version, v1.Node{}, m.NewObject()).
		MustImport(&Version, v1.NodeStatus{}, nodeStatusOverride{}).
		MustImport(&Version, v1.Node{})
}
func namespaceTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.NamespaceStatus{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&m.Drop{"phase"},
			},
		}).
		AddMapperForType(&Version, v1.NamespaceSpec{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&m.Drop{"finalizers"},
			},
		}).
		AddMapperForType(&Version, v1.Namespace{}, m.NewObject()).
		MustImport(&Version, v1.Namespace{})
}

func podTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.Capabilities{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Move{From: "add", To: "capAdd"},
				m.Move{From: "drop", To: "capDrop"},
			},
		}).
		AddMapperForType(&Version, v1.PodSecurityContext{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Drop{"seLinuxOptions"},
				m.Move{From: "runAsUser", To: "uid"},
				m.Move{From: "supplementalGroups", To: "gids"},
				m.Move{From: "fsGroup", To: "fsgid"},
			},
		}).
		AddMapperForType(&Version, v1.SecurityContext{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&m.Embed{Field: "capabilities"},
				m.Drop{"seLinuxOptions"},
				m.Move{From: "readOnlyRootFilesystem", To: "readOnly"},
				m.Move{From: "runAsUser", To: "uid"},
			},
		}).
		AddMapperForType(&Version, v1.Container{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Move{"command", "entrypoint"},
				m.Move{"args", "command"},
				m.Move{"livenessProbe", "healthcheck"},
				m.Move{"readinessProbe", "readycheck"},
				m.Move{"imagePullPolicy", "pullPolicy"},
				mapper.EnvironmentMapper{},
				&m.Embed{Field: "securityContext"},
				&m.Embed{Field: "lifecycle"},
			},
		}).
		AddMapperForType(&Version, v1.ContainerPort{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Drop{"name"},
			},
		}).
		AddMapperForType(&Version, v1.VolumeMount{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Enum{
					Field: "mountPropagation",
					Values: map[string][]string{
						"HostToContainer": []string{"rslave"},
						"Bidirectional":   []string{"rshared", "shared"},
					},
				},
			},
		}).
		AddMapperForType(&Version, v1.Handler{}, handlerMapper).
		AddMapperForType(&Version, v1.Probe{}, handlerMapper).
		AddMapperForType(&Version, v1.PodSpec{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				m.Move{From: "restartPolicy", To: "restart"},
				m.Move{From: "imagePullSecrets", To: "pullSecrets"},
				mapper.NamespaceMapper{},
				mapper.InitContainerMapper{},
				mapper.SchedulingMapper{},
				&m.Embed{Field: "securityContext"},
				&m.SliceToMap{
					Field: "volumes",
					Key:   "name",
				},
				&m.SliceToMap{
					Field: "containers",
					Key:   "name",
				},
				&m.SliceToMap{
					Field: "hostAliases",
					Key:   "ip",
				},
			},
		}).
		AddMapperForType(&Version, v1.Pod{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				&m.Drop{"status"},
				&m.Embed{Field: "metadata"},
				&m.Embed{Field: "spec"},
			},
		}).
		AddMapperForType(&Version, v1.ResourceRequirements{}, &types.TypeMapper{
			Mappers: []types.Mapper{
				mapper.PivotMapper{Plural: true},
			},
		}).
		// Must import handlers before Container
		MustImport(&Version, v1.Handler{}, handlerOverride{}).
		MustImport(&Version, v1.Probe{}, handlerOverride{}).
		MustImport(&Version, v1.Container{}, struct {
			Scheduling      *Scheduling
			Resources       *Resources
			Environment     map[string]string
			EnvironmentFrom []EnvironmentFrom
			InitContainer   bool
		}{}).
		MustImport(&Version, v1.PodSpec{}, struct {
			Net string
			PID string
			IPC string
		}{}).
		MustImport(&Version, v1.Pod{})
}
