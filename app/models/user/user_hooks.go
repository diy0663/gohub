package user

import (
	"github.com/diy0663/go_project_packages/hash"
	"gorm.io/gorm"
)

func (user *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(user.Password) {
		user.Password = hash.BcryptHash(user.Password)
	}
	return
}
