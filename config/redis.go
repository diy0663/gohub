package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("redis", map[string]interface{}{
		"host":     configPkg.Env("REDIS_HOST", "127.0.0.1"),
		"port":     configPkg.Env("REDIS_PORT", "6379"),
		"password": configPkg.Env("REDIS_PASSWORD", ""),

		// 业务类存储使用 1 (图片验证码、短信验证码、会话)
		"database": configPkg.Env("REDIS_MAIN_DB", 1),
	})
}
