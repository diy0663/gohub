package topic

import (
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// todo 保存之后自动import
// util 里面存的都是直接对数据表的查询操作, 是函数, 不是结构体实现的方法

func Get(idstr string) (topic Topic) {
	database.DB.Where("id", idstr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页列表查询
func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {

	query := database.DB.Model(&Topic{})

	// 条件查询

	//	if name, isExists := c.GetQuery("name"); isExists {
	//	query = query.Where("name=?", string(name))
	//	}

	paging = paginator.Paginate(
		c,
		query,
		&topics,
		app.V1URL(database.TableNameByStruct(Topic{})),
		perPage,
	)

	return
}
