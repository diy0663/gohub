package user

import (
	"github.com/diy0663/go_project_packages/hash"
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/spf13/cast"
)

type User struct {
	models.BaseModel

	// 自身的数据表字段
	Name string `json:"name,omitempty" gorm:"name,not null;" valid:"name"`
	// 不希望将敏感信息输出给用户，所以这里 Email 、Phone 、Password 后面设置了 json:"-"
	Email    string `json:"-" gorm:"email,not null;index;" valid:"email"`
	Phone    string `json:"-" gorm:"phone,not null;index;" valid:"phone"`
	Password string `json:"-" gorm:"password,not null;" valid:"password"`
	RoleId   int64  `json:"-" gorm:"role_id,not null;index;" valid:"role_id"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// 密码对比
func (userModel *User) Comparepassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// Grom 里面的 Create 方法跟Save 方法有啥区别
// Create方法是用于创建新的记录，而Save方法是用于保存已有记录的变更。Create方法会在数据库中插入一条新的记录，
//	而Save方法会根据记录的主键（如果存在）进行更新或插入操作

// 根据model 的值更新表数据
func (userModel *User) Save() (rowAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

func (userModel *User) GetStringRoleID() string {
	return cast.ToString(userModel.RoleId)
}
