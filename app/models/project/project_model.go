// Package project 模型
package project

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

// todo  需要手动保存之后 自动 import 进来,同时自行检查引入的包是否正确

type Project struct {
	models.BaseModel

	// Put fields in here
	Name     string `json:"name,omitempty" gorm:"name,not null;" valid:"name"`
	Comments string `json:"comments,omitempty" gorm:"comments,not null;" valid:"comments"`

	models.CommonTimestampsField
}

func (project *Project) Create() {
	database.DB.Create(&project)
}

// 更新
func (project *Project) Save() (rowsAffected int64) {
	result := database.DB.Save(&project)
	return result.RowsAffected
}

func (project *Project) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&project)
	return result.RowsAffected
}
