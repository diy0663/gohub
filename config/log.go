package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("log", map[string]interface{}{

		"level": configPkg.Env("LOG_LEVEL", "debug"),

		// 日志的类型，可选：
		// "single" 独立的文件
		// "daily" 按照日期每日一个
		"type": configPkg.Env("LOG_TYPE", "single"),

		/* ------------------ 滚动日志配置 ------------------ */
		// 日志文件路径
		"filename": configPkg.Env("LOG_NAME", "storage/logs/logs.log"),
		// 每个日志文件保存的最大尺寸 单位：M
		"max_size": configPkg.Env("LOG_MAX_SIZE", 64),
		// 最多保存日志文件数，0 为不限，MaxAge 到了还是会删
		"max_backup": configPkg.Env("LOG_MAX_BACKUP", 5),
		// 最多保存多少天，7 表示一周前的日志会被删除，0 表示不删
		"max_age": configPkg.Env("LOG_MAX_AGE", 30),
		// 是否压缩，压缩日志不方便查看，我们设置为 false（压缩可节省空间）
		"compress": configPkg.Env("LOG_COMPRESS", false),
	})
}
