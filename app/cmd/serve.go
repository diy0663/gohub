package cmd

import (
	"github.com/diy0663/gohub/bootstrap"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// 启动这个web-api服务 的方式改用这个命令去做

var CmdServe = &cobra.Command{
	// 命令关键字
	Use: "serve",
	//简介
	Short: "start web server",
	// 处理逻辑对应的函数
	Run: runWeb,
	// todo  ?? 为啥这里不要参数
	Args: cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	bootstrap.SetupRoute(r)
	// mac M1环境下 用 127.0.0.1:3000 就不会频繁提示防火墙了
	err := r.Run(config.Get("app.mac_m1") + ":" + config.Get("app.port"))

	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}

}
