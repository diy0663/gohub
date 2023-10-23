package cmd

import (
	"fmt"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/gohub/pkg/cache"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

func init() {
	// 注册 cache 命令的子命令
	CmdCache.AddCommand(CmdCacheClear)

}

// 子命令
var CmdCacheClear = &cobra.Command{
	Use: "clear",
	// todo 这里为何获取不到正确的 redis.database_cache 值 ? 明明是2 ,结果得到 0
	Short: fmt.Sprintf("Clear cache of cache db,the db num of redis is  %v ", config.GetString("redis.host")),
	Run:   runCacheClear,
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}
