package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("mail", map[string]interface{}{

		// 默认是 Mailhog 的配置
		"smtp": map[string]interface{}{
			"host":     configPkg.Env("MAIL_HOST", "localhost"),
			"port":     configPkg.Env("MAIL_PORT", 1025),
			"username": configPkg.Env("MAIL_USERNAME", ""),
			"password": configPkg.Env("MAIL_PASSWORD", ""),
		},

		"from": map[string]interface{}{
			"address": configPkg.Env("MAIL_FROM_ADDRESS", "gohub@example.com"),
			"name":    configPkg.Env("MAIL_FROM_NAME", "Gohub"),
		},
	})
}
