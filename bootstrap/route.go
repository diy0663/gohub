package bootstrap

import (
	"net/http"
	"strings"

	"github.com/diy0663/gohub/routes"
	"github.com/gin-gonic/gin"
)

// 注册路由
func SetupRoute(router *gin.Engine) {

	// 注册全局路由
	registerGlobalMiddleWare(router)

	//  注册API路由 , 这里的路由定义放到另外的一个文件去
	routes.RegisterAPIRoutes(router)

	// 配置404 路由
	setup404Handler(router)

}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(gin.Logger(), gin.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 返回json格式的提醒
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "请确认 url 和请求方法是否正确",
			})
		}
	})
}
