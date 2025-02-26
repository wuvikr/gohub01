package requests

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
	"github.com/wuvikr/gohub01/pkg/captcha"
)

type VerifyCodeEmailRequest struct {
	Email string `json:"email" valid:"email"`

	CaptchaID     string `json:"captcha_id" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer" valid:"captcha_answer"`
}

func VerifyCodeEmail(data any) url.Values {
	rules := govalidator.MapData{
		"email":          []string{"required", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	message := govalidator.MapData{
		"email": []string{
			"required:邮箱不能为空",
			"email:邮箱格式不正确",
		},
		"captcha_id": []string{
			"required:验证码ID不能为空",
		},
		"captcha_answer": []string{
			"required:验证码不能为空",
			"digits:验证码长度必须为6位",
		},
	}

	errs := validator(data, rules, message)

	// 图片验证码
	_data := data.(*VerifyCodeEmailRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}
