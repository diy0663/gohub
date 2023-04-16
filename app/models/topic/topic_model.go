package topic

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/app/models/category"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/pkg/database"
)

type Topic struct {
	models.BaseModel

	// Put fields in here
	Title      string `gorm:"type:varchar(255);not null;index" json:"title,omitempty" `
	Body       string `gorm:"type:longtext;not null" json:"body,omitempty" `
	UserID     string `gorm:"type:bigint;not null;index" json:"user_id,omitempty"`
	CategoryID string `gorm:"type:bigint;not null;index" json:"category_id,omitempty"`

	//Name string `json:"name,omitempty"`
	// - 表示忽略,不输出此类敏感信息
	//Email    string `json:"-"`
	// 注意, 嵌入了另外两个表的struct后,使用gorm的自动迁移会生成对应的外键
	User     user.User         `json:"user"`
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
