package user

import "github.com/diy0663/gohub/pkg/database"

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
