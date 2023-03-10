package config

import "github.com/diy0663/gohub/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{

			// 验证码的长度
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),

			// 过期时间，单位是分钟
			"expire_minute": config.Env("VERIFY_CODE_EXPIRE", 15),

			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_minute": 10080,
			// 本地开发环境验证码使用 debug_code
			"debug_code": 123456,

			// 方便本地和 API 自动测试
			// 测试手机号前缀
			"debug_phone_prefix": "0000",
			// 测试邮箱的后缀
			"debug_email_suffix": "1@testing.com",
		}
	})
}
