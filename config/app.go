package config

import "github.com/diy0663/gohub/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name": config.Env("APP_NAME", "Gohub"),
			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": config.Env("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),

			// 应用服务端口
			"port":   config.Env("APP_PORT", "3000"),
			"mac_m1": config.Env("APP_MAC_M1", ""),
			// 加密会话、JWT 加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea0111160a532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:3000"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
