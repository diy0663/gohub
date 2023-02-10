package bootstrap

import (
	"net/http"
	"strings"

	"github.com/diy0663/gohub/routes"
	"github.com/gin-gonic/gin"
)

// 在这里做路由相关的启动项加载或者初始化工作
// 例如 全局中间件开启  设置404 对应的默认路由 设定一堆可直接调用的函数

func SetupRoute(router *gin.Engine) {
	// 注意顺序
	// 先注册全局路由
	registerGlobalMiddleWare(router)
	// 注册API层的路由 , api 层面的路由会有很多,所以要拆分到特定的文件那边去, 方便维护
	routes.RegisterAPIRoutes(router)

	// 配置404 路由
	setup404Handler(router)

}

// 只在本文件用到, 直接小写字母开头
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
