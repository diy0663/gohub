package auth

import (
	"github.com/diy0663/go_project_packages/response"
	v1 "github.com/diy0663/gohub/app/http/controllers/v1"

	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 表单验证
	request := requests.LoginByPasswordRequest{}
	if ok := requests.RequestValidate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	// 尝试登录
	userModel, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, "账号不存在或密码错误")
	} else {
		response.JSON(c, gin.H{
			"token": jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name),
		})
	}

}
