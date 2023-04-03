package cmd

import (
	"os"

	"github.com/diy0663/gohub/pkg/helper"
	"github.com/spf13/cobra"
)

// 全局变量 --env 的值放这里
var Env string

// 读取输入参数 --env 的值,赋予Env 值
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	// 参数依次是 要赋值的变量名, 命令行输入的参数名, 命令行输入的参数名简写, 默认值, 使用说明
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, example: --env=testing will use .env.testing file")
}

// 注册默认命令 (感觉没啥卵用. 看着是把子命令的参数也给一份到主命令的参数上 )
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {

	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helper.FirstElement(os.Args[1:])
	if err != nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

}
