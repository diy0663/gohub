package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"

	"github.com/diy0663/go_project_packages/response"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// panic("测试邮件发送")

	request := requests.SignupPhoneExistRequest{}

	// 必须 &request ,这样才能确保进行参数验证之后会赋值
	ok := requests.Validate(c, &request, requests.SignupPhoneExist)

	if !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}
