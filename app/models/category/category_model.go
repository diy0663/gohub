// Package category 模型
package category

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

//   需要手动保存之后 自动 import 进来,同时自行检查引入的包是否正确

type Category struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name,omitempty" gorm:"name,type:varchar(255);not null;index;" valid:"name"`
	Description string `json:"description,omitempty" gorm:"description,type:varchar(255);default:''" valid:"description"`

	models.CommonTimestampsField
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

// 更新
func (category *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}
