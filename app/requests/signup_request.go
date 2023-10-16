package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 处理表单验证 的报错处理结果

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty"  valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone": {
			"required",
			"digits:11",
		},
	}

	messages := govalidator.MapData{
		"phone": {
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	options := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(options).ValidateStruct()
}
