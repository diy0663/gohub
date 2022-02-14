package main

import (
	"fmt"

	"github.com/diy0663/gohub/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	bootstrap.SetUpRoute(router)

	// 指定8000端口
	err := router.Run(":8000")
	if err != nil {
		// 打印可能出现的错误,例如端口占用
		fmt.Println(err.Error())
	}
}
