package client

const (
	TemplateVersionSpecType                     = "templateVersionSpec"
	TemplateVersionSpecFieldAppReadme           = "appReadme"
	TemplateVersionSpecFieldDigest              = "digest"
	TemplateVersionSpecFieldExternalID          = "externalId"
	TemplateVersionSpecFieldFiles               = "files"
	TemplateVersionSpecFieldKubeVersion         = "kubeVersion"
	TemplateVersionSpecFieldQuestions           = "questions"
	TemplateVersionSpecFieldRancherMaxVersion   = "rancherMaxVersion"
	TemplateVersionSpecFieldRancherMinVersion   = "rancherMinVersion"
	TemplateVersionSpecFieldRancherVersion      = "rancherVersion"
	TemplateVersionSpecFieldReadme              = "readme"
	TemplateVersionSpecFieldRequiredNamespace   = "requiredNamespace"
	TemplateVersionSpecFieldUpgradeVersionLinks = "upgradeVersionLinks"
	TemplateVersionSpecFieldVersion             = "version"
	TemplateVersionSpecFieldVersionDir          = "versionDir"
	TemplateVersionSpecFieldVersionName         = "versionName"
	TemplateVersionSpecFieldVersionURLs         = "versionUrls"
)

type TemplateVersionSpec struct {
	AppReadme           string            `json:"appReadme,omitempty" yaml:"app_readme,omitempty"`
	Digest              string            `json:"digest,omitempty" yaml:"digest,omitempty"`
	ExternalID          string            `json:"externalId,omitempty" yaml:"external_id,omitempty"`
	Files               map[string]string `json:"files,omitempty" yaml:"files,omitempty"`
	KubeVersion         string            `json:"kubeVersion,omitempty" yaml:"kube_version,omitempty"`
	Questions           []Question        `json:"questions,omitempty" yaml:"questions,omitempty"`
	RancherMaxVersion   string            `json:"rancherMaxVersion,omitempty" yaml:"rancher_max_version,omitempty"`
	RancherMinVersion   string            `json:"rancherMinVersion,omitempty" yaml:"rancher_min_version,omitempty"`
	RancherVersion      string            `json:"rancherVersion,omitempty" yaml:"rancher_version,omitempty"`
	Readme              string            `json:"readme,omitempty" yaml:"readme,omitempty"`
	RequiredNamespace   string            `json:"requiredNamespace,omitempty" yaml:"required_namespace,omitempty"`
	UpgradeVersionLinks map[string]string `json:"upgradeVersionLinks,omitempty" yaml:"upgrade_version_links,omitempty"`
	Version             string            `json:"version,omitempty" yaml:"version,omitempty"`
	VersionDir          string            `json:"versionDir,omitempty" yaml:"version_dir,omitempty"`
	VersionName         string            `json:"versionName,omitempty" yaml:"version_name,omitempty"`
	VersionURLs         []string          `json:"versionUrls,omitempty" yaml:"version_urls,omitempty"`
}
