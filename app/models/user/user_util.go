package user

import "github.com/diy0663/gohub/pkg/database"

// 存放user 模型的相关操作方法

// 新增用户的时候检查email是否已存在,避免重复注册
func IsEmailExist(email string) bool {

	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	// todo 假如sql报错咋办
	return count > 0
}

// 判断手机号是否已注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	// todo 假如sql报错咋办
	return count > 0
}
