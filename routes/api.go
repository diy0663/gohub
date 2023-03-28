package routes

import (
	"net/http"

	"github.com/diy0663/gohub/app/http/controllers/api/v1/auth"
	"github.com/diy0663/gohub/app/http/middlewares"
	auth_jwt "github.com/diy0663/gohub/pkg/auth"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// 在这里做 api 服务的 每一个具体的路由定义
func RegisterAPIRoutes(r *gin.Engine) {
	// 使用路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "V1",
			})
		})

		authGroup := v1.Group("/auth")
		{
			signup_controller := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", signup_controller.IsPhoneExists)
			authGroup.POST("/signup/email/exist", signup_controller.IsEmailExists)

			// 传入账号密码手机验证码 通过手机号完成注册
			authGroup.POST("/signup/using-phone", signup_controller.SignupUsingPhone)
			// 根据邮件创建用户
			authGroup.POST("/signup/using-email", signup_controller.SignupUsingEmail)

			// 手机号+短信验证码进行验证  ,前提是要触发一次发送短信的操作
			login_controller := new(auth.LoginController)
			authGroup.POST("/login/using-phone", login_controller.LoginByPhone)
			authGroup.POST("/login/using-password", login_controller.LoginByPassword)
			authGroup.POST("/login/refresh-token", login_controller.RefreshToken)

			verify_code_controller := new(auth.VerifyCodeController)
			// 生成图片验证码
			authGroup.POST("/verify-codes/captcha", verify_code_controller.ShowCaptcha)
			// 发短信验证码
			authGroup.POST("/verify-codes/phone", verify_code_controller.SendUsingPhone)
			// 发邮件验证码
			authGroup.POST("/verify-codes/email", verify_code_controller.SendUsingEmail)

			password_reset_controller := new(auth.PasswordResetController)
			authGroup.POST("/password-reset/using-phone", password_reset_controller.ResetByPhone)

		}
		v1.GET("/test_auth", middlewares.AuthJWT(), func(c *gin.Context) {
			userModel := auth_jwt.CurrentUser(c)
			response.Data(c, userModel)
		})

	}
}
