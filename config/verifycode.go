package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("verifycode", map[string]interface{}{
		"code_length": configPkg.Env("VERIFY_CODE_LENGTH", 6),

		// 过期时间，单位是分钟
		"expire_time": configPkg.Env("VERIFY_CODE_EXPIRE", 15),

		// debug 模式下的过期时间，方便本地开发调试
		"debug_expire_time": 10080,
		// 本地开发环境验证码使用 debug_code
		"debug_code": 123456,

		// 方便本地和 API 自动测试
		"debug_phone_prefix": "000",
		"debug_email_suffix": "@testing.com",
	})
}
