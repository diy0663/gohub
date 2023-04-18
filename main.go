package main

import (
	"fmt"
	"os"

	"github.com/diy0663/gohub/app/cmd"
	"github.com/diy0663/gohub/app/cmd/make"
	"github.com/diy0663/gohub/bootstrap"
	"github.com/spf13/cobra"

	btsConfig "github.com/diy0663/gohub/config"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/console"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	//Cobra 的所有命令都基于 root 命令，相当于整个程序的入口，我们将其放置于 main.go 中
	var rootCmd = &cobra.Command{
		// 把这个设置为主命令
		Use:   "Gohub",
		Short: " Gohub 命令行应用",
		// 详细介绍
		Long: ` you can use "-h" flag to see all subcommands`,
		// rootCmd 的所有子命令 执行之前,都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// --env 的值 读取环境变量,才能知道用哪个配置文件作为配置

			config.InitConfig(cmd.Env)
			bootstrap.SetupLogger()

			// 设置 gin 的运行模式，支持 debug, release, test
			// release 会屏蔽调试信息，官方建议生产环境中使用

			// 连接数据库
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
			// 初始化缓存
			bootstrap.SetupCache()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		make.CmdMake,
		cmd.CmdSeed,
		cmd.CmdCache,
	)

	// 设置默认运行的命令 (这里设置为启动web服务 CmdServe )
	// 这里  RegisterDefaultCmd 感觉没啥卵用... 还是要用go run main.go serve   或者 air serve  去启动
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	// 注册全局参数 --env
	cmd.RegisterGlobalFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
