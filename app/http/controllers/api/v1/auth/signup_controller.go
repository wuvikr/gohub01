package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wuvikr/gohub01/app/http/controllers/api/v1"
	"github.com/wuvikr/gohub01/app/models/user"
	"github.com/wuvikr/gohub01/app/requests"
	"github.com/wuvikr/gohub01/pkg/response"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validator(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{"exist": user.IsPhoneExist(request.Phone)})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 请求对象
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validator(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{"exist": user.IsEmailExist(request.Email)})
}
