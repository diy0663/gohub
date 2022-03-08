package requests

import (
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 把通用的govalidator验证方法抽取到这个方法里面来
func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}

// 每个控制器里面的参数验证处理在这里完成, 只需传入 规则,对应验证规则的函数,gin的上下文 context

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

	// 限定json格式
	if err := c.ShouldBindJSON(obj); err != nil {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		// 	"error":   err.Error(),
		// })
		// fmt.Println(err.Error())

		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return false
	}

	// 用自个传入的具体验证方法去做验证
	errs := handler(obj, c)
	if len(errs) > 0 {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "请求验证不通过，具体请查看 errors",
		// 	"errors":  errs,
		// })
		response.ValidationError(c, errs)
		return false
	}

	return true
}
