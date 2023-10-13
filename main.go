package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// cmd+shift+K 用于vscode 删除 单行
	// 自行配置 cmd+D 用于复制单行
	r := gin.New()
	// 中间件
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "Gin"})
	})

	// 专门对应404 的路由
	r.NoRoute(func(c *gin.Context) {
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

	//  http://127.0.0.1:8080
	r.Run(":8080")

}
