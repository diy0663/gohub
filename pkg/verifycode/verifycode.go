package verifycode

import (
	"fmt"
	"strings"
	"sync"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/helper"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/diy0663/gohub/pkg/mail"
	"github.com/diy0663/gohub/pkg/redis"
	"github.com/diy0663/gohub/pkg/sms"
)

type VerifyCode struct {
	Store Store
}

// 以下once, internalVerifyCode, NewVerifyCode() 一起配合实现单例模式
var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				// 在这里指定了 key 的前缀
				KeyPrefix: config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

// 对于短信验证码, 发送短信的时候只需要传手机号码, 其余参数,(模板,内容可以写死, 数字验证码在内部调用生成)
func (vc *VerifyCode) SendSMS(phone string) bool {
	// 生成数字验证码(生成的同时存入redis中)
	code := vc.generateVerifyCode(phone)

	// 本地测试不需要真实发短信
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return sms.NewSMS().Send(phone, sms.Message{
		// 从配置文件里面指定 发短信验证码的第三方短信模板编号,算是写死
		Template: config.GetString("sms.aliyun.template_code"),
		Data: map[string]string{
			"code": code,
		},
		Content: "",
	})

}

// key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})
	//方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	if !app.IsProduction() && (strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
		strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}
	return vc.Store.Verify(key, answer, false)

}

func (vc *VerifyCode) SendEmail(email string) bool {
	// 生成数字验证码(生成的同时存入redis中)
	code := vc.generateVerifyCode(email)

	// 本地测试不需要真实发出
	// if !app.IsProduction() && strings.HasSuffix(email, config.GetString("verifycode.debug_email_suffix")) {
	// 	return true
	// }
	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)

	return mail.NewMail().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Bcc:     []string{},
		Cc:      []string{},
		Subject: "Email 验证码",
		Text:    []byte{},
		// 强制类型转换, string 转 []byte
		HTML: []byte(content),
	})

}

// 内部私有方法
func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := helper.RandomNumber(config.GetInt("verifycode.code_length"))

	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})
	// 这里不判断存数据到redis 的时候是否出错??
	vc.Store.Set(key, code)
	return code
}
