package main

import (
	"fmt"
	"gohub/bootstrap"
	btsConfig "gohub/config"

	"github.com/diy0663/go_project_packages/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.GetString("app.port"))
	if err != nil {
		// 错误打印出来 (例如端口占用)
		fmt.Println(err.Error())
	}
}
