package sms

import (
	"encoding/json"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"

	"github.com/wuvikr/gohub01/pkg/config"
	"github.com/wuvikr/gohub01/pkg/logger"
)

type Aliyun struct {
}

func (a Aliyun) Send(phone string, message Message) bool {
	// 从配置中获取阿里云短信配置
	accessKeyId := config.MustGet[string]("sms.aliyun.access_key_id")
	accessKeySecret := config.MustGet[string]("sms.aliyun.access_key_secret")
	signName := config.MustGet[string]("sms.aliyun.sign_name")

	logger.DebugJson("Aliyun.Send", "阿里云SMS配置信息", map[string]string{
		"accessKeyId":     accessKeyId,
		"accessKeySecret": accessKeySecret,
		"signName":        signName,
	})

	// 检查配置信息是否完整
	if len(accessKeyId) == 0 || len(accessKeySecret) == 0 || len(signName) == 0 {
		logger.ErrorStr("Aliyun.Send", "阿里云SMS配置信息不完整，请检查配置信息！")
		return false
	}

	// 创建阿里云短信客户端配置
	endpoint := "dysmsapi.aliyuncs.com"
	regionId := "cn-hangzhou"

	clientConfig := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
		RegionId:        &regionId,
	}

	// 创建阿里云短信客户端
	smsClient, err := dysmsapi.NewClient(clientConfig)
	if err != nil {
		logger.ErrorStr("dysmsapi.NewClient failed", err.Error())
		return false
	}

	// 构建request
	dataBytes, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorStr("json.Marshal failed", err.Error())
		return false
	}

	templateParam := string(dataBytes)
	smsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  &phone,
		SignName:      &signName,
		TemplateCode:  &message.Template,
		TemplateParam: &templateParam,
	}

	// 发送短信
	resp, err := smsClient.SendSms(smsRequest)
	if err != nil {
		logger.ErrorStr("smsClient.SendSms", err.Error())
		return false
	}

	if resp.Body == nil {
		logger.ErrorStr("smsClient.SendSms", "response body is nil")
		return false
	}

	if *resp.Body.Code != "OK" {
		logger.ErrorStr("smsClient.SendSms", *resp.Body.Message)
		return false
	}

	logger.DebugJson("Aliyun 短信发送成功", "response", resp)

	return true
}
