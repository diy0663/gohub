package user

import "github.com/diy0663/gohub/pkg/database"

// 实用工具类, 特定查询

// 判断email 是否已经被注册过了
func IsEmailExists(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ? ", email).Count(&count)
	return count > 0
}

func IsPhoneExists(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone = ? ", phone).Count(&count)
	return count > 0
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone=?", phone).First(&userModel)
	return

}

func Get(idStr string) (userModel User) {
	database.DB.Where("id=?", idStr).First(&userModel)
	return
}

func GetByMulti(loginId string) (userModel User) {
	database.DB.Where("phone=?", loginId).Or("email=?", loginId).Or("name=?", loginId).First(&userModel)
	return
}
