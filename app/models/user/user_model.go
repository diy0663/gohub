package user

import (
	"gohub/app/models"

	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
