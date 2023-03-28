package user

import "github.com/diy0663/gohub/pkg/database"

// user_util 里面的函数,都是简单传参

// 通过查表来判断email 是否已存在
func IsEmailExists(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email=?", email).Count(&count)
	return count > 0

}

func IsPhoneExists(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ? ", phone).Count(&count)
	return count > 0

}

func GetByMulti(loginId string) (userModel User) {
	database.DB.Where("phone= ? ", loginId).Or("email = ? ", loginId).Or("name = ? ", loginId).First(&userModel)
	return userModel
}

func GetByPone(phone string) (userModel User) {
	database.DB.Where("phone = ? ", phone).First(&userModel)
	return userModel

}

func Get(id string) (userModel User) {
	database.DB.Where("id = ? ", id).First(&userModel)
	return userModel
}
