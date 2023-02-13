package user

import "github.com/diy0663/gohub/pkg/database"

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
