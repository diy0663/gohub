// Package tag 模型
package tag

import (
	"github.com/diy0663/gohub/app/models"
	"github.com/diy0663/gohub/pkg/database"
)

// todo  需要手动保存之后 自动 import 进来,同时自行检查引入的包是否正确

// 表结构定义, 不推荐使用自动迁移
// 基本用来指定单条数据增删改的数据结构
type Tag struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;unique"`

	models.CommonTimestampsField
}

// 推荐直接写明 数据表名称
func (tag *Tag) TableName() string {
	return "tags"
	// 注意校对好 表名
	//FIXME()
}

// 下面这种 Create,Save,Delete 相对来说比较局限,因为都直接写死依赖了被初始化的全局变量 database.DB ,换个库或者多个库就没法整了. 不推荐
func (tag *Tag) Create() {
	database.DB.Create(&tag)
}

// 更新
func (tag *Tag) Save() (rowsAffected int64) {
	result := database.DB.Save(&tag)
	return result.RowsAffected
}

func (tag *Tag) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&tag)
	return result.RowsAffected
}
