package redis

import (
	"context"
	"sync"
	"time"

	"github.com/diy0663/gohub/pkg/logger"
	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// 定义一个全局变量并确保只实例化一次 搭配sync 包使用.
var once sync.Once

var Redis *RedisClient

// 里面使用了 once.Do , 确保只把全局redis对象实例化一次, 假如要实例化另外的redis对象用于别的用途,就后面的 NewClient 方法
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

func NewClient(address string, username string, password string, db int) *RedisClient {
	rds := &RedisClient{}
	rds.Context = context.Background()
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试redis 连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

// 把redis 底层包的方法  再封装一层
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

//实现 Set  Get  Has Del FlushDB Increment Decrement

func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	err := rds.Client.Set(rds.Context, key, value, expiration).Err()
	if err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// 判断可以是否存在 ,  内部错误和 redis.Nil 都返回 false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

func (rds RedisClient) Del(keys ...string) bool {
	err := rds.Client.Del(rds.Context, keys...).Err()
	if err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

func (rds RedisClient) FlushDB() bool {
	err := rds.Client.FlushDB(rds.Context).Err()
	if err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// // Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ,第二个参数为要增加的值 int64 类型。
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		//  强制类型转换
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)

		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}

	default:
		logger.ErrorString("Redis", "Increment", "参数过多")
		return false
	}

	return true
}

func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false
	}
	return true
}
