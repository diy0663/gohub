package link

import (
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
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
