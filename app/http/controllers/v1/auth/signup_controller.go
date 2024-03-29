package auth

import (
	"github.com/diy0663/go_project_packages/response"
	v1 "github.com/diy0663/gohub/app/http/controllers/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/jwt"
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
	// panic("验证recovery")

	request := requests.SignupPhoneExistRequest{}
	// 要求传过来为json 格式
	if ok := requests.RequestValidate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// 返回数据库的查询结果
	// c.JSON(http.StatusOK, gin.H{
	// 	"exists": user.IsPhoneExists(request.Phone),
	// })
	response.JSON(c, gin.H{
		"exists": user.IsPhoneExists(request.Phone),
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}

	// 底层的验证有做断言类型判断,要求 验证数据参数传进去的是一个指针类型,所以才要 &request
	ok := requests.RequestValidate(c, &request, requests.ValidateSignupEmailExist)
	if !ok {
		return
	}
	// 返回数据库的查询结果
	response.JSON(c, gin.H{
		"exists": user.IsEmailExists(request.Email),
	})

}

// 根据邮件创建用户
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	request := requests.SignupUsingEmailRequest{}
	ok := requests.RequestValidate(c, &request, requests.SignupUsingEmail)
	if !ok {
		return
	}
	userModel := user.User{

		Name:  request.Name,
		Email: request.Email,

		Password: request.Password,
	}
	userModel.Create()
	if userModel.ID > 0 {
		response.JSON(c, gin.H{
			"token": jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name, 0),
			"data":  userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}

}
