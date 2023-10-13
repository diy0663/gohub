package bootstrap

import (
	"net/http"
	"strings"

	"github.com/diy0663/gohub/routes"
	"github.com/gin-gonic/gin"
)

// 统一被外面引用初始化全部的路由入口
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件 ,先写,之后提取出去作为单独的函数
	registerGlobalMiddleWare(router)

	// 注册 API类型的路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 的专门路由
	setup404Handler(router)

}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "查无该页面")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "请检查请求路由",
			})
		}
	})

}
