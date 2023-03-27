package middlewares

import (
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWt().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, "授权不通过")
			return
		}

		//这里应该用缓存
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "查无对应用户")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringId())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)
		c.Next()

	}
}
