package routes

import (
	"net/http"

	"github.com/diy0663/gohub/app/http/controllers/v1/auth"
	"github.com/diy0663/gohub/app/http/middlewares"
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
			// 登录 (加上游客中间件)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), middlewares.LimitIP("5-M"), lc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lc.RefreshToken)

			// 使用邮件进行注册 (在邮件注册这个路由上,每个IP每分钟最多请求5次  ,注意失败次数也算在内)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), middlewares.LimitPerRoute("5-M"), sc.SignupUsingEmail)

			vc := new(auth.VerifyCodeController)
			// 基于IP做限量,获取图片验证码,每分钟限制20次
			authGroup.POST("/verify-codes/captcha", middlewares.LimitIP("20-M"), vc.ShowCaptcha)
			authGroup.POST("/verify-codes/email", vc.SendUsingEmail)

			pwc := new(auth.PasswordController)
			authGroup.POST("password-reset/using-email", middlewares.AuthJWT(), pwc.ResetByEmail)

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
