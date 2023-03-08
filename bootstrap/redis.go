package bootstrap

import (
	"fmt"

	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/redis"
)

// 初始化redis
func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
