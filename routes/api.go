package routes

import (
	"net/http"

	"github.com/diy0663/gohub/app/http/controllers/api/v1/auth"
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

			verify_code_controller := new(auth.VerifyCodeController)
			// 生成图片验证码
			authGroup.POST("/verify-codes/captcha", verify_code_controller.ShowCaptcha)
			// 发短信验证码
			authGroup.POST("/verify-codes/phone", verify_code_controller.SendUsingPhone)
			// 发邮件验证码
			authGroup.POST("/verify-codes/email", verify_code_controller.SendUsingEmail)

		}

	}
}
