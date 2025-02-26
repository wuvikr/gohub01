package bootstrap

import (
	"fmt"

	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%s:%s",
			config.MustGet[string]("redis.host"),
			config.MustGet[string]("redis.port"),
		),
		config.GetWithDefault[string]("redis.username", ""),
		config.GetWithDefault[string]("redis.password", ""),
		config.GetWithDefault[int]("redis.db", 0),
	)
}
