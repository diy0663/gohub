package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	options := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(options).ValidateStruct()
}

// 验证函数类型
type ValidatorFunc func(data interface{}, c *gin.Context) map[string][]string

// Validate 控制器里调用示例：
//
//	if ok := requests.Validate(c, &requests.UserSaveRequest{}, requests.UserSave); !ok {
//	    return
//	}
//
// 参数分别是 gin的context , 请求空结构, 请求的对应验证方法(里面有rules 以及 message 报错信息的定制)
// 专门给控制器调用
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

	//	todo  是不是提前判断报错  obj 必须是指针类型 会好一点
	if err := c.ShouldBindJSON(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	errs := handler(obj, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}

	// 先解析json格式并且做绑定关系
	return true

}
