package auth

import (
	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type PasswordResetController struct {
	v1.BaseApiController
}

func (prc *PasswordResetController) ResetByPhone(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	userModel := user.GetByPone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		// todo 要是上面save不成功咋办??
		response.Success(c)
	}
}
