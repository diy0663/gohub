package routes

import (
	"fmt"
	"net/http"
	"time"

	// 别名, 因为 v1 以及早被 路由分组用了
	v1_controller "github.com/diy0663/gohub/app/http/controllers/v1"
	"github.com/diy0663/gohub/app/http/controllers/v1/auth"
	"github.com/diy0663/gohub/app/http/middlewares"
	"github.com/diy0663/gohub/app/models/role"
	casbinpkg "github.com/diy0663/gohub/pkg/casbinPkg"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// 路由分组, 版本号控制,用于方便升级接口
	v1 := r.Group("/v1")

	{
		// http://127.0.0.1:3000/v1/
		v1.GET("/", func(c *gin.Context) {
			// 验证全局中间件10s 超时是否有效
			time.Sleep(11 * time.Second)
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

		// 这里面需要注意, 我们路由分组用了v1, 第一版本的控制器包名也是v1,所以只能使用别名 v1_controller 来给控制器这边用
		uc_controller := new(v1_controller.UsersController)
		v1.GET("/user", middlewares.AuthJWT(), middlewares.CasbinRule(), uc_controller.CurrentUser)

		// 用户列表
		userGroup := v1.Group("/users")
		{
			userGroup.GET("", uc_controller.Index)
		}

		projectGroup := v1.Group("/projects")
		pj_controller := new(v1_controller.ProjectsController)
		{
			projectGroup.GET("", pj_controller.Index)
		}

		cgc := new(v1_controller.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
			cgcGroup.GET("", cgc.Index)
		}
		tpc := new(v1_controller.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
			tpcGroup.GET("/:id", tpc.Show)
		}
		lsc := new(v1_controller.LinksController)
		linksGroup := v1.Group("/links")
		{
			linksGroup.GET("", lsc.Index)
			linksGroup.GET("/:id", lsc.Show)
		}
	}

	v2 := r.Group("/v2")

	// 下面的{} 仅仅用于 类似 括起来,方便查看,以及不会搞错作用域之类的用途,用于独立处理
	{
		// http://127.0.0.1:3000/v2/
		v2.GET("/", func(c *gin.Context) {
			// 停3秒,验证是否能优雅关闭(处理完请求再关闭 超时5秒关闭)
			time.Sleep(3 * time.Second)
			c.JSON(http.StatusOK, gin.H{
				"Hello": "V2!",
			})
		})
		v2.GET("/timeout", timeout.New(
			// 设置超时3秒
			timeout.WithTimeout(time.Second*3),
			// 设置超时之后的返回提示结果
			timeout.WithResponse(func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"Hello": "the request is time-out !!",
				})
			}),
			// 设置这个路由本身的handler
			timeout.WithHandler(func(c *gin.Context) {
				time.Sleep(1 * time.Second)
				c.JSON(http.StatusOK, gin.H{
					"Hello": " the request is finish  !",
				})
			}),
		))
	}

	r.GET("getAllRoutes", func(ctx *gin.Context) {
		all := r.Routes()
		permissions := make([]casbinpkg.PermissionInfo, 1)
		for _, v := range all {
			permission := casbinpkg.PermissionInfo{
				Path:   v.Path,
				Method: v.Method,
			}

			permissions = append(permissions, permission)
		}
		fmt.Println(permissions)
		// 指定针对 超级管理员做所有的授权记录插入
		//superRole := role.GetByName("超级管理员")
		superRole := role.GetBy("name", "超级管理员")
		if superRole.ID != 0 {
			err := casbinpkg.UpdatePermissionByRoleId(int(superRole.ID), permissions)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("超级管理员权限插入成功")
			}

		} else {
			fmt.Println("超级管理员角色不存在")
		}

	})

}
