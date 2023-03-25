package auth

import (
	"net/http"

	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExists(c *gin.Context) {

	//panic("故意出错")
	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExists(request.Phone),
	})

}

func (sc *SignupController) IsEmailExists(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	// 解析请求
	ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist)
	// 根据验证规则校验字段数据  todo &request
	if !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExists(request.Email),
	})

}
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	ok := requests.Validate(c, &request, requests.SignupUsingPhone)
	if !ok {
		return
	}
	_user := user.User{

		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		// 注册成功之后生成token 返回
		// todo
		token := jwt.NewJWt().IssueToken(_user.GetStringId(), _user.Name)
		// 这里哪怕生成的token是空串也暂时不管了
		response.CreatedJSON(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试")
	}

}

func (sc *SignupController) SignupUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingEmailRequest{}
	ok := requests.Validate(c, &request, requests.SignupUsingEmail)
	if !ok {
		return
	}
	_user := user.User{

		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试")
	}

}
