// Package role 模型
package role

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

// todo  需要手动保存之后 自动 import 进来,同时自行检查引入的包是否正确

type Role struct {
	models.BaseModel

	// Put fields in here
	Name string `json:"name,omitempty" gorm:"name,not null;type:varchar(30);uniqueIndex;" valid:"name"`

	models.CommonTimestampsField
}

func (role *Role) Create() {
	database.DB.Create(&role)
}

// 更新
func (role *Role) Save() (rowsAffected int64) {
	result := database.DB.Save(&role)
	return result.RowsAffected
}

func (role *Role) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&role)
	return result.RowsAffected
}
