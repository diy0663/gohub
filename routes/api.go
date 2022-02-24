package routes

import (
	"github.com/diy0663/gohub/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

// 把api相关的具体路由放到这里,方法要大写,这样才能被bootstrap包引用

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	// 这里其实就是个分组而已
	{
		authGroup := v1.Group("/auth")

		{
			signupController := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", signupController.IsPhoneExist)
			authGroup.POST("/signup/email/exist", signupController.IsEmailExist)
		}

	}
}
