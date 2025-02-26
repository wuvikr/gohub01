package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cast"
	viperLib "github.com/spf13/viper"
)

// viper 实例
var viper *viperLib.Viper

// ConfigFunc 配置函数
type ConfigFunc func() map[string]any

// ConfigFuncs 配置函数集合
var ConfigFuncs map[string]ConfigFunc

func init() {

	// 1. 初始化 viper
	viper = viperLib.New()

	viper.SetConfigType("env")

	viper.AddConfigPath(".")

	viper.SetEnvPrefix("appenv")

	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)

	// 2. 注册配置信息
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	// 1. 默认加载.env文件, 如果有传参 --env=name 则加载.env.name文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := fmt.Sprintf(".env.%s", envSuffix)
		if _, err := os.Stat(filePath); err != nil {
			panic(fmt.Sprintf("配置文件%s不存在", filePath))
		}
		envPath = filePath
	}

	// 2. 加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("加载配置文件失败: %s", err.Error()))
	}

	// 3. 监控配置文件变化
	viper.WatchConfig()

}

// Add 添加配置信息
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// ConfigValue 是可以从配置中获取的类型
type ConfigValue interface {
	~string | ~int | ~int64 | ~float64 | ~bool | ~map[string]string
}

// Get 获取配置项，返回值和是否成功
func Get[T ConfigValue](path string) (T, bool) {
	// 定义一个泛型默认值，如果没有取到值则返回零值和 false
	var fallback T

	value := viper.Get(path)
	if value == nil {
		return fallback, false
	}

	// 获取目标类型
	targetType := reflect.TypeOf(fallback)
	valueType := reflect.TypeOf(value)

	// 如果类型已经匹配，直接返回
	if valueType == targetType {
		return value.(T), true
	}

	// 尝试类型转换
	switch any(fallback).(type) {
	case string:
		// 转换为字符串
		strValue := cast.ToString(value)
		return any(strValue).(T), true
	case int:
		// 转换为整数
		intValue := cast.ToInt(value)
		return any(intValue).(T), true
	case int64:
		// 转换为整数
		int64Value := cast.ToInt64(value)
		return any(int64Value).(T), true
	case bool:
		// 转换为布尔值
		boolValue := cast.ToBool(value)
		return any(boolValue).(T), true
	case float64:
		// 转换为浮点数
		float64Value := cast.ToFloat64(value)
		return any(float64Value).(T), true
	case map[string]string:
		// 转换为字符串映射
		stringMapValue := cast.ToStringMapString(value)
		return any(stringMapValue).(T), true
	default:
		return fallback, false
	}
}

// GetWithDefault 获取配置项，如果不存在则返回默认值
func GetWithDefault[T ConfigValue](path string, defaultValue T) T {
	value, ok := Get[T](path)
	if !ok {
		return defaultValue
	}
	return value
}

// MustGet 获取配置项，如果失败则panic
func MustGet[T ConfigValue](path string) T {
	value, ok := Get[T](path)
	if !ok {
		panic(fmt.Sprintf("配置项 [%s] 不存在或类型转换失败", path))
	}
	return value
}
