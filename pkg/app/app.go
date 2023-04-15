package app

import (
	"time"

	"github.com/diy0663/gohub/pkg/config"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}
func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// 根据配置的时区获取当前的时间 ,因为还要做计算,所以在这里的结果就不直接转 int64了
func TimenowInTimezone() time.Time {
	timezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(timezone)
}

func URL(path string) string {
	return config.Get("app.url") + path
}

func V1URL(path string) string {
	//return URL("/v1/" + path)
	return config.Get("app.url") + "/v1/" + path
}
