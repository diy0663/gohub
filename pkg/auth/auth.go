package auth

import (
	"errors"

	"github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/gin-gonic/gin"
)

// 登录判断 ,传进来的得确保是加密的密码?
func Attempt(loginId string, password string) (user.User, error) {
	// 先查用户是否存在
	userModel := user.GetByMulti(loginId)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	// 验证密码
	if !userModel.Comparepassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

// 获取当前用户

func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel

}

// 获取当前用户id
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
