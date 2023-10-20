package cmd

import (
	"github.com/diy0663/gohub/database/seeders"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/seed"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database , the max num of args is 1  ",
	Run:   runSeed,
	// 最多只能有1个参数, 传1个参数就是要运行指定的seeder, 否则就是全部运行
	Args: cobra.MaximumNArgs(1), //
}

func runSeed(cmd *cobra.Command, args []string) {
	// 加载全部的 seeder 到全局变量里面去
	seeders.InitAllSeeder()
	if len(args) > 0 {
		name := args[0]
		thisSeeder := seed.GetSeeder(name)
		if thisSeeder.Name != "" {
			// 说明真的存在
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}

	} else {
		seed.RunAll()
		console.Success(" All is Done")
	}

}
