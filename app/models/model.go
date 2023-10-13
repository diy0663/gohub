package models

import "time"

type BaseModel struct {
	ID uint64 `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement;" valid:"id"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at;index;" valid:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;index;" valid:"updated_at"`
}
