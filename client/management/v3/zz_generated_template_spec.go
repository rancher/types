package client

const (
	TemplateSpecType                          = "templateSpec"
	TemplateSpecFieldCatalogID                = "catalogId"
	TemplateSpecFieldCategories               = "categories"
	TemplateSpecFieldCategory                 = "category"
	TemplateSpecFieldClusterCatalogID         = "clusterCatalogId"
	TemplateSpecFieldClusterID                = "clusterId"
	TemplateSpecFieldDefaultTemplateVersionID = "defaultTemplateVersionId"
	TemplateSpecFieldDefaultVersion           = "defaultVersion"
	TemplateSpecFieldDescription              = "description"
	TemplateSpecFieldDisplayName              = "displayName"
	TemplateSpecFieldFolderName               = "folderName"
	TemplateSpecFieldIcon                     = "icon"
	TemplateSpecFieldIconFilename             = "iconFilename"
	TemplateSpecFieldMaintainer               = "maintainer"
	TemplateSpecFieldPath                     = "path"
	TemplateSpecFieldProjectCatalogID         = "projectCatalogId"
	TemplateSpecFieldProjectID                = "projectId"
	TemplateSpecFieldProjectURL               = "projectURL"
	TemplateSpecFieldReadme                   = "readme"
	TemplateSpecFieldUpgradeFrom              = "upgradeFrom"
	TemplateSpecFieldVersions                 = "versions"
)

type TemplateSpec struct {
	CatalogID                string                `json:"catalogId,omitempty" yaml:"catalog_id,omitempty"`
	Categories               []string              `json:"categories,omitempty" yaml:"categories,omitempty"`
	Category                 string                `json:"category,omitempty" yaml:"category,omitempty"`
	ClusterCatalogID         string                `json:"clusterCatalogId,omitempty" yaml:"cluster_catalog_id,omitempty"`
	ClusterID                string                `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	DefaultTemplateVersionID string                `json:"defaultTemplateVersionId,omitempty" yaml:"default_template_version_id,omitempty"`
	DefaultVersion           string                `json:"defaultVersion,omitempty" yaml:"default_version,omitempty"`
	Description              string                `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName              string                `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	FolderName               string                `json:"folderName,omitempty" yaml:"folder_name,omitempty"`
	Icon                     string                `json:"icon,omitempty" yaml:"icon,omitempty"`
	IconFilename             string                `json:"iconFilename,omitempty" yaml:"icon_filename,omitempty"`
	Maintainer               string                `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`
	Path                     string                `json:"path,omitempty" yaml:"path,omitempty"`
	ProjectCatalogID         string                `json:"projectCatalogId,omitempty" yaml:"project_catalog_id,omitempty"`
	ProjectID                string                `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	ProjectURL               string                `json:"projectURL,omitempty" yaml:"project_url,omitempty"`
	Readme                   string                `json:"readme,omitempty" yaml:"readme,omitempty"`
	UpgradeFrom              string                `json:"upgradeFrom,omitempty" yaml:"upgrade_from,omitempty"`
	Versions                 []TemplateVersionSpec `json:"versions,omitempty" yaml:"versions,omitempty"`
}
