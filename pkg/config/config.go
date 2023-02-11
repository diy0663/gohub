package config

import (
	"os"
	"strings"

	"github.com/diy0663/gohub/pkg/helper"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

// 自动调用初始化
func init() {
	viper = viperlib.New()

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("appenv")
	viper.AutomaticEnv()
	ConfigFuncs = make(map[string]ConfigFunc)
}

// 被外部调用
func InitConfig(env string) {
	loadEnv(env)
	loadConfig()

}

func loadEnv(envSuffix string) {
	// 找到对应指定环境的配置文件进行加载
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := envPath + "." + strings.TrimSpace(envSuffix)
		if _, err := os.Stat(filePath); err == nil {
			envPath = filePath
		} else {
			// 没有这个文件就报错
			panic(err)
		}
	}
	viper.SetConfigName(envPath)
	// 文件读取有错也报错
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.WatchConfig()

}

func loadConfig() {
	// viper 是本包内的全局变量
	// ConfigFuncs 是上面最开始定义的全局变量
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helper.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
