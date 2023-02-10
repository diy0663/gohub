package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "gin from :8000",
		})
	})

	// 处理 404 page not found 的无效路由
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": " 路由未定义",
			})
		}
	})

	r.Run("127.0.0.1:8000")
}
