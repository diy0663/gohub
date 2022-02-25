package config

import "github.com/diy0663/go_project_packages/config"

func init() {

	config.Add("redis", config.StrMap{

		"host":     config.Env("REDIS_HOST", "127.0.0.1"),
		"port":     config.Env("REDIS_PORT", "63799"),
		"password": config.Env("REDIS_PASSWORD", ""),
		// 业务类存储使用 1 (图片验证码、短信验证码、会话) ,注意要跟缓存那边的分开
		"database": config.Env("REDIS_MAIN_DB", 1),
	})
}
