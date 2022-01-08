package main

import (
	"fmt"

	"github.com/diy0663/gohub/bootstrap"
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
	bootstrap.SetupRoute(r)

	// 指定运行端口
	err := r.Run(":3000")
	if err != nil {
		// 打印错误
		fmt.Println(err.Error())
	}

}
