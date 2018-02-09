package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapper"
	"github.com/rancher/types/apis/project.cattle.io/v3"
	"github.com/rancher/types/factory"
	"github.com/rancher/types/mapper"
	"k8s.io/api/apps/v1beta2"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
)

var (
	Version = types.APIVersion{
		Version: "v3",
		Group:   "project.cattle.io",
		Path:    "/v3/projects",
		SubContexts: map[string]bool{
			"projects": true,
		},
	}

	Schemas = factory.Schemas(&Version).
		Init(namespaceTypes).
		// volume before pod types.  pod types uses volume things, so need to register mapper
		Init(volumeTypes).
		Init(ingressTypes).
		Init(secretTypes).
		Init(serviceTypes).
		Init(podTypes).
		Init(deploymentTypes).
		Init(replicationControllerTypes).
		Init(replicaSetTypes).
		Init(statefulSetTypes).
		Init(daemonSetTypes).
		Init(jobTypes).
		Init(cronJobTypes).
		Init(podTemplateSpecTypes).
		Init(workloadTypes).
		Init(configMapTypes)
)

func configMapTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.MustImport(&Version, v1.ConfigMap{})
}

type DeploymentConfig struct {
}

type StatefulSetConfig struct {
}

type ReplicaSetConfig struct {
}

type ReplicationControllerConfig struct {
}

type DaemonSetConfig struct {
}

type CronJobConfig struct {
}

type JobConfig struct {
}

type deploymentConfigOverride struct {
	Deployment DeploymentConfig
}

type statefulSetConfigOverride struct {
	StatefulSet StatefulSetConfig
}

type replicaSetConfigOverride struct {
	ReplicaSet ReplicaSetConfig
}

type replicationControllerConfigOverride struct {
	ReplicationController ReplicationControllerConfig
}

type daemonSetOverride struct {
	DaemonSet DaemonSetConfig
}

type cronJobOverride struct {
	CronJob CronJobConfig
}

type jobOverride struct {
	Job JobConfig
}

func workloadTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.MustImportAndCustomize(&Version, v3.Workload{},
		func(schema *types.Schema) {
			toInclude := []string{"deployment", "replicationController", "statefulSet",
				"daemonSet", "job", "cronJob"}
			for _, name := range toInclude {
				baseSchema := schemas.Schema(&Version, name)
				if baseSchema == nil {
					continue
				}
				for name, field := range baseSchema.ResourceFields {
					if name == "template" {
						templateSchema := schemas.Schema(&Version, field.Type)
						for name, field := range templateSchema.ResourceFields {
							schema.ResourceFields[name] = field
						}
					} else {
						schema.ResourceFields[name] = field
					}
				}

			}
		})
}

func statefulSetTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta2.StatefulSetSpec{},
			&m.Move{
				From:        "replicas",
				To:          "scale",
				DestDefined: true,
			},
			&m.Move{
				From: "updateStrategy/rollingUpdate/partition",
				To:   "statefulSet/partition",
			},
			&m.Move{
				From: "volumeClaimTemplates",
				To:   "statefulSet/volumeClaimTemplates",
			},
			&m.Move{
				From: "serviceName",
				To:   "statefulSet/serviceName",
			},
			&m.Move{
				From: "revisionHistoryLimit",
				To:   "statefulSet/revisionHistoryLimit",
			},
			&m.Move{
				From: "podManagementPolicy",
				To:   "statefulSet/podManagementPolicy",
			},
			&m.Move{
				From: "updateStrategy",
				To:   "statefulSet/updateStrategy",
			},
		).
		AddMapperForType(&Version, v1beta2.StatefulSet{},
			&m.Move{
				From: "status",
				To:   "statefulSetStatus",
			},
		).
		MustImport(&Version, v1beta2.StatefulSetSpec{}, statefulSetConfigOverride{}).
		MustImport(&Version, v1beta2.StatefulSet{}, projectOverride{})
}

func replicaSetTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta1.ReplicaSetSpec{},
			&m.Move{
				From:        "replicas",
				To:          "scale",
				DestDefined: true,
			},
			&m.Move{
				From: "minReadySeconds",
				To:   "replicaSet/minReadySeconds",
			},
		).
		AddMapperForType(&Version, v1beta1.ReplicaSet{},
			&m.Move{
				From: "status",
				To:   "replicaSetStatus",
			},
		).
		MustImport(&Version, v1beta1.ReplicaSetSpec{}, replicaSetConfigOverride{}).
		MustImportAndCustomize(&Version, v1beta1.ReplicaSet{}, func(schema *types.Schema) {
			schema.BaseType = "workload"
		}, projectOverride{})
}

func replicationControllerTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.ReplicationControllerSpec{},
			&m.Move{
				From:        "replicas",
				To:          "scale",
				DestDefined: true,
			},
			&m.Move{
				From: "minReadySeconds",
				To:   "replicationController/minReadySeconds",
			},
		).
		AddMapperForType(&Version, v1.ReplicationController{},
			&m.Move{
				From: "status",
				To:   "replicationControllerStatus",
			},
		).
		MustImport(&Version, v1.ReplicationControllerSpec{}, replicationControllerConfigOverride{}).
		MustImportAndCustomize(&Version, v1.ReplicationController{}, func(schema *types.Schema) {
			schema.BaseType = "workload"
		}, projectOverride{})
}

func daemonSetTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta2.DaemonSetSpec{},
			&m.Move{
				From: "minReadySeconds",
				To:   "daemonSet/minReadySeconds",
			},
			&m.Move{
				From: "updateStrategy",
				To:   "daemonSet/updateStrategy",
			},
			&m.Move{
				From: "revisionHistoryLimit",
				To:   "daemonSet/revisionHistoryLimit",
			},
		).
		AddMapperForType(&Version, v1beta2.DaemonSet{},
			&m.Move{
				From: "status",
				To:   "daemonSetStatus",
			},
		).
		MustImport(&Version, v1beta2.DaemonSetSpec{}, daemonSetOverride{}).
		MustImportAndCustomize(&Version, v1beta2.DaemonSet{}, func(schema *types.Schema) {
			schema.BaseType = "workload"
		}, projectOverride{})
}

func jobTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, batchv1.JobSpec{},
			&m.Move{
				From: "parallelism",
				To:   "job/parallelism",
			},
			&m.Move{
				From: "completions",
				To:   "job/completions",
			},
			&m.Move{
				From: "activeDeadlineSeconds",
				To:   "job/activeDeadlineSeconds",
			},
			&m.Move{
				From: "backoffLimit",
				To:   "job/backoffLimit",
			},
			&m.Move{
				From: "manualSelector",
				To:   "job/manualSelector",
			},
		).
		AddMapperForType(&Version, batchv1.Job{},
			&m.Move{
				From: "status",
				To:   "jobStatus",
			},
		).
		MustImport(&Version, batchv1.JobSpec{}, jobOverride{}).
		MustImportAndCustomize(&Version, batchv1.Job{}, func(schema *types.Schema) {
			schema.BaseType = "workload"
		}, projectOverride{})
}

func cronJobTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, batchv1beta1.CronJobSpec{},
			&m.Move{
				From: "schedule",
				To:   "cronJob/schedule",
			},
			&m.Move{
				From: "startingDeadlineSeconds",
				To:   "cronJob/startingDeadlineSeconds",
			},
			&m.Move{
				From: "concurrencyPolicy",
				To:   "cronJob/concurrencyPolicy",
			},
			&m.Move{
				From: "suspend",
				To:   "cronJob/suspend",
			},
			&m.Move{
				From: "successfulJobsHistoryLimit",
				To:   "cronJob/successfulJobsHistoryLimit",
			},
			&m.Move{
				From: "failedJobsHistoryLimit",
				To:   "cronJob/failedJobsHistoryLimit",
			},
			&m.Move{
				From: "jobTemplate/spec/selector",
				To:   "selector",
			},
			&m.Move{
				From: "jobTemplate/spec/template",
				To:   "template",
			},
			&m.Move{
				From: "jobTemplate/spec/job/parallelism",
				To:   "cronJob/parallelism",
			},
			&m.Move{
				From: "jobTemplate/spec/job/completions",
				To:   "cronJob/completions",
			},
			&m.Move{
				From: "jobTemplate/spec/job/activeDeadlineSeconds",
				To:   "cronJob/activeDeadlineSeconds",
			},
			&m.Move{
				From: "jobTemplate/spec/job/backoffLimit",
				To:   "cronJob/backoffLimit",
			},
			&m.Move{
				From: "jobTemplate/spec/job/manualSelector",
				To:   "cronJob/manualSelector",
			},
			&m.Drop{Field: "jobTemplate"},
		).
		AddMapperForType(&Version, batchv1beta1.CronJob{},
			&m.Move{
				From: "status",
				To:   "cronJobStatus",
			},
		).
		MustImport(&Version, batchv1beta1.CronJobSpec{}, cronJobOverride{}).
		MustImportAndCustomize(&Version, batchv1beta1.CronJob{}, func(schema *types.Schema) {
			schema.BaseType = "workload"
		}, projectOverride{})
}

func deploymentTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta2.DeploymentSpec{},
			&m.Move{
				From: "replicas",
				To:   "scale",
			},
			&m.Move{
				From: "minReadySeconds",
				To:   "deployment/minReadySeconds",
			},
			&m.Move{
				From: "strategy",
				To:   "deployment/strategy",
			},
			&m.Move{
				From: "revisionHistoryLimit",
				To:   "deployment/revisionHistoryLimit",
			},
			&m.Move{
				From: "paused",
				To:   "deployment/paused",
			},
			&m.Move{
				From: "progressDeadlineSeconds",
				To:   "deployment/progressDeadlineSeconds",
			},
		).
		AddMapperForType(&Version, v1beta2.Deployment{},
			&m.Move{
				From: "status",
				To:   "deploymentStatus",
			},
		).
		MustImport(&Version, v1beta2.DeploymentSpec{}, deploymentConfigOverride{}).
		MustImportAndCustomize(&Version, v1beta2.Deployment{}, func(schema *types.Schema) {
			schema.BaseType = "workload"
		}, projectOverride{})
}

func podTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.PodTemplateSpec{},
			&m.Embed{Field: "spec"},
		).
		AddMapperForType(&Version, v1.HTTPGetAction{},
			&m.Drop{Field: "host"},
		).
		AddMapperForType(&Version, v1.Capabilities{},
			m.Move{From: "add", To: "capAdd"},
			m.Move{From: "drop", To: "capDrop"},
		).
		AddMapperForType(&Version, v1.PodSecurityContext{},
			m.Drop{Field: "seLinuxOptions"},
			m.Move{From: "runAsUser", To: "uid"},
			m.Move{From: "supplementalGroups", To: "gids"},
			m.Move{From: "fsGroup", To: "fsgid"},
		).
		AddMapperForType(&Version, v1.SecurityContext{},
			&m.Embed{Field: "capabilities"},
			m.Drop{Field: "seLinuxOptions"},
			m.Move{From: "readOnlyRootFilesystem", To: "readOnly"},
			m.Move{From: "runAsUser", To: "uid"},
		).
		AddMapperForType(&Version, v1.Container{},
			m.Move{From: "command", To: "entrypoint"},
			m.Move{From: "args", To: "command"},
			mapper.EnvironmentMapper{},
			&m.Embed{Field: "securityContext"},
			&m.Embed{Field: "lifecycle"},
		).
		AddMapperForType(&Version, v1.ContainerPort{},
			m.Move{From: "hostIP", To: "hostIp"},
		).
		AddMapperForType(&Version, v1.Handler{}, handlerMapper).
		AddMapperForType(&Version, v1.Probe{}, handlerMapper).
		AddMapperForType(&Version, v1.PodStatus{},
			m.Move{From: "hostIP", To: "nodeIp"},
			m.Move{From: "podIP", To: "podIp"},
		).
		AddMapperForType(&Version, v1.PodSpec{},
			mapper.InitContainerMapper{},
			mapper.SchedulingMapper{},
			m.Move{From: "tolerations", To: "scheduling/tolerations", DestDefined: true},
			&m.Embed{Field: "securityContext"},
			&m.Drop{Field: "serviceAccount"},
		).
		AddMapperForType(&Version, v1.ResourceRequirements{},
			mapper.PivotMapper{Plural: true},
		).
		MustImport(&Version, v3.PublicEndpoint{}).
		AddMapperForType(&Version, v1.Pod{},
			&m.AnnotationField{Field: "description"},
			&m.AnnotationField{Field: "publicEndpoints", List: true},
		).
		// Must import handlers before Container
		MustImport(&Version, v1.Capabilities{}, struct {
			Add  []string `norman:"type=array[enum],options=AUDIT_CONTROL|AUDIT_WRITE|BLOCK_SUSPEND|CHOWN|DAC_OVERRIDE|DAC_READ_SEARCH|FOWNER|FSETID|IPC_LOCK|IPC_OWNER|KILL|LEASE|LINUX_IMMUTABLE|MAC_ADMIN|MAC_OVERRIDE|MKNOD|NET_ADMIN|NET_BIND_SERVICE|NET_BROADCAST|NET_RAW|SETFCAP|SETGID|SETPCAP|SETUID|SYSLOG|SYS_ADMIN|SYS_BOOT|SYS_CHROOT|SYS_MODULE|SYS_NICE|SYS_PACCT|SYS_PTRACE|SYS_RAWIO|SYS_RESOURCE|SYS_TIME|SYS_TTY_CONFIG|WAKE_ALARM"`
			Drop []string `norman:"type=array[enum],options=AUDIT_CONTROL|AUDIT_WRITE|BLOCK_SUSPEND|CHOWN|DAC_OVERRIDE|DAC_READ_SEARCH|FOWNER|FSETID|IPC_LOCK|IPC_OWNER|KILL|LEASE|LINUX_IMMUTABLE|MAC_ADMIN|MAC_OVERRIDE|MKNOD|NET_ADMIN|NET_BIND_SERVICE|NET_BROADCAST|NET_RAW|SETFCAP|SETGID|SETPCAP|SETUID|SYSLOG|SYS_ADMIN|SYS_BOOT|SYS_CHROOT|SYS_MODULE|SYS_NICE|SYS_PACCT|SYS_PTRACE|SYS_RAWIO|SYS_RESOURCE|SYS_TIME|SYS_TTY_CONFIG|WAKE_ALARM"`
		}{}).
		MustImport(&Version, v1.Handler{}, handlerOverride{}).
		MustImport(&Version, v1.Probe{}, handlerOverride{}).
		MustImport(&Version, v1.Container{}, struct {
			Resources       *Resources
			Environment     map[string]string
			EnvironmentFrom []EnvironmentFrom
			InitContainer   bool
		}{}).
		MustImport(&Version, v1.PodSpec{}, struct {
			Scheduling *Scheduling
			NodeName   string `norman:"type=reference[node]"`
		}{}).
		MustImport(&Version, v1.Pod{}, projectOverride{}, struct {
			Description     string `json:"description"`
			WorkloadID      string `norman:"type=reference[workload]"`
			PublicEndpoints string `json:"publicEndpoints" norman:"type=array[publicEndpoint],nocreate,noupdate"`
		}{})
}

func serviceTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		Init(addServiceType).
		Init(addDNSRecord)
}

func addServiceType(schemas *types.Schemas) *types.Schemas {
	return schemas.AddSchema(*factory.Schemas(&Version).
		Init(addServiceOrDNSRecord(false)).
		Schema(&Version, "service"))
}

func addDNSRecord(schemas *types.Schemas) *types.Schemas {
	return schemas.
		Init(addServiceOrDNSRecord(true))
}

func addServiceOrDNSRecord(dns bool) types.SchemasInitFunc {
	return func(schemas *types.Schemas) *types.Schemas {
		if dns {
			schemas = schemas.
				TypeName("dnsRecord", v1.Service{})
		}

		schemas = schemas.
			AddMapperForType(&Version, v1.ServiceSpec{},
				&m.Move{From: "externalName", To: "hostname"},
				&ServiceSpecMapper{},
				&m.Move{From: "type", To: "serviceKind"},
				&m.SetValue{
					Field: "clusterIP",
					IfEq:  "None",
					Value: nil,
				},
				&m.Move{From: "clusterIP", To: "clusterIp"},
			).
			AddMapperForType(&Version, v1.Service{},
				&m.Drop{Field: "status"},
				&m.LabelField{Field: "workloadId"},
				&m.AnnotationField{Field: "description"},
				&m.AnnotationField{Field: "ipAddresses", List: true},
				&m.AnnotationField{Field: "targetWorkloadIds", List: true},
				&m.AnnotationField{Field: "targetDnsRecordIds", List: true},
				&m.AnnotationField{Field: "publicEndpoints", List: true},
				&m.Move{From: "serviceKind", To: "kind"},
			)

		if dns {
			schemas = schemas.
				AddMapperForType(&Version, v1.Service{},
					&m.Drop{Field: "kind"},
					&m.Drop{Field: "externalIPs"},
					&m.Drop{Field: "externalTrafficPolicy"},
					&m.Drop{Field: "healthCheckNodePort"},
					&m.Drop{Field: "loadBalancerIP"},
					&m.Drop{Field: "loadBalancerSourceRanges"},
					&m.Drop{Field: "ports"},
					&m.Drop{Field: "publishNotReadyAddresses"},
					&m.Drop{Field: "sessionAffinity"},
					&m.Drop{Field: "sessionAffinityConfig"},
				)
		}

		return schemas.MustImportAndCustomize(&Version, v1.Service{}, func(schema *types.Schema) {
			if dns {
				schema.CodeName = "DNSRecord"
				schema.MustCustomizeField("clusterIp", func(f types.Field) types.Field {
					f.Create = false
					f.Update = false
					return f
				})
			}
		}, projectOverride{}, struct {
			Description        string   `json:"description"`
			IPAddresses        []string `json:"ipAddresses"`
			WorkloadID         string   `json:"workloadId" norman:"type=reference[workload],nocreate,noupdate"`
			TargetWorkloadIDs  string   `json:"targetWorkloadIds" norman:"type=array[reference[workload]]"`
			TargetDNSRecordIDs string   `json:"targetDnsRecordIds" norman:"type=array[reference[dnsRecord]]"`
			PublicEndpoints    string   `json:"publicEndpoints" norman:"type=array[publicEndpoint],nocreate,noupdate"`
		}{})
	}
}

func ingressTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1beta1.HTTPIngressRuleValue{},
			&m.SliceToMap{Field: "paths", Key: "path"},
		).
		AddMapperForType(&Version, v1beta1.HTTPIngressPath{},
			&m.Embed{Field: "backend"},
		).
		AddMapperForType(&Version, v1beta1.IngressRule{},
			&m.Embed{Field: "http"},
		).
		AddMapperForType(&Version, v1beta1.Ingress{},
			&m.AnnotationField{Field: "description"},
			&m.Move{From: "backend", To: "defaultBackend"},
		).
		AddMapperForType(&Version, v1beta1.IngressTLS{},
			&m.Move{From: "secretName", To: "certificateName"},
		).
		AddMapperForType(&Version, v1beta1.IngressBackend{},
			&m.Move{From: "servicePort", To: "targetPort"},
		).
		MustImport(&Version, v1beta1.IngressBackend{}, struct {
			WorkloadIDs string `json:"workloadIds" norman:"type=array[reference[workload]]"`
			ServiceName string `norman:"type=reference[service]"`
		}{}).
		MustImportAndCustomize(&Version, v1beta1.IngressRule{}, func(schema *types.Schema) {
			schema.MustCustomizeField("paths", func(f types.Field) types.Field {
				f.Type = "map[ingressBackend]"
				return f
			})
		}).
		MustImport(&Version, v1beta1.IngressTLS{}, struct {
			SecretName string `norman:"type=reference[certificate]"`
		}{}).
		MustImport(&Version, v1beta1.Ingress{}, projectOverride{}, struct {
			Description string `json:"description"`
		}{})
}

func volumeTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v1.PersistentVolumeClaim{}, projectOverride{})
}

func podTemplateSpecTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v1.PodTemplateSpec{})
}

func namespaceTypes(schemas *types.Schemas) *types.Schemas {
	return NamespaceTypes(&Version, schemas)
}

func NamespaceTypes(version *types.APIVersion, schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(version, v1.NamespaceStatus{},
			&m.Drop{Field: "phase"},
		).
		AddMapperForType(version, v1.NamespaceSpec{},
			&m.Drop{Field: "finalizers"},
		).
		AddMapperForType(version, v1.Namespace{},
			&m.AnnotationField{Field: "description"},
			&m.AnnotationField{Field: "projectId"},
			&m.AnnotationField{Field: "externalId"},
			&m.AnnotationField{Field: "templates", Object: true},
			&m.AnnotationField{Field: "prune"},
			&m.AnnotationField{Field: "answers", Object: true},
		).
		MustImport(version, v1.Namespace{}, struct {
			Description string                 `json:"description"`
			ProjectID   string                 `norman:"type=reference[/v3/schemas/project]"`
			Templates   map[string]string      `json:"templates"`
			Answers     map[string]interface{} `json:"answers"`
			Prune       bool                   `json:"prune"`
			ExternalID  string                 `json:"externalId"`
			Tags        []string               `json:"tags"`
		}{})
}
