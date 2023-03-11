package config

import "github.com/diy0663/gohub/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认是 Mailhog 的配置
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_SMTP_HOST", "localhost"),
				"port":     config.Env("MAIL_SMTP_PORT", 1025),
				"username": config.Env("MAIL_SMTP_USERNAME", ""),
				"password": config.Env("MAIL_SMTP_PASSWORD", ""),
			},

			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "xx@xx.com"),
				"name":    config.Env("MAIL_FROM_NAME", "theSender"),
			},
		}
	})
}
