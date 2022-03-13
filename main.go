package main

import (
	"fmt"
	"os"

	"github.com/diy0663/gohub/app/cmd"
	"github.com/diy0663/gohub/app/cmd/make"

	"github.com/diy0663/gohub/bootstrap"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"

	"github.com/diy0663/go_project_packages/config"
	btsConfig "github.com/diy0663/gohub/config"
)

func init() {
	btsConfig.Initialize()

}

func main() {

	var rootCmd = &cobra.Command{
		Use:   config.GetString("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		//  rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.KeyCmd,
		cmd.TinkerCmd,
		make.CmdMake,
		cmd.CmdTestCommand,
	)

	// 默认执行 CmdServe 命令
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

	// gin.SetMode(gin.ReleaseMode)
	// router := gin.New()

	// bootstrap.SetUpRoute(router)

	// // 指定端口
	// err := router.Run(":" + config.GetString("app.port"))
	// if err != nil {
	// 	// 打印可能出现的错误,例如端口占用
	// 	fmt.Println(err.Error())
	// }
}
