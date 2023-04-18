package cache

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/diy0663/gohub/pkg/logger"
	"github.com/spf13/cast"
)

type CacheService struct {
	Store Store
}

var once sync.Once

// 单例变量
var Cache *CacheService

func InitWithCacheStore(store Store) {
	once.Do(func() {
		Cache = &CacheService{Store: store}
	})
}

// 缓存(redis/memcache)底层存的都是字符串
// 缓存服务这里, 存任意的数据结果的Json格式, 取数据的时候再还原即可

func Set(key string, object interface{}, expireTime time.Duration) {
	value, err := json.Marshal(object)
	logger.LogIf(err)
	Cache.Store.Set(key, string(value), expireTime)
}

func Get(key string) interface{} {
	stringValue := Cache.Store.Get(key)
	var wanted interface{}
	err := json.Unmarshal(([]byte(stringValue)), &wanted)
	logger.LogIf(err)
	return wanted
}

func Has(key string) bool {
	return Cache.Store.Has(key)
}
func Forget(key string) {
	Cache.Store.Forget(key)
}

func Forever(key string, value string) {
	Cache.Store.Set(key, value, 0)
}

func Flush() {
	Cache.Store.Flush()
}

func Increment(parameters ...interface{}) {
	Cache.Store.Increment(parameters...)
}

func Decrement(parameters ...interface{}) {
	Cache.Store.Decrement(parameters...)
}

func IsAlive() error {
	return Cache.Store.IsAlive()
}

// GetObject 应该传地址变量 &，用法如下:

// model := user.User{}
// cache.GetObject("key", &model)
func GetObject(key string, wanted interface{}) {
	val := Cache.Store.Get(key)
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &wanted)
		logger.LogIf(err)
	}
}

// 获取数据的同时配合cast包, 获取自己想要的数据格式
func GetString(key string) string {
	return cast.ToString(Get(key))
}

func GetBool(key string) bool {
	return cast.ToBool(Get(key))
}

func GetInt(key string) int {
	return cast.ToInt(Get(key))
}

func GetInt32(key string) int32 {
	return cast.ToInt32(Get(key))
}

func GetInt64(key string) int64 {
	return cast.ToInt64(Get(key))
}

func GetUint(key string) uint {
	return cast.ToUint(Get(key))
}

func GetUint32(key string) uint32 {
	return cast.ToUint32(Get(key))
}

func GetUint64(key string) uint64 {
	return cast.ToUint64(Get(key))
}

func GetFloat64(key string) float64 {
	return cast.ToFloat64(Get(key))
}

func GetTime(key string) time.Time {
	return cast.ToTime(Get(key))
}

func GetDuration(key string) time.Duration {
	return cast.ToDuration(Get(key))
}

func GetIntSlice(key string) []int {
	return cast.ToIntSlice(Get(key))
}

func GetStringSlice(key string) []string {
	return cast.ToStringSlice(Get(key))
}

func GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(Get(key))
}

func GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(Get(key))
}

func GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(Get(key))
}
