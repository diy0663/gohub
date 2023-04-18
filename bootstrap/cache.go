package bootstrap

import (
	"fmt"

	"github.com/diy0663/gohub/pkg/cache"
	"github.com/diy0663/gohub/pkg/config"
)

func SetupCache() {
	cache.InitWithCacheStore(cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	))
}
