package helpers

import (
	"crypto/rand"
	"math/big"
)

// import "reflect"

// // Empty 判断是否为空
// func Empty(val any) bool {
// 	if val == nil {
// 		return true
// 	}

// 	v := reflect.ValueOf(val)
// 	switch v.Kind() {
// 	case reflect.String, reflect.Array:
// 		return v.Len() == 0
// 	case reflect.Map, reflect.Slice:
// 		return v.IsNil() || v.Len() == 0
// 	case reflect.Bool:
// 		return !v.Bool()
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		return v.Int() == 0
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
// 		return v.Uint() == 0
// 	case reflect.Float32, reflect.Float64:
// 		return v.Float() == 0
// 	case reflect.Interface, reflect.Ptr:
// 		return v.IsNil()
// 	}

// 	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
// }

// GenerateSecureNumber 生成密码学安全的随机数字字符串
func GenerateSecureNumber(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(10)) // 生成 0-9 的均匀分布
		if err != nil {
			return "", err
		}
		result[i] = '0' + byte(num.Int64()) // 直接转换为 ASCII 字符
	}
	return string(result), nil
}
