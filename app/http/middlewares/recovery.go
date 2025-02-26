package middlewares

import (
	"errors"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wuvikr/gohub01/pkg/logger"
	"github.com/wuvikr/gohub01/pkg/response"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取用户的请求信息
				httpRequest, _ := httputil.DumpRequest(c.Request, true)

				// 链接中断，客户端中断连接为正常行为，不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)))

					c.Error(err.(error))
					c.Abort()
					// 如果连接已被关闭，我们无法写入状态码
					return
				}

				// 如果不是链接中断，就开始记录堆栈信息
				logger.Error("[Recovery from panic]",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.Stack("stacktrace"),
				)

				response.Abort500(c, errors.New("panic recovery"))
			}
		}()
		c.Next()
	}
}
