package cmd

import (
	"fmt"

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

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key , example: cache forget cache-key ",
	Run:   runCacheForget,
}

var cacheKey string

func init() {
	CmdCache.AddCommand(CmdCacheClear)
	CmdCache.AddCommand(CmdCacheForget)
	// 设置 cache forget  的 命令选项值 --key 或 -k,
	// 示例 go run main.go cache forget --key=links:all
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	CmdCacheForget.MarkFlagRequired("key")
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}