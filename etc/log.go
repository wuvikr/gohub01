package etc

import "github.com/wuvikr/gohub01/pkg/config"

func init() {
	config.Add("log", func() map[string]any {
		return map[string]any{
			// 日志级别
			"level": config.GetWithDefault("LOG_LEVEL", "debug"),

			// 日志类型，可选：
			// single 单文件
			// daily 按日期，每天一个日志文件
			"type": config.GetWithDefault("LOG_TYPE", "single"),

			// 日志文件路径
			"filename": config.GetWithDefault("LOG_NAME", "storage/logs/logs.log"),

			// 日志文件名最大 Size, 单位：M
			"max_size": config.GetWithDefault("LOG_MAX_SIZE", 64),
			// 日志文件最多保存多少个备份
			"max_backup": config.GetWithDefault("LOG_MAX_BACKUP", 5),
			// 日志文件最多保存多少天
			"max_age": config.GetWithDefault("LOG_MAX_AGE", 30),
			// 是否压缩日志文件，默认为 false 不压缩
			"compress": config.GetWithDefault("LOG_COMPRESS", false),
		}
	})
}
