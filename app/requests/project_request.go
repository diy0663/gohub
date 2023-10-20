package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// todo  一个控制器里面, 多个1个api 就得用到 1个 struct 以及一个对应的验证方法

type ProjectRequest struct {
	Name     string `json:"name,omitempty" valid:"name"`
	Comments string `json:"comments,omitempty" valid:"comments"`
}

func ProjectSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":     []string{"required", "min_cn:2", "max_cn:8", "not_exists:projects,name"},
		"comments": []string{"required", "min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
			"not_exists:项目名称已存在",
		},
		"comments": []string{
			"required:备注为必填项",
			"min_cn:备注长度需至少 3 个字",
			"max_cn:备注长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}
