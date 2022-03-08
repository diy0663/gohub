package auth

import (
	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 初始化验证参数request数据
	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// 检查数据库并返回响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"exists": user.IsPhoneExist(request.Phone),
	// })
	response.JSON(c, gin.H{
		"exists": user.IsPhoneExist(request.Phone),
	})

}

// 验证邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 初始化验证参数request数据
	request := requests.SignupEmailExistRequest{}

	// 这里面传 &request , 经过 ShouldBindJSON并且验证通过是会对request进行赋值的
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	// 检查数据库并返回响应

	response.JSON(c, gin.H{
		"exists": user.IsEmailExist(request.Email),
	})

}
