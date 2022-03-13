package cmd

import (
	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/gohub/bootstrap"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	// 指定这个命令要运行的方法
	Run: runWeb,

	Args: cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	// release 会屏蔽调试信息，官方建议生产环境中使用
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	bootstrap.SetUpRoute(router)

	err := router.Run(":" + config.GetString("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}

}
