package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

func ErrorStr(moduleName, message string) {
	Logger.Error(moduleName, zap.String("msg", message))
}

func json2Str(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		Logger.Error("json.Marshal failed", zap.Error(err))
		return "", err
	}

	return string(b), nil
}

// DebugJson 将对象类型转换为 json 字符串并记录信息日志
// 使用 json.Marshal 进行转换，调用示例：DebugJson("getuser", "用户信息", user)
func DebugJson(moduleName string, key string, value any) {
	str, err := json2Str(value)
	if err != nil {
		return
	}
	Logger.Debug(moduleName, zap.String(key, str))
}

func InfoJson(moduleName string, key string, value any) {
	str, err := json2Str(value)
	if err != nil {
		return
	}
	Logger.Info(moduleName, zap.String(key, str))
}

func WarnJson(moduleName string, key string, value any) {
	str, err := json2Str(value)
	if err != nil {
		return
	}
	Logger.Warn(moduleName, zap.String(key, str))
}

func ErrorJson(moduleName string, key string, value any) {
	str, err := json2Str(value)
	if err != nil {
		return
	}
	Logger.Error(moduleName, zap.String(key, str))
}
