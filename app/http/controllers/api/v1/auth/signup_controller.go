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
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 初始化验证参数request数据
	request := requests.SignupPhoneExistRequest{}

	// 解析json (在这里限定了传参要json格式)
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	err := requests.ValidValidateSignupPhoneExist(&request, c)
	if len(err) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err,
		})
		return
	}
	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exists": user.IsPhoneExist(request.Phone),
	})

}
