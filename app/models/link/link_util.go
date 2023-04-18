package link

import (
	"time"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/cache"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/helper"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// FIXME_BY_DELE_TO_IMPORT_PACKAGE()

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
	// 设置缓存 key
	cacheKey := "links:all"
	// 设置过期时间
	expireTime := 120 * time.Minute
	// 取数据
	cache.GetObject(cacheKey, &links)

	// 如果数据为空
	if helper.Empty(links) {
		// 查询数据库
		links = All()
		if helper.Empty(links) {
			return links
		}
		// 设置缓存
		cache.Set(cacheKey, links, expireTime)
	}
	return

}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页相关
func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(c, database.DB.Model(Link{}), &links, app.V1URL(database.TableName(&Link{})), perPage)
	return
}
