package user

import (
	"github.com/diy0663/go_project_packages/hash"
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

type User struct {
	models.BaseModel

	// 自身的数据表字段
	Name string `json:"name,omitempty" gorm:"name,not null;" valid:"name"`
	// 不希望将敏感信息输出给用户，所以这里 Email 、Phone 、Password 后面设置了 json:"-"
	Email    string `json:"-" gorm:"email,not null;index;" valid:"email"`
	Phone    string `json:"-" gorm:"phone,not null;index;" valid:"phone"`
	Password string `json:"-" gorm:"password,not null;" valid:"password"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// 密码对比
func (userModel *User) Comparepassword(_password string) bool {
	return hash.BcryptCheck(userModel.Password, _password)
}
