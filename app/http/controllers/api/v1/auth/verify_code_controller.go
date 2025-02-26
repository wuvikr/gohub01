package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	v1 "github.com/wuvikr/gohub01/app/http/controllers/api/v1"
	"github.com/wuvikr/gohub01/app/requests"
	"github.com/wuvikr/gohub01/pkg/captcha"
	"github.com/wuvikr/gohub01/pkg/logger"
	"github.com/wuvikr/gohub01/pkg/response"
	"github.com/wuvikr/gohub01/pkg/verifycode"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, _, err := captcha.NewCaptcha().GenerateCaptcha()
	if err != nil {
		logger.ErrorStr("captcha.NewCaptcha().GenerateCaptcha()", err.Error())
		response.Abort500(c, err)
	}

	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validator(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 2. 发送验证码
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, errors.New("发送短信失败"))
	}

	response.Success(c)
}

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validator(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	// 2. 发送验证码
	if err := verifycode.NewVerifyCode().SendEmail(request.Email); err != nil {
		response.Abort500(c, errors.New("发送 Email 验证码失败"))
		return
	}

	response.Success(c)
}
