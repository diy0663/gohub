package main

import (
	"net/http"

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
	// http://127.0.0.1:8000/ 即可访问
	// 指定8000端口
	r.Run(":8000")
}
