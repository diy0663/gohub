package auth

import (
	"errors"

	"github.com/diy0663/gohub/app/models/user"
)

// 登录,授权验证登, 基于userModel 做的处理,不具备通用性,其实不应该放pkg模块

// 基于用户名/邮箱/手机号 + 密码尝试登录, 注册时候 用户名/邮箱/手机号 得经过not_exists 规则校验
func Attempt(loginId string, password string) (user.User, error) {
	userModel := user.GetByMulti(loginId)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}
