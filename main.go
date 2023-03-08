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

	// 加载完读取配置之后,优先初始日志
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	// 连接数据库
	bootstrap.SetupDB()
	bootstrap.SetupRedis()
	bootstrap.SetupRoute(r)
	// mac M1环境下 用 127.0.0.1:3000 就不会频繁提示防火墙了
	err := r.Run(config.Get("app.mac_m1") + ":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
