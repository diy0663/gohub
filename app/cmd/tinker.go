package cmd

import (
	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var TinkerCmd = &cobra.Command{
	Use:   "tinker",
	Short: "run your test code of main  in tinker ",
	Run:   tinker,
	Args:  cobra.NoArgs, //
}

func tinker(cmd *cobra.Command, args []string) {
	console.Success("这里是 cmd.tinker 方法的测试")
	// 在这里写准备在main.go 里面的调试代码
	// 即使调试代码忘记删掉, 也不会影响主流程(主流程不会是用这个tinker 命令去启动)
	// redis.Redis.Set("rdkey", "11111", 10*time.Minute)
	// console.Success("rdkey is :" + redis.Redis.Get("rdkey"))
}
