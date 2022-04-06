package config

import "github.com/diy0663/go_project_packages/config"

// 等待被自动触发
func init() {

	config.Add("app", config.StrMap{
		"name": config.Env("APP_NAME", "GoHub"),
		// 当前环境，用以区分多环境，一般为 local, stage, production, test
		"env": config.Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),

		// 应用服务端口
		"port": config.Env("APP_PORT", "3000"),

		// 加密会话、JWT 加密
		"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

		// 用以生成链接
		"url": config.Env("APP_URL", "http://localhost:3000"),

		// 设置时区，JWT 里会使用，日志记录里也会使用到
		"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
	})

}
