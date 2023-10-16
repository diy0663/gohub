package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("app", map[string]interface{}{
		"name": configPkg.Env("APP_NAME", "Gohub"),
		// 当前环境，用以区分多环境，一般为 local, stage, production, test
		"env": configPkg.Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": configPkg.Env("APP_DEBUG", false),

		// 应用服务端口
		"port": configPkg.Env("APP_PORT", "3000"),

		// 加密会话、JWT 加密
		"key": configPkg.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

		// 用以生成链接
		"url": configPkg.Env("APP_URL", "http://localhost:3000"),

		// 设置时区，JWT 里会使用，日志记录里也会使用到
		"timezone": configPkg.Env("TIMEZONE", "Asia/Shanghai"),
	})
}
