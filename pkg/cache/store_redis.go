package cache

import (
	"time"

	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/redis"
)

type RedisStore struct {
	RedisClient redis.RedisClient
	KeyPrefix   string
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	return &RedisStore{
		RedisClient: *redis.NewClient(address, username, password, db),
		KeyPrefix:   config.GetString("app.name") + ":cache:",
	}

}

// 挨个实现 Store 这个interface
func (s *RedisStore) Set(key string, value string, expireTime time.Duration) bool {
	return s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)

}
func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}
func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

// 删除缓存key
func (s *RedisStore) Forget(key string) bool {
	return s.RedisClient.Del(s.KeyPrefix + key)
}

// 让缓存过期
func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}
func (s *RedisStore) Flush() bool {
	return s.RedisClient.FlushDB()
}
func (s *RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}

func (s *RedisStore) Increment(parameters ...interface{}) bool {
	return s.RedisClient.Increment(parameters...)
}
func (s *RedisStore) Decrement(parameters ...interface{}) bool {
	return s.RedisClient.Decrement(parameters...)
}
