// Package topic 模型
package topic

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/app/models/category"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/pkg/database"
)

// todo  需要手动保存之后 自动 import 进来,同时自行检查引入的包是否正确

type Topic struct {
	models.BaseModel

	// Put fields in here,

	Title string `json:"title,omitempty" gorm:"title,not null;" valid:"title"`
	Body  string `json:"body,omitempty" gorm:"body,not null;" valid:"body"`

	UserID     string `json:"user_id,omitempty" gorm:"user_id,not null;" valid:"user_id"`
	CategoryID string `json:"category_id,omitempty" gorm:"category_id,not null;" valid:"category_id"`

	// todo 嵌入其他model的时候, 会自动创建外键 , user_id 和 category_id 是外键名
	User     user.User         `json:"user,omitempty" `
	Category category.Category `json:"category,omitempty" `

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

// 更新
func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
