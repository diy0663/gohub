package cmd

import (
	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/gohub/bootstrap"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use: "serve",

	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 路由 + 中间件
	bootstrap.SetupRoute(r)
	//  http://127.0.0.1:8080
	err := r.Run(":" + config.GetString("app.port"))
	if err != nil {
		//  错误处理，端口被占用了或者其他错误
		//fmt.Println(err.Error())
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
