package main

import (
	"flag"
	"fmt"

	"github.com/diy0663/gohub/bootstrap"

	btsConfig "github.com/diy0663/gohub/config"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 指定.env 文件, 例如 --env=testing 对应 .env.testing , 不传则默认.env 文件")
	flag.Parse()
	config.InitConfig(env)

	r := gin.New()
	bootstrap.SetupRoute(r)
	// r.Use(gin.Logger(), gin.Recovery())

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"hello": "gin from :8000",
	// 	})
	// })

	// // 处理 404 page not found 的无效路由
	// r.NoRoute(func(c *gin.Context) {
	// 	acceptString := c.Request.Header.Get("Accept")
	// 	if strings.Contains(acceptString, "text/html") {
	// 		c.String(http.StatusNotFound, "页面返回 404")
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"error_code":    404,
	// 			"error_message": " 路由未定义",
	// 		})
	// 	}
	// })

	// mac M1环境下 用 127.0.0.1:3000 就不会频繁提示防火墙了
	err := r.Run(config.Get("app.mac_m1") + ":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
