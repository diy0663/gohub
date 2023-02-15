package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"` //  valid 对应 TagIdentifier ,表明出现关键字 valid 就验证,验证规则就是 valid 后面跟着的规则
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	opts := govalidator.Options{
		Data: data,

		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
		FormSize:      0,
	}
	return govalidator.New(opts).ValidateStruct()

}
