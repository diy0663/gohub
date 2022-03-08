package captcha

import (
	"errors"
	"time"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/redis"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) error {
	// 验证码过期时间(分钟)
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码到redis")
	}
	return nil
}

// 第二个参数的意思是取完就顺便删除
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val

}

// 验证用户输入的验证码是否正确
func (s *RedisStore) Verify(key string, user_input string, clear bool) bool {

	value := s.Get(key, clear)
	return value == user_input
}
