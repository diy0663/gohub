package requests

import (
	"github.com/diy0663/go_project_packages/captcha"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 触发 发送邮件验证码 的请求
type VerifyCodeEmailRequest struct {
	Email         string `json:"email,omitempty"  valid:"email"`
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty"  valid:"captcha_answer"`
}

func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{
			"required",
			"min:4",
			"max:30",
			"email",
		},
		"captcha_id": []string{"required"},
		// 在生成的时候就已经限定是数字了, 假如不是数字验证码,就得去掉 digits
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:email为必填项,参数名称 email",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为6位数字",
		},
	}
	errsMap := validate(data, rules, messages)

	// 在这里面自行验证验证码是否正确,去redis 里面查
	_data := data.(*VerifyCodeEmailRequest)

	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errsMap["captcha_answer"] = append(errsMap["captcha_answer"], "图片验证码错误")
	}

	return errsMap
}
