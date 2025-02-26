package etc

import "github.com/wuvikr/gohub01/pkg/config"

func init() {
	config.Add("redis", func() map[string]any {
		return map[string]any{
			"host":     config.MustGet[string]("REDIS_HOST"),
			"port":     config.MustGet[string]("REDIS_PORT"),
			"username": config.GetWithDefault("REDIS_USERNAME", ""),
			"password": config.GetWithDefault("REDIS_PASSWORD", ""),
			"db":       config.GetWithDefault("REDIS_DB", 0),
		}
	})
}
