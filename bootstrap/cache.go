package bootstrap

import (
	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/gohub/pkg/cache"
)

func SetCache() {
	// 初始化全局变量 缓存服务, 这个缓存服务用的是redis 作为 store 存储
	rds := cache.NewRedisStore(
		config.GetString("redis.host")+":"+config.GetString("redis.port"),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"))
	cache.NewCacheService(rds)

}
