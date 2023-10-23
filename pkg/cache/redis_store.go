package cache

import (
	"time"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/redis"
)

// 先定义结构体,之后利用结构体去实现 interface
type RedisStore struct {
	// 客户端链接
	RedisClient *redis.RedisClient
	// 缓存整体前缀
	KeyPrefix string
}

// 先在interface 文件里面,把光标移动到最底下空白行腾出空间
// 使用vscode 插件 shift +cmd + p  选择 Go: Generate Interface Stubs , 按格式输入: redisStore *RedisStore Store
// 就能让 RedisStore 去实现 接口 Store, 之后再把  RedisStore 以及生成的方法都移动到别的文件去单独放

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.KeyPrefix = config.GetString("app.name") + ":cache:"
	rs.RedisClient = redis.NewClient(address, username, password, db)
	return rs
}

func (redisStore *RedisStore) Set(key string, value string, expireTime time.Duration) {
	redisStore.RedisClient.Set(redisStore.KeyPrefix+key, value, expireTime)
}

func (redisStore *RedisStore) Get(key string) string {

	return redisStore.RedisClient.Get(redisStore.KeyPrefix + key)
}

func (redisStore *RedisStore) Has(key string) bool {
	return redisStore.RedisClient.Has(redisStore.KeyPrefix + key)
}

func (redisStore *RedisStore) Forget(key string) {
	redisStore.RedisClient.Del(redisStore.KeyPrefix + key)
}

// 永久缓存
func (redisStore *RedisStore) Forever(key string, value string) {
	redisStore.RedisClient.Set(redisStore.KeyPrefix+key, value, 0)

}

func (redisStore *RedisStore) Flush() {
	redisStore.RedisClient.FlushDB()
}

func (redisStore *RedisStore) IsAlive() error {
	return redisStore.RedisClient.Ping()
}

// Increment 当参数只有 1 个时，为 key，增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
func (redisStore *RedisStore) Increment(parameters ...interface{}) {
	redisStore.RedisClient.Increment(parameters...)
}

// Decrement 当参数只有 1 个时，为 key，减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (redisStore *RedisStore) Decrement(parameters ...interface{}) {
	redisStore.RedisClient.Decrement(parameters...)
}
