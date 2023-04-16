package link

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

type Link struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	//Name string `json:"name,omitempty"`
	// - 表示忽略,不输出此类敏感信息
	//Email    string `json:"-"`
	Name string `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	URL  string `gorm:"type:varchar(255);default:null" json:"url,omitempty"`

	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
