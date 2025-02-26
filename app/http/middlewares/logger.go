package middlewares

import (
	"bytes"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wuvikr/gohub01/pkg/logger"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 response 内容
		w := &responseBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 由于读取后 body 会被清空，需要重新设置 body 供后续中间件使用
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		start := time.Now()
		c.Next()

		// 计算耗时
		cost := time.Since(start)
		responStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", cost.String()),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			logFields = append(logFields, zap.String("request_body", string(requestBody)))

			logFields = append(logFields, zap.String("response_body", w.body.String()))
		}

		if responStatus > 400 && responStatus <= 499 {
			logger.Warn("HTTP WARNING"+strconv.Itoa(responStatus), logFields...)
		}
		if responStatus >= 500 && responStatus <= 599 {
			logger.Error("HTTP ERROR"+strconv.Itoa(responStatus), logFields...)
		}
		logger.Debug("HTTP ACCESS LOG", logFields...)
	}
}
