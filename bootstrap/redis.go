package bootstrap

import (
	"fmt"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/redis"
)

func SetupRedis() {

	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
