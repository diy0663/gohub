package auth

import (
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

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExists(request.Phone),
	})

}

func (sc *SignupController) IsEmailExists(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	// 解析请求
	ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist)
	// 根据验证规则校验字段数据  todo &request
	if !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExists(request.Email),
	})

}
