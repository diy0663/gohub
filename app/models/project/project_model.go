package project

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

type Project struct {
	models.BaseModel

	// Put fields in here
	Name string `json:"name,omitempty"`

	models.CommonTimestampsField
}

func (project *Project) Create() {
	database.DB.Create(&project)
}

func (project *Project) Save() (rowsAffected int64) {
	result := database.DB.Save(&project)
	return result.RowsAffected
}

func (project *Project) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&project)
	return result.RowsAffected
}
