package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("database", map[string]interface{}{
		"connection": configPkg.Get("DB_CONNECTION", "mysql"),
		"mysql": map[string]interface{}{
			// 数据库连接信息
			"host":     configPkg.Env("DB_HOST", "127.0.0.1"),
			"port":     configPkg.Env("DB_PORT", "3306"),
			"database": configPkg.Env("DB_DATABASE", "gohub"),
			"username": configPkg.Env("DB_USERNAME", ""),
			"password": configPkg.Env("DB_PASSWORD", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": configPkg.Env("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": configPkg.Env("DB_MAX_OPEN_CONNECTIONS", 100),
			"max_life_seconds":     configPkg.Env("DB_MAX_LIFE_SECONDS", 10),
		},
	})
}
