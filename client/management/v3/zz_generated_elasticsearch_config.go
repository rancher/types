package client

const (
	ElasticsearchConfigType               = "elasticsearchConfig"
	ElasticsearchConfigFieldAuthPassword  = "authPassword"
	ElasticsearchConfigFieldAuthUserName  = "authUsername"
	ElasticsearchConfigFieldCertificate   = "certificate"
	ElasticsearchConfigFieldClientCert    = "clientCert"
	ElasticsearchConfigFieldClientKey     = "clientKey"
	ElasticsearchConfigFieldClientKeyPass = "clientKeyPass"
	ElasticsearchConfigFieldDateFormat    = "dateFormat"
	ElasticsearchConfigFieldEndpoint      = "endpoint"
	ElasticsearchConfigFieldIndexPrefix   = "indexPrefix"
	ElasticsearchConfigFieldSSLVerify     = "sslVerify"
	ElasticsearchConfigFieldSSLVersion    = "sslVersion"
)

type ElasticsearchConfig struct {
	AuthPassword  string `json:"authPassword,omitempty" yaml:"auth_password,omitempty"`
	AuthUserName  string `json:"authUsername,omitempty" yaml:"auth_username,omitempty"`
	Certificate   string `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	ClientCert    string `json:"clientCert,omitempty" yaml:"client_cert,omitempty"`
	ClientKey     string `json:"clientKey,omitempty" yaml:"client_key,omitempty"`
	ClientKeyPass string `json:"clientKeyPass,omitempty" yaml:"client_key_pass,omitempty"`
	DateFormat    string `json:"dateFormat,omitempty" yaml:"date_format,omitempty"`
	Endpoint      string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	IndexPrefix   string `json:"indexPrefix,omitempty" yaml:"index_prefix,omitempty"`
	SSLVerify     bool   `json:"sslVerify,omitempty" yaml:"ssl_verify,omitempty"`
	SSLVersion    string `json:"sslVersion,omitempty" yaml:"ssl_version,omitempty"`
}
