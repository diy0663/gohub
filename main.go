package main

import (
	"fmt"

	"github.com/diy0663/gohub/bootstrap"

	"github.com/diy0663/go_project_packages/config"
	btsConfig "github.com/diy0663/gohub/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()

}

func main() {

	router := gin.New()
	bootstrap.SetupDB()
	bootstrap.SetUpRoute(router)

	// 指定端口
	err := router.Run(":" + config.GetString("app.port"))
	if err != nil {
		// 打印可能出现的错误,例如端口占用
		fmt.Println(err.Error())
	}
}
