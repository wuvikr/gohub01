package requests

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone" valid:"phone"`
}

func SignupPhoneExist(data any) url.Values {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return validator(data, rules, messages)
}

type SignupEmailExistRequest struct {
	Email string `json:"email" valid:"email"`
}

func SignupEmailExist(data any) url.Values {
	rules := govalidator.MapData{
		"email": []string{"required", "email", "min:4", "max:30"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:email为必填项",
			"email:email格式错误，参数名称 email",
			"min:email长度需大于 4",
			"max:email长度需小于 30",
		},
	}

	return validator(data, rules, messages)
}
