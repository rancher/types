package client

const (
	PublishImageConfigType                = "publishImageConfig"
	PublishImageConfigFieldBuildContext   = "buildContext"
	PublishImageConfigFieldDockerfilePath = "dockerfilePath"
	PublishImageConfigFieldPushRemote     = "pushRemote"
	PublishImageConfigFieldRegistry       = "registry"
	PublishImageConfigFieldTag            = "tag"
)

type PublishImageConfig struct {
	BuildContext   string `json:"buildContext,omitempty" yaml:"build_context,omitempty"`
	DockerfilePath string `json:"dockerfilePath,omitempty" yaml:"dockerfile_path,omitempty"`
	PushRemote     bool   `json:"pushRemote,omitempty" yaml:"push_remote,omitempty"`
	Registry       string `json:"registry,omitempty" yaml:"registry,omitempty"`
	Tag            string `json:"tag,omitempty" yaml:"tag,omitempty"`
}
