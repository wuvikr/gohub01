package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuvikr/gohub01/app/http/controllers/api/v1/auth"
)

func RegisterAPIRoutes(r gin.IRouter) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 邮箱是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			vcc := new(auth.VerifyCodeController)
			// 验证码
			authGroup.GET("/verify-codes/captcha", vcc.ShowCaptcha)
			// 发送验证码
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			// 发送验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
		}
	}
}
