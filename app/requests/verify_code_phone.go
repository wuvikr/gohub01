package requests

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
	"github.com/wuvikr/gohub01/app/requests/validators"
)

type VerifyCodePhoneRequest struct {
	Phone string `json:"phone" valid:"phone"`

	CaptchaID     string `json:"captcha_id" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer" valid:"captcha_answer"`
}

func VerifyCodePhone(data any) url.Values {
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	message := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号长度必须为11位",
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
	_data := data.(*VerifyCodePhoneRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
