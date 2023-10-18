package requests

import (
	"github.com/diy0663/go_project_packages/captcha"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" gorm:"captcha_id,not null;" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" gorm:"captcha_answer,not null;" valid:"captcha_answer"`
	LoginID       string `json:"login_id,omitempty" gorm:"login_id,not null;" valid:"login_id"`
	Password      string `json:"password,omitempty" gorm:"password,not null;" valid:"password"`
}

func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"login_id": []string{
			"required:登录 ID 为必填项，支持手机号、邮箱和用户名",
			"min:登录 ID 长度需大于 3",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度最小为 6",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	//额外验证
	_data := data.(*LoginByPasswordRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "验证码不正确")
	}
	return errs

}
