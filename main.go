package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// cmd+shift+K 用于vscode 删除 单行
	// 自行配置 cmd+D 用于复制单行

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "World"})
	})

	r.Run()
}
