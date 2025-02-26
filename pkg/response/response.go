package response

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// JSON 响应 200 和 JSON 数据
func JSON(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// 执行某些没有具体返回数据的变更操作成功后调用
func Success(c *gin.Context) {
	JSON(c, gin.H{
		"success": true,
		"message": "操作成功",
	})
}

// Error 响应错误
func Error(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "操作失败",
		"error":   err.Error(),
	})
}

func Abort500(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "服务器内部错误，请稍后再试",
		"error":   err.Error(),
	})
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		"errors":  err.Error(),
	})
}

func ValidationFailed(c *gin.Context, errors url.Values) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "请求验证不通过，具体请查看 errors",
		"errors":  errors,
	})
}
