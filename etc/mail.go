package etc

import (
	"github.com/wuvikr/gohub01/pkg/config"
)

func init() {
	config.Add("mail", func() map[string]any {
		return map[string]any{

			// SMTP 服务器配置
			"smtp": map[string]any{
				"host":     config.GetWithDefault("MAIL_SMTP_HOST", "localhost"),
				"port":     config.GetWithDefault("MAIL_SMTP_PORT", 1025),
				"username": config.GetWithDefault("MAIL_SMTP_USERNAME", ""),
				"password": config.GetWithDefault("MAIL_SMTP_PASSWORD", ""),
			},

			// 发件人配置
			"from": map[string]any{
				"address": config.GetWithDefault("MAIL_FROM_ADDRESS", "gohub01@example.com"),
				"name":    config.GetWithDefault("MAIL_FROM_NAME", "gohub01 测试邮件"),
			},
		}
	})
}
