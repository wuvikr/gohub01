package validators

import (
	"net/url"

	"github.com/wuvikr/gohub01/pkg/captcha"
)

func ValidateCaptcha(captchaID, captchaAnswer string, errs url.Values) url.Values {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}
