package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/wuvikr/gohub01/bootstrap"
	"github.com/wuvikr/gohub01/pkg/config"

	_ "github.com/wuvikr/gohub01/etc"
)

func main() {

	// 添加命令行参数 --env
	var env string
	flag.StringVar(&env, "env", "", "加载 env 配置，默认为当前目录下的.env文件，可以使用--env指定加载文件名称，如：--env=test 加载 .env.test 文件")
	flag.Parse()

	// 初始化配置
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化数据库
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 设置 gin 为发布模式
	gin.SetMode(gin.ReleaseMode)

	// 实例化 gin Engine
	r := gin.New()

	// 初始化路由
	bootstrap.SetupRoute(r)

	// 测试短信发送功能
	// sms.NewSMS().SendMessage("17717817438", sms.Message{
	// 	Template: config.MustGet[string]("sms.aliyun.template_code"),
	// 	Data:     map[string]string{"code": "48484"},
	// })

	// 运行服务
	err := r.Run(":" + config.MustGet[string]("app.port"))
	if err != nil {
		panic(err)
	}
}
