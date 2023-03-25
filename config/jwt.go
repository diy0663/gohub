package config

import "github.com/diy0663/gohub/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{

			// 生成 jwt 的密钥 ,在底层直接用了 app 配置文件的密钥 (app.key 这个配置)
			// 过期时间，单位是分钟，一般不超过两个小时
			"expire_minute": config.Env("EXPIRE_MINUTE", 120),
			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_minute": 86400,

			// 允许刷新时间，单位分钟，86400 为两个月，从 Token 的签名时间算起
			"max_refresh_minute": config.Env("JWT_MAX_REFRESH_TIME", 86400),
		}
	})
}
