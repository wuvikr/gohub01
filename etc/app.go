package etc

import "github.com/wuvikr/gohub01/pkg/config"

func init() {
	config.Add("app", func() map[string]any {
		return map[string]any{
			// 应用名称
			"name": config.GetWithDefault("APP_NAME", "gohub01"),

			// 应用环境, 一般分为 本地开发环境, 测试环境, 生产环境
			"env": config.GetWithDefault("APP_ENV", "production"),

			// 是否开启应用调试模式
			"debug": config.GetWithDefault("APP_DEBUG", false),

			// 应用加密key
			"key": config.GetWithDefault("APP_KEY", "22222222222"),

			// 服务url
			"url": config.GetWithDefault("APP_URL", "http://localhost:3000"),

			// 服务端口
			"port": config.GetWithDefault("APP_PORT", "3000"),

			// 时区
			"timezone": config.GetWithDefault("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
