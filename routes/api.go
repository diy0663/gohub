package routes

import (
	"net/http"

	"github.com/diy0663/gohub/app/http/controllers/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// 路由分组, 版本号控制,用于方便升级接口
	v1 := r.Group("/v1")

	{
		// http://127.0.0.1:8080/v1/
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "V1!",
			})
		})

		authGroup := v1.Group("/auth")
		{
			sc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", sc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", sc.IsEmailExist)

			lc := new(auth.LoginController)
			// 登录
			authGroup.POST("/login/using-password", lc.LoginByPassword)

			// 使用邮件进行注册
			authGroup.POST("/signup/using-email", sc.SignupUsingEmail)

			vc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vc.ShowCaptcha)
			authGroup.POST("/verify-codes/email", vc.SendUsingEmail)

		}

	}

	v2 := r.Group("/v2")

	// 下面的{} 仅仅用于 类似 括起来,方便查看,以及不会搞错作用域之类的用途,用于独立处理
	{
		// http://127.0.0.1:8080/v2/
		v2.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "V2!",
			})
		})
	}

}
