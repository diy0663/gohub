package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    ":" + config.GetString("app.port"),
		Handler: r,
	}

	// 启动服务
	go func() {
		//  http://127.0.0.1:8080
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			//  错误处理，
			//fmt.Println(err.Error())
			logger.ErrorString("CMD", "serve", err.Error())
			console.Exit("Unable to start server, error:" + err.Error())
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	console.Warning("ready to Shutdowned Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		console.Error("Server Shutdown Error:" + err.Error())
	}

	console.Success("Server exiting")

}
