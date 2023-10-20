package main

import (
	"fmt"
	"os"

	"github.com/diy0663/gohub/app/cmd"
	"github.com/diy0663/gohub/app/cmd/make"
	"github.com/diy0663/gohub/bootstrap"
	btsConfig "github.com/diy0663/gohub/config"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

func init() {
	btsConfig.InitAllConfig()
	bootstrap.SetupLogger()
	//连接数据库
	bootstrap.SetupDB()
	bootstrap.SetupRedis()

}

// main  入口保持整洁
func main() {
	// cmd+shift+K 用于vscode 删除 单行
	// 自行配置 cmd+D 用于复制单行

	var rootCmd = &cobra.Command{
		// go run main.go 的作用相当于执行 Gohub
		// 剩下的,就是在后面拼接 子命令, 不传就默认拼接 serve
		Use:     "Gohub",
		Short:   "",
		Long:    `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		Example: "Gohub serve",
		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 其实, root 以及其子命令也依赖main , main 执行前 也有init方法要执行
			// 所以  某种层面这个 PersistentPreRun 达到的效果也跟 init 一样
		},
	}

	// 添加注册子命令
	rootCmd.AddCommand(
		// serve 被设置为了默认命令, 所以  go run main.go 或者  go run main.go serve 都可
		cmd.CmdServe,
		// go run main.go key 即可执行9
		cmd.CmdKey,
		make.CmdMake,
		cmd.CmdTestCommand,
		cmd.CmdSeed,
	)
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
