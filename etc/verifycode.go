package etc

import "github.com/wuvikr/gohub01/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{

			// 验证码位数
			"code_length": config.GetWithDefault("VERIFY_CODE_LENGTH", 6),

			// 过期时间，单位是分钟
			"expire_time": config.GetWithDefault("VERIFY_CODE_EXPIRE", 15),

			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_time": 10080,
			// 本地开发环境使用此 debug_code，方便测试
			"debug_code": "666666",

			// 本地开发环境使用，方便测试
			"debug_phone_prefix": "000",
			"debug_email_suffix": "@testing.com",
		}
	})
}
