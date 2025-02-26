package verifycode

import (
	"fmt"
	"strings"
	"sync"

	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/helpers"
	"github.com/wuvikr/gohub01/pkg/mail"
	"github.com/wuvikr/gohub01/pkg/redis"
	"github.com/wuvikr/gohub01/pkg/sms"
)

type VerifyCode struct {
	Store Store
}

var (
	// 单例
	internalVerifyCode *VerifyCode
	once               sync.Once
)

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   fmt.Sprintf("%s:verifycode:", config.MustGet[string]("app.name")),
			},
		}
	})

	return internalVerifyCode
}

func (vc *VerifyCode) SendSMS(phone string) bool {
	// 1. 生成验证码
	code := vc.genVerifyCode(phone)

	// 1.1 如果不是生产环境，并且是特殊号码，跳过验证
	if config.MustGet[string]("app.env") != "production" && strings.HasPrefix(phone, config.MustGet[string]("verifycode.debug_phone_prefix")) {
		return true
	}

	// 2. 发送验证码
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.MustGet[string]("sms.aliyun.template_code"),
		Data: map[string]string{
			"code": code,
		},
	})
}

func (vc *VerifyCode) genVerifyCode(key string) string {
	// 1. 生成 6 位数字的验证码
	code, err := helpers.GenerateSecureNumber(config.MustGet[int]("verifycode.code_length"))
	if err != nil {
		return ""
	}
	// 1.1 如果不是生产环境，设置为固定验证码
	if config.MustGet[string]("app.env") != "production" {
		code = config.MustGet[string]("verifycode.debug_code")
	}
	// 2. 存储验证码
	if ok := vc.Store.Set(key, code); !ok {
		return ""
	}
	// 3. 返回验证码
	return code
}

func (vc *VerifyCode) Verify(key, answer string) bool {
	// 1. 如果不是生产环境，并且是特殊号码和邮箱，跳过验证
	if config.MustGet[string]("app.env") != "production" && (strings.HasPrefix(key, config.MustGet[string]("verifycode.debug_phone_prefix")) || strings.HasSuffix(key, config.MustGet[string]("verifycode.debug_email_suffix"))) {
		return true
	}
	// 2. 验证验证码
	return vc.Store.Verify(key, answer, false)
}

func (vc *VerifyCode) SendEmail(email string) error {
	// 1. 生成验证码
	code := vc.genVerifyCode(email)

	// 1.1 如果不是生产环境，并且是特殊邮箱，跳过验证
	if config.MustGet[string]("app.env") != "production" && strings.HasSuffix(email, config.MustGet[string]("verifycode.debug_email_suffix")) {
		return nil
	}

	// 2. 发送验证码
	mailData := mail.Email{
		From: mail.From{
			Name:    config.MustGet[string]("mail.from.name"),
			Address: config.MustGet[string]("mail.from.address"),
		},
		To:      []string{email},
		Subject: fmt.Sprintf("%s 验证码", config.MustGet[string]("app.name")),
		HTML:    fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code),
	}

	return mail.NewMailer().Send(mailData)
}
