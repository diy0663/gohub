package middlewares

import (
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWt().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
