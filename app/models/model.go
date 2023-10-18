package models

import (
	"time"

	"github.com/spf13/cast"
)

type BaseModel struct {
	ID uint64 `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement;" valid:"id"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at;index;" valid:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;index;" valid:"updated_at"`
}

func (model *BaseModel) GetStringID() string {
	return cast.ToString(model.ID)
}
