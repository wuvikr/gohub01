package bootstrap

import (
	"fmt"
	"time"

	"github.com/wuvikr/gohub01/app/models/user"

	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() {
	// 数据库连接配置
	var dbConfig gorm.Dialector

	v := config.MustGet[string]("database.type")
	switch v {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.MustGet[string]("database.mysql.username"),
			config.MustGet[string]("database.mysql.password"),
			config.MustGet[string]("database.mysql.host"),
			config.MustGet[string]("database.mysql.port"),
			config.MustGet[string]("database.mysql.database"),
			config.MustGet[string]("database.mysql.charset"),
		)
		fmt.Println("-------------: " + dsn)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		dbFile := config.MustGet[string]("database.sqlite.db_file")
		dbConfig = sqlite.Open(dbFile)
	default:
		panic("database type is not supported")
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.MustGet[int]("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.MustGet[int]("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.MustGet[int]("database.mysql.max_life_seconds")) * time.Second)

	database.DB.AutoMigrate(&user.User{})
}
