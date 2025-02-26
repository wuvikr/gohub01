package etc

import "github.com/wuvikr/gohub01/pkg/config"

func init() {
	config.Add("database", func() map[string]any {
		return map[string]any{
			// 数据库连接信息
			"type": config.GetWithDefault("DB_TYPE", "mysql"),

			"mysql": map[string]any{
				"host":     config.GetWithDefault("DB_HOST", "127.0.0.1"),
				"port":     config.GetWithDefault("DB_PORT", "3306"),
				"database": config.GetWithDefault("DB_DATABASE", "gohub01"),
				"username": config.GetWithDefault("DB_USERNAME", "root"),
				"password": config.GetWithDefault("DB_PASSWORD", "Admin#098"),
				"charset":  "utf8mb4",
				// 连接池配置
				"max_idle_connections": config.GetWithDefault("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": config.GetWithDefault("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds":     config.GetWithDefault("DB_MAX_LIFE_SECONDS", 5*60),
			},

			"sqlite": map[string]any{
				"db_file": config.GetWithDefault("DB_SQL_FILE", "database/gohub01.db"),
			},
		}
	})
}
