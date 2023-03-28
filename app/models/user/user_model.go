package user

import (
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/hash"
)

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowAffected int64) {
	// todo 这里为啥要传  &userModel ??
	result := database.DB.Save(&userModel)
	// 返回影响行数
	return result.RowsAffected
}
