package bootstrap

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wuvikr/gohub01/app/http/middlewares"
	"github.com/wuvikr/gohub01/routes"
)

func SetupRoute(r *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(r)

	// 注册路由
	routes.RegisterAPIRoutes(r)

	// 处理404请求
	setup404Handler(r)
}

func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		// gin.Logger(),
		// gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
