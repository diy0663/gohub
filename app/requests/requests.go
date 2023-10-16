package requests

import (
	requestsPkg "github.com/diy0663/go_project_packages/requests"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func RequestValidate(c *gin.Context, obj interface{}, handler requestsPkg.ValidatorFunc) bool {
	return requestsPkg.ValidateInAPI(c, obj, handler)
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	return requestsPkg.ValidateInRequest(data, rules, messages)
}
