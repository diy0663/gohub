package captcha

import (
	"sync"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/diy0663/gohub/pkg/redis"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once
var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	// 单例模式, 为上面的  internalCaptcha 做初始化
	once.Do(func() {
		internalCaptcha = &Captcha{}

		store := RedisStore{
			RedisClient: redis.Redis, // 已经被初始化的全局单例redis对象
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}

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

// Generate generates a random id, base64 image string or an error if any
func (c *Captcha) Generate() (id, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// Verify by a given id key and remove the captcha value in store,
// return boolean value.
// if you has multiple captcha instances which share a same store.
// You may want to call `store.Verify` method instead.
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {

	logger.DebugString("验证对比", id, config.GetString("captcha.testing_key"))
	// 方便本地和 API 自动测试  , 传特定验证key 过来的直接允许跳过验证码验证
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}

	// 验证错了也不能直接清除.. 要允许多试几次
	return c.Base64Captcha.Verify(id, answer, false)
}
