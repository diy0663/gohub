package user

import (
	"gohub/app/models"

	"github.com/diy0663/go_project_packages/hash"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// 把明文密码 跟数据库里面加密的密码做处理
func (user *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, user.Password)
}
