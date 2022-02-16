package user

import "github.com/diy0663/gohub/app/models"

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	//  - JSON 解析器忽略改字段,不要输出该字段
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
