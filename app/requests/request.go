package requests

import (
	"github.com/diy0663/go_project_packages/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 定义函数
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

//  被控制器直接使用
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	if err := c.ShouldBindJSON(obj); err != nil {

		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")

		return false
	}

	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)

		return false
	}
	return true

}

// 重构提取重复代码
func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
