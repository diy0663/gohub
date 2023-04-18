package cache

import "time"

// 每一种缓存(redis /memcached 等) 都必须实现的接口/功能
type Store interface {
	// Set 这里用不到 bool 类型的返回值? 如何知道是否设置成功呢??
	Set(key string, value string, expireTime time.Duration) bool
	Get(key string) string
	Has(key string) bool
	// 删除缓存key
	Forget(key string) bool
	// 让缓存过期
	Forever(key string, value string)
	// 清空缓存
	Flush() bool

	IsAlive() error

	// Increment 当参数只有 1 个时，为 key，增加 1。
	// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
	Increment(parameters ...interface{}) bool
	Decrement(parameters ...interface{}) bool
}
