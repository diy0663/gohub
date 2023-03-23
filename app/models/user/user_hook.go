package user

import (
	hash "github.com/diy0663/gohub/pkg/hash"

	"gorm.io/gorm"
)

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
