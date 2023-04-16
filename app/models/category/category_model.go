package category

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

type Category struct {
	models.BaseModel

	// Put fields in here
	Name        string `gorm:"type:varchar(255);not null;index" json:"name,omitempty"`
	Description string `gorm:"type:varchar(255);default:null" json:"description,omitempty"`
	//Name string `json:"name,omitempty"`
	// - 表示忽略,不输出此类敏感信息
	//Email    string `json:"-"`

	models.CommonTimestampsField
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

func (category *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}
