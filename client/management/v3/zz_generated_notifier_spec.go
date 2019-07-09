package client

const (
	NotifierSpecType                 = "notifierSpec"
	NotifierSpecFieldClusterID       = "clusterId"
	NotifierSpecFieldDescription     = "description"
	NotifierSpecFieldDisplayName     = "displayName"
	NotifierSpecFieldPagerdutyConfig = "pagerdutyConfig"
	NotifierSpecFieldSMTPConfig      = "smtpConfig"
	NotifierSpecFieldSlackConfig     = "slackConfig"
	NotifierSpecFieldWebhookConfig   = "webhookConfig"
	NotifierSpecFieldWechatConfig    = "wechatConfig"
)

type NotifierSpec struct {
	ClusterID       string           `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Description     string           `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName     string           `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty" yaml:"pagerduty_config,omitempty"`
	SMTPConfig      *SMTPConfig      `json:"smtpConfig,omitempty" yaml:"smtp_config,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty" yaml:"slack_config,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty" yaml:"webhook_config,omitempty"`
	WechatConfig    *WechatConfig    `json:"wechatConfig,omitempty" yaml:"wechat_config,omitempty"`
}
