package cache

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/diy0663/go_project_packages/logger"
	"github.com/spf13/cast"
)

type CacheService struct {
	// 这里放接口, 实例化的时候把 实现了这个接口的 struct 传进去
	Store Store
}

// 使用单例模式
var once sync.Once

// 使用全局变量 , 这个全局变量的赋值,是在bootstrap 那边调用这里面的 NewCacheService 来得到的
var Cache *CacheService

// 根据传入的store 类型实例化对应的缓存类型
func NewCacheService(store Store) {
	once.Do(func() {
		//对全局变量进行赋值
		Cache = &CacheService{
			Store: store,
		}
	})
}

// 这里针对缓存, 用的是函数,而不是方法 , 因为包名比较容易记住,便于使用
// 因为 函数里面都是针对全局变量的操作, 调用方式为 包名.函数() 全局变量在最开始就已经初始化了
// 假如要使用方法, 也可以,那么调用方式就得变为 全局变量.方法()

// 这里用了序列化
func Set(key string, obj interface{}, expireTime time.Duration) {
	// 传 &
	// 参数传入一个指针地址的作用是为了将数据序列化为JSON格式的字节流。JSON编码器将通过指针获取数据的值，并将该值转换为JSON格式的字节流。
	// 通过传递指针地址，可以避免数据的复制
	// 通过传递指针地址，可以减少内存使用
	byteSlice, err := json.Marshal(&obj)
	logger.LogIf(err)
	// 使用全局变量做操作
	Cache.Store.Set(key, string(byteSlice), expireTime)
}
func Get(key string) interface{} {

	stringValue := Cache.Store.Get(key)
	var want interface{}
	// 反序列化
	err := json.Unmarshal([]byte(stringValue), &want)
	logger.LogIf(err)
	return want
}

func Has(key string) bool {
	return Cache.Store.Has(key)
}

// GetObject 应该传地址，因为下面定义的是没返回值的,所以只能在传参做手脚
// 用法如下:
//
//	model := user.User{}
//	cache.GetObject("key", &model)
func GetObject(key string, wanted interface{}) {
	val := Cache.Store.Get(key)
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &wanted)
		logger.LogIf(err)
	}
}

func GetString(key string) string {
	return cast.ToString(Get(key))
}

func GetBool(key string) bool {
	return cast.ToBool(Get(key))
}

// 下面是一堆基于 cast 获取指定数据类型的缓存内容
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
