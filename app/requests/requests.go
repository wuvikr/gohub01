package requests

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"github.com/wuvikr/gohub01/pkg/response"
)

func validator(data any, rules govalidator.MapData, messages govalidator.MapData) url.Values {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).ValidateStruct()
}

type validatorFunc func(any) url.Values

func Validator(c *gin.Context, request any, validator validatorFunc) bool {
	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		response.BadRequest(c, err)
		fmt.Println(err.Error())
		return false
	}

	// 表单验证
	errs := validator(request)

	if len(errs) > 0 {
		response.ValidationFailed(c, errs)
		return false
	}

	return true
}
