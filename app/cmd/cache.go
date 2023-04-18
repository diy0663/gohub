package cmd

import (
	"github.com/diy0663/gohub/pkg/cache"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "",
}

// 子命令 clear
var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache ",
	Run:   runCacheClear,
}

func init() {
	CmdCache.AddCommand(CmdCacheClear)
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared")
}
