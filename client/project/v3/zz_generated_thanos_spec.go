package client

const (
	ThanosSpecType                     = "thanosSpec"
	ThanosSpecFieldBaseImage           = "baseImage"
	ThanosSpecFieldGRPCServerTLSConfig = "grpcServerTlsConfig"
	ThanosSpecFieldImage               = "image"
	ThanosSpecFieldListenLocal         = "listenLocal"
	ThanosSpecFieldLogFormat           = "logFormat"
	ThanosSpecFieldLogLevel            = "logLevel"
	ThanosSpecFieldMinTime             = "minTime"
	ThanosSpecFieldObjectStorageConfig = "objectStorageConfig"
	ThanosSpecFieldResources           = "resources"
	ThanosSpecFieldSHA                 = "sha"
	ThanosSpecFieldTag                 = "tag"
	ThanosSpecFieldTracingConfig       = "tracingConfig"
	ThanosSpecFieldVersion             = "version"
)

type ThanosSpec struct {
	BaseImage           string                `json:"baseImage,omitempty" yaml:"baseImage,omitempty"`
	GRPCServerTLSConfig *TLSConfig            `json:"grpcServerTlsConfig,omitempty" yaml:"grpcServerTlsConfig,omitempty"`
	Image               string                `json:"image,omitempty" yaml:"image,omitempty"`
	ListenLocal         bool                  `json:"listenLocal,omitempty" yaml:"listenLocal,omitempty"`
	LogFormat           string                `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`
	LogLevel            string                `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	MinTime             string                `json:"minTime,omitempty" yaml:"minTime,omitempty"`
	ObjectStorageConfig *SecretKeySelector    `json:"objectStorageConfig,omitempty" yaml:"objectStorageConfig,omitempty"`
	Resources           *ResourceRequirements `json:"resources,omitempty" yaml:"resources,omitempty"`
	SHA                 string                `json:"sha,omitempty" yaml:"sha,omitempty"`
	Tag                 string                `json:"tag,omitempty" yaml:"tag,omitempty"`
	TracingConfig       *SecretKeySelector    `json:"tracingConfig,omitempty" yaml:"tracingConfig,omitempty"`
	Version             string                `json:"version,omitempty" yaml:"version,omitempty"`
}
