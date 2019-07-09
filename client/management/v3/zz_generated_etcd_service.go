package client

const (
	ETCDServiceType              = "etcdService"
	ETCDServiceFieldBackupConfig = "backupConfig"
	ETCDServiceFieldCACert       = "caCert"
	ETCDServiceFieldCert         = "cert"
	ETCDServiceFieldCreation     = "creation"
	ETCDServiceFieldExternalURLs = "externalUrls"
	ETCDServiceFieldExtraArgs    = "extraArgs"
	ETCDServiceFieldExtraBinds   = "extraBinds"
	ETCDServiceFieldExtraEnv     = "extraEnv"
	ETCDServiceFieldImage        = "image"
	ETCDServiceFieldKey          = "key"
	ETCDServiceFieldPath         = "path"
	ETCDServiceFieldRetention    = "retention"
	ETCDServiceFieldSnapshot     = "snapshot"
)

type ETCDService struct {
	BackupConfig *BackupConfig     `json:"backupConfig,omitempty" yaml:"backup_config,omitempty"`
	CACert       string            `json:"caCert,omitempty" yaml:"ca_cert,omitempty"`
	Cert         string            `json:"cert,omitempty" yaml:"cert,omitempty"`
	Creation     string            `json:"creation,omitempty" yaml:"creation,omitempty"`
	ExternalURLs []string          `json:"externalUrls,omitempty" yaml:"external_urls,omitempty"`
	ExtraArgs    map[string]string `json:"extraArgs,omitempty" yaml:"extra_args,omitempty"`
	ExtraBinds   []string          `json:"extraBinds,omitempty" yaml:"extra_binds,omitempty"`
	ExtraEnv     []string          `json:"extraEnv,omitempty" yaml:"extra_env,omitempty"`
	Image        string            `json:"image,omitempty" yaml:"image,omitempty"`
	Key          string            `json:"key,omitempty" yaml:"key,omitempty"`
	Path         string            `json:"path,omitempty" yaml:"path,omitempty"`
	Retention    string            `json:"retention,omitempty" yaml:"retention,omitempty"`
	Snapshot     *bool             `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`
}
