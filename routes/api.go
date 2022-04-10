package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	// 分组
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})

		authGroup := v1.Group("/auth")
		{
			// 验证手机号是否已存在
			sign_up_controller := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", sign_up_controller.IsPhoneExist)

			// 生产图片验证码
			verify_controller := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", verify_controller.ShowCaptcha)
		}

	}
}
