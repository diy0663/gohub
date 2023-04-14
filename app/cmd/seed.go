package cmd

import (
	"github.com/diy0663/gohub/database/seeders"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/seed"
	"github.com/spf13/cobra"
)

var CmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeed,
	Args:  cobra.MaximumNArgs(1), // 最多传 1 个参数
}

func runSeed(cmd *cobra.Command, args []string) {

	seeders.Initialize()
	if len(args) > 0 {
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error(" Seeder not found :" + name)
		}

	} else {
		seed.RunAll()
		console.Success("Done all seeding")
	}
}
