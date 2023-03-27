package auth

import (
	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseApiController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	// 表单验证通过之后尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "该账号不存在")
	}
	token := jwt.NewJWt().IssueToken(user.GetStringId(), user.Name)
	response.JSON(c, gin.H{
		"token": token,
	})

}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}
	// 表单验证通过之后尝试登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		// 不要给出精确的报错提示
		response.Error(c, err, "该账号不存在或密码错误")
	}
	token := jwt.NewJWt().IssueToken(user.GetStringId(), user.Name)
	response.JSON(c, gin.H{
		"token": token,
	})

}
