package bootstrap

import (
	"net/http"
	"strings"

	"github.com/diy0663/gohub/app/http/middlewares"
	"github.com/diy0663/gohub/routes"
	"github.com/gin-gonic/gin"
)

// 把main 里面的路由初始化引入到这里

func SetUpRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册具体的API路由
	routes.RegisterAPIRoutes(router)

	// 处理404
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(middlewares.Logger(), middlewares.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
