package etc

import "github.com/wuvikr/gohub01/pkg/config"

func init() {
	config.Add("sms", func() map[string]any {
		return map[string]any{
			// 阿里云短信配置
			"aliyun": map[string]any{
				"access_key_id":     config.GetWithDefault("SMS_ALIYUN_ACCESS_ID", ""),
				"access_key_secret": config.GetWithDefault("SMS_ALIYUN_ACCESS_SECRET", ""),
				"sign_name":         config.GetWithDefault("SMS_ALIYUN_SIGN_NAME", "阿里云短信测试"),
				"template_code":     config.GetWithDefault("SMS_ALIYUN_TEMPLATE_CODE", "SMS_154950909"),
			},
		}
	})
}
