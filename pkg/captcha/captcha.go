package captcha

import (
	"fmt"
	"sync"

	"github.com/mojocn/base64Captcha"
	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/redis"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// 单例
var once sync.Once
var internalCaptcha *Captcha

// NewCaptcha 单例模式获取
func NewCaptcha() *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用 redis 实例初始化 store
		store := &RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   fmt.Sprintf("%s:captcha:", config.MustGet[string]("app.name")),
		}

		// 配置 base64Captcha 驱动
		driver := base64Captcha.NewDriverDigit(
			config.MustGet[int]("captcha.height"),
			config.MustGet[int]("captcha.width"),
			config.MustGet[int]("captcha.length"),
			config.MustGet[float64]("captcha.maxskew"),
			config.MustGet[int]("captcha.dotcount"),
		)

		// 实例化 base64Captcha
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, store)
	})

	return internalCaptcha
}

// GenerateCaptcha 生成验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, answer string, err error) {
	// 生成验证码，Generate 内部已经做了 err 处理，直接返回即可，返回 err 在调用逻辑中处理
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) bool {
	// 非生产环境，跳过验证
	if config.MustGet[string]("app.env") != "production" && id == config.MustGet[string]("captcha.testing_key") {
		return true
	}
	// 这里不删除验证码，让验证码自动过期
	// 这样用户多次提交表单，验证码都有效，防止表单提交错误需要多次刷新验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
