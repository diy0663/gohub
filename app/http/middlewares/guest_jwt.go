package middlewares

import (
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				// 没报错说明有传token, 但是这个中间件是用来判断不需要传token的
				response.Unauthorized(c, "请使用游客身份访问")

				//c.Abort() 和 return 都可以达到中断后续处理逻辑的效果，
				//但在使用时需要考虑是否需要执行后续的 defer 函数或其他代码。
				//如果不需要执行后续的代码，则推荐使用 c.Abort()，以提高处理性能。
				//如果需要执行后续的代码，则可以使用 return 来中断处理逻辑

				c.Abort() // Abort 不会再有任何后续处理
				return

			}
		}

		c.Next()
	}
}
