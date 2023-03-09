package config

import "github.com/diy0663/gohub/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			"expire_minute":       config.Env("CAPTCHA_EXPIRE_MINUTE", 10),
			"debug_expire_minute": config.Env("CAPTCHA_DEBUG_EXPIRE_MINUTE", 50000),
			// 验证码图片高度
			"height": 80,

			// 验证码图片宽度
			"width": 240,

			// 验证码的长度
			"length": 6,

			// 数字的最大倾斜角度
			"maxskew": 0.7,

			// 图片背景里的混淆点数量
			"dotcount": 80,
			// 非 production 环境，使用此 key 可跳过验证，方便测试
			"testing_key": "captcha_skip_special",
		}
	})
}
