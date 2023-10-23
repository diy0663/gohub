// Package link 模型
package link

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

// todo  需要手动保存之后 自动 import 进来,同时自行检查引入的包是否正确

type Link struct {
	models.BaseModel

	// Put fields in here
	Name string `json:"name,omitempty" gorm:"name,not null;type:varchar(255);" valid:"name"`
	Url  string `json:"url,omitempty" gorm:"url,not null;type:varchar(600);" valid:"url"`

	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

// 更新
func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
