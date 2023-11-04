package role

import (
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// todo 保存之后自动import
// util 里面存的都是直接对数据表的查询操作, 是函数, 不是结构体实现的方法

func Get(idstr string) (role Role) {
	database.DB.Where("id", idstr).First(&role)
	return
}

func GetByName(value string) (role Role) {
	database.DB.Where("name= ?", value).First(&role)
	return
}

func All() (roles []Role) {
	database.DB.Find(&roles)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Role{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页列表查询
func Paginate(c *gin.Context, perPage int) (roles []Role, paging paginator.Paging) {

	query := database.DB.Model(&Role{})

	// 条件查询

	//	if name, isExists := c.GetQuery("name"); isExists {
	//	query = query.Where("name=?", string(name))
	//	}

	paging = paginator.Paginate(
		c,
		query,
		&roles,
		app.V1URL(database.TableNameByStruct(Role{})),
		perPage,
	)

	return
}
