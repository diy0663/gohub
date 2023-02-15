package auth

import (
	"fmt"
	"net/http"

	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExists(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	// 解析请求
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	// 根据验证规则校验字段数据  todo &request
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExists(request.Phone),
	})

}
