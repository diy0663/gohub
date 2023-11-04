package auth

import (
	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/response"
	v1 "github.com/diy0663/gohub/app/http/controllers/v1"
	"github.com/diy0663/gohub/app/models/token"

	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}

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

		tokenData := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name, userModel.RoleId)

		//非生产环境, 保存一次 token记录
		if config.GetString("app.env") != "production" {
			tokenModel := token.Token{
				UserID:      userModel.ID,
				TokenString: tokenData,
				LoginId:     request.LoginID,
				ExpireTime:  uint64(jwt.NewJWT().ExpireAtTime()),
			}
			tokenModel.Create()
			if tokenModel.ID == 0 {
				panic("token 记录出错")
			}
		}

		response.JSON(c, gin.H{
			"token": tokenData,
		})
	}

}
