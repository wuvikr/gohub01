package bootstrap

import (
	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/logger"
)

func SetupLogger() {
	// 初始化 Logger
	logger.InitLogger(
		config.MustGet[string]("log.filename"),
		config.MustGet[int]("log.max_size"),
		config.MustGet[int]("log.max_backup"),
		config.MustGet[int]("log.max_age"),
		config.MustGet[bool]("log.compress"),
		config.MustGet[string]("log.type"),
		config.MustGet[string]("log.level"),
	)
}
