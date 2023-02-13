package user

import "github.com/diy0663/gohub/app/models"

type User struct {
	models.BaseModel
	Name string `json:"name,omitempty"`
	// - 表示忽略,不输出此类敏感信息
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
	models.CommonTimestampsField
}
