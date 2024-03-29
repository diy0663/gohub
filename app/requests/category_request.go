package requests

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// todo  一个控制器里面, 多个1个api 就得用到 1个 struct 以及一个对应的验证方法

type CategoryRequest struct {
	Name        string `valid:"name" json:"name"`
	Description string `valid:"description" json:"description,omitempty"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {

	not_exists_by_name := "not_exists:categories,name"
	if c.Param("id") != "" {
		id, _ := strconv.Atoi(c.Param("id"))
		if id > 0 {
			not_exists_by_name += ",id<>" + strconv.Itoa(id)
		}
	}
	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", not_exists_by_name},
		"description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
			"not_exists:分类名称已存在",
		},
		"description": []string{
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}
