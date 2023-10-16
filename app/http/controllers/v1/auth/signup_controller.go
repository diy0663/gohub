package auth

import (
	"fmt"
	"net/http"

	v1 "github.com/diy0663/gohub/app/http/controllers/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/gin-gonic/gin"
)

// // 本API必须要实现的功能
// type SignupInterface interface {
// 	IsPhoneExist(c *gin.Context)
// 	IsEmailExist(c *gin.Context)
// }

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	// 要求传过来为json 格式
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	// 表单验证

	// 底层的验证有做断言类型判断,要求 验证数据参数传进去的是一个指针类型,所以才要 &request
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		// 说明有报错,验证不通过
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 返回数据库的查询结果
	c.JSON(http.StatusOK, gin.H{
		"exists": user.IsPhoneExists(request.Phone),
	})

}
