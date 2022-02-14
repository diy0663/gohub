package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// 注册两个用到的全局中间件
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"key": "value",
		})
	})

	// 新增一个专门的404处理路由
	r.NoRoute(func(c *gin.Context) {
		// 分情况处理,看要返回json格式还是返回html格式
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认要返回json格式
			c.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_msg":  " 未定义的路由!",
			})
		}
	})

	// http://127.0.0.1:8000/ 即可访问
	// 指定8000端口
	r.Run(":8000")
}
