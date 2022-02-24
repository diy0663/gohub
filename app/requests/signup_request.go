package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 注意首字母大写
type SignupPhoneExistRequest struct {
	// valid 标签跟 govalidator 中配置项定义的 TagIdentifier 取值对应起来
	// 注意字段首字母大写
	//  omitempty 意思是在转换json数据的时候,该字段为零值的时候就不再显示这个json字段
	Phone string `json:"phone,omitempty" valid:"phone"`
}

// 被控制器调用
func ValidValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的对应报错信息
	messages := govalidator.MapData{
		"phone": []string{
			"required: 手机号为必填项，参数名称 phone",
			"digits:  手机号长度必须为 11 位的数字",
		},
	}
	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}