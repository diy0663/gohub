// Package token 模型
package token

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

// todo  需要手动保存之后 自动 import 进来,同时自行检查引入的包是否正确

type Token struct {
	models.BaseModel

	// Put fields in here
	UserID uint64 `json:"user_id,omitempty" gorm:"user_id,not null;index;" valid:"user_id"`

	TokenString string `json:"-" gorm:"token_string,not null;" valid:"token_string"`

	LoginId string `json:"login_id,omitempty" gorm:"login_id,not null;index;" valid:"login_id"`

	ExpireTime uint64 `json:"expire_time,omitempty" gorm:"expire_time,not null;" valid:"expire_time"`

	models.CommonTimestampsField
}

func (token *Token) Create() {
	database.DB.Create(&token)
}

// 更新
func (token *Token) Save() (rowsAffected int64) {
	result := database.DB.Save(&token)
	return result.RowsAffected
}

func (token *Token) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&token)
	return result.RowsAffected
}
