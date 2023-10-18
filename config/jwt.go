package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("jwt", map[string]interface{}{
		// 过期时间，注意 单位是分钟，一般不超过两个小时
		"expire_minute": configPkg.Env("JWT_EXPIRE_MINUTE", 120),

		// 允许刷新时间，单位分钟，86400 为两个月，从 Token 的签名时间算起
		"max_refresh_minute": configPkg.Env("JWT_MAX_REFRESH_MINUTE", 86400),

		// debug 模式下的过期时间，方便本地开发调试
		"debug_expire_minute": 86400,
	})
}
