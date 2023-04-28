package user

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/hash"
)

type User struct {
	models.BaseModel
	Name string `gorm:"type:varchar(50);index;" json:"name,omitempty"`

	City         string `gorm:"type:varchar(50);" json:"city,omitempty"`
	Introduction string `gorm:"type:varchar(255);" json:"introduction,omitempty"`
	Avatar       string `gorm:"type:varchar(255);default:null" json:"avatar,omitempty"`

	// - 表示忽略,不输出此类敏感信息
	Email    string `gorm:"type:varchar(30);index;" json:"-"`
	Phone    string `gorm:"type:varchar(50);index;" json:"-"`
	Password string `gorm:"type:varchar(100);" json:"-"`
	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowAffected int64) {
	// todo 这里为啥要传  &userModel ??
	result := database.DB.Save(&userModel)
	// 返回影响行数
	return result.RowsAffected
}
