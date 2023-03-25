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
