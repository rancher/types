package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapping/mapper"
	"github.com/rancher/types/commonmappers"
	"github.com/rancher/types/io.cattle.workload/v1/schema/mapper"
	"k8s.io/api/core/v1"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "io.cattle.workload",
		Path:    "/v1-app",
		SubContexts: map[string]bool{
			"projects": true,
		},
	}

	Schemas = commonmappers.Add(&Version, types.NewSchemas()).
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
				mapper.ResourceRequirementsMapper{},
			},
		}).
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
)
