package captcha

import (
	"sync"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/redis"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}
		store := RedisStore{
			// redis.Redis 是仅被实例化一次的全局变量
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}
		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // 宽
			config.GetInt("captcha.width"),       // 高
			config.GetInt("captcha.length"),      // 长度
			config.GetFloat64("captcha.maxskew"), // 数字的最大倾斜角度
			config.GetInt("captcha.dotcount"),    // 图片背景里的混淆点数量
		)
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})
	return internalCaptcha
}

func (c *Captcha) Verify(id, answer string) (match bool) {
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	// todo , 这里为何不用store里面的Verify
	return c.Base64Captcha.Verify(id, answer, false)
}
