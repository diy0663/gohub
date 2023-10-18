package requests

import (
	"github.com/diy0663/go_project_packages/verifycode"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" gorm:"email,not null;" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" gorm:"verify_code,not null;" valid:"verify_code"`
	Password   string `json:"password,omitempty" gorm:"password,not null;" valid:"password"`
}

func ResetByEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "max:30", "email"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度z最少为 6",
		},
	}
	errs := validate(data, rules, messages)

	// 验证码判断
	_data := data.(*ResetByEmailRequest)
	if ok := verifycode.NewVerifyCode().CheckAnswer(_data.Email, _data.VerifyCode); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码不正确")
	}
	return errs
}
