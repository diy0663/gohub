package middlewares

import (
	"fmt"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		// 解析成功
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}
		c.Set("current_user_id", userModel.ID)
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
