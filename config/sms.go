package config

import "github.com/diy0663/gohub/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认是阿里云的测试 sign_name 和 template_code
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME", "阿里云短信测试"),
				// 下面这个 template_code 就是专门发短信验证码的那个 特定模板code
				"template_code": config.Env("SMS_ALIYUN_TEMPLATE_CODE", "SMS_154950909"),
			},
		}
	})
}
