package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		// 错误打印出来 (例如端口占用)
		fmt.Println(err.Error())
	}
}
