package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"hello": "world!"})
	// })
	// r.Run()
	r := gin.New()
	// 使用中间件
	r.Use(gin.Logger(), gin.Recovery())
	// 定义一个路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world!"})
	})
	// 没有可匹配的路由的处理方式
	r.NoRoute(func(c *gin.Context) {
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

	// 指定运行端口
	r.Run(":8000")

}
