package app

import (
	"time"

	"github.com/diy0663/go_project_packages/config"
)

// 当前时间
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}
