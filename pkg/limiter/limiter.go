package limiter

import (
	"strings"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/go_project_packages/redis"
	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	// 初始化存储，使用我们程序里共用的 redis.Redis 对象
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiterlib.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	limiterObj := limiterlib.New(store, rate)
	// 获取限流的结果
	if c.GetBool("limiter-once") {
		// Peek() 取结果，不增加访问次数
		return limiterObj.Peek(c, key)
	} else {
		// 确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问次数。
		c.Set("limiter-once", true)
		// Get() 取结果且增加访问次数
		return limiterObj.Get(c, key)
	}

}

// 将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName

}
