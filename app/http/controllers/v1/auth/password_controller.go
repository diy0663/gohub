package auth

import (
	"github.com/diy0663/go_project_packages/response"
	v1 "github.com/diy0663/gohub/app/http/controllers/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	// 表单验证
	request := requests.ResetByEmailRequest{}
	if ok := requests.RequestValidate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	//更新密码
	userModel := user.GetByMulti(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c, "查无该用户")
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}

}
