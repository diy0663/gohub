package category

import (
	"fmt"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// todo 保存之后自动import
// util 里面存的都是直接对数据表的查询操作, 是函数, 不是结构体实现的方法

func Get(idstr string) (category Category) {
	database.DB.Where("id", idstr).First(&category)
	return
}

func GetBy(field, value string) (category Category) {
	str := fmt.Sprintf("%v= ?", field)
	database.DB.Where(str, value).First(&category)
	return
}

func All() (categories []Category) {
	database.DB.Find(&categories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页列表查询
func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {

	query := database.DB.Model(&Category{})

	// 条件查询

	//	if name, isExists := c.GetQuery("name"); isExists {
	//	query = query.Where("name=?", string(name))
	//	}

	paging = paginator.Paginate(
		c,
		query,
		&categories,
		app.V1URL(database.TableNameByStruct(Category{})),
		perPage,
	)

	return
}
