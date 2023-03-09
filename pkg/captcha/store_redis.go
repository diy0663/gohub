package captcha

import (
	"errors"
	"time"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/redis"
)

// 实现基于redis 的验证码存储 b
// 实现下面的这个 interface 即可
/*

type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string) error

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string

	//Verify captcha's answer directly
	Verify(id, answer string, clear bool) bool
}

*/

type RedisStore struct {
	RedisClient *redis.RedisClient
	// 特定前缀, 凡是验证码相关的用到的redis都加上这个
	KeyPrefix string
}

func (s *RedisStore) Set(key string, value string) error {
	// 过期时间通过配置文件来设置 , 这里按照分钟来作为单位
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_minute"))
	// 新增一个本地调试用的专用的过期时间
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_minute"))
	}
	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

func (s *RedisStore) Get(key string, clear bool) string {
	value := s.RedisClient.Get(s.KeyPrefix + key)
	if clear {
		// 这个操作有个返回值, 可以不接收??
		_ = s.RedisClient.Del(s.KeyPrefix + key)
	}
	return value
}

func (s *RedisStore) Verify(key string, answer string, clear bool) bool {
	value := s.Get(key, clear)
	return value == answer
}
