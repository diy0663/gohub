package link

import (
	"time"

	"github.com/diy0663/go_project_packages/helper"
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/cache"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// todo 保存之后自动import
// util 里面存的都是直接对数据表的查询操作, 是函数, 不是结构体实现的方法

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}
func AllCached() (links []Link) {
	cacheKey := "links:all"
	// 129分钟的缓存
	expireTime := 120 * time.Minute

	cache.GetObject(cacheKey, &links)
	if helper.Empty(links) {
		links = All()
		cache.Set(cacheKey, links, expireTime)
	}
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页列表查询
func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {

	query := database.DB.Model(&Link{})

	// 条件查询

	//	if name, isExists := c.GetQuery("name"); isExists {
	//	query = query.Where("name=?", string(name))
	//	}

	paging = paginator.Paginate(
		c,
		query,
		&links,
		app.V1URL(database.TableNameByStruct(Link{})),
		perPage,
	)

	return
}
