package middlewares

import (
	"net/http"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/pkg/limiter"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 格式
//
// * 5 reqs/second: "5-S"   每秒 5次
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
//

func LimitIP(limitFormat string) gin.HandlerFunc {
	if config.GetString("app.env") == "test" {
		// 测试环境调大限流量
		limitFormat = "10000000-H"
	}
	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limitFormat); !ok {
			return
		}
		c.Next()
		//if ok:=
	}
}

// IP+路由 作为组合一起限制
func LimitPerRoute(limitFormat string) gin.HandlerFunc {
	if config.GetString("app.env") == "test" {
		// 测试环境调大限流量
		limitFormat = "10000000-H"
	}
	return func(c *gin.Context) {
		c.Set("limiter-once", false)
		// limiter-once 最终还会在  limitHandler-> limiter.CheckRate 里面用到
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limitFormat); !ok {
			return
		}

		c.Next()
	}
}

// 检查某个key 的访问是否达到了指定的限量
func limitHandler(c *gin.Context, key string, limitFormat string) bool {
	// 使用 CheckRate ,里面构造并存放了限流实例
	rate, err := limiter.CheckRate(c, key, limitFormat)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}
	// 本限量 总共数量多少
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	// 还剩下多少量
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	// 啥时间戳之后会重置限量
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))
	if rate.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁",
		})
		return false
	}
	return true
}
