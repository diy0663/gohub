package topic

import (
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

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

// 分页相关
func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(c, database.DB.Model(Topic{}), &topics, app.V1URL(database.TableName(&Topic{})), perPage)
	return
}
