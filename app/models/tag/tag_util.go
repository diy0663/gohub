package tag

import (
	"context"
	"errors"
	"fmt"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// util 里面存的都是直接对数据表的查询操作, 是函数, 不是结构体实现的方法

func Get(idstr string) (tag Tag) {
	database.DB.Where("id", idstr).First(&tag)
	return
}

func GetBy(field, value string) (tag Tag) {

	str := fmt.Sprintf("%v= ?", field)
	database.DB.Where(str, value).First(&tag)
	return
}

func All() (tags []Tag) {
	database.DB.Find(&tags)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Tag{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页列表查询
func Paginate(c *gin.Context, perPage int) (tags []*Tag, paging paginator.Paging) {

	query := database.DB.Model(&Tag{})

	// 条件查询

	//	if name, isExists := c.GetQuery("name"); isExists {
	//	query = query.Where("name=?", string(name))
	//	}

	paging = paginator.Paginate(
		c,
		query,
		&tags,
		app.V1URL(database.TableNameByStruct(Tag{})),
		perPage,
	)

	return
}

// 以下是不依赖 database.DB 的写法
type TagModel struct {
	db *gorm.DB
}

// 在使用 NewCategoryModel 的时候可以把 经过初始化的*gorm.DB传进来
func NewTagModel(db *gorm.DB) *TagModel {
	if db == nil {
		panic("NewTagModel , db is nil")
	}
	return &TagModel{
		//
		db: db,
	}
}

func (tag *TagModel) FindOne(ctx context.Context, id int64) (*Tag, error) {
	var result Tag
	err := tag.db.WithContext(ctx).Where("id =?", id).First(&result).Error
	return &result, err
}

func (tag *TagModel) Insert(ctx context.Context, data *Tag) error {
	return tag.db.WithContext(ctx).Create(data).Error

}

func (tag *TagModel) Update(ctx context.Context, data *Tag) error {
	return tag.db.WithContext(ctx).Save(data).Error
}

func (tag *TagModel) UpdateFiels(ctx context.Context, id int64, data map[string]interface{}) error {
	return tag.db.WithContext(ctx).Model(&Tag{}).Where("id =?", id).Updates(data).Error
}

func (tag *TagModel) FindByIds(ctx context.Context, ids []int64) ([]Tag, error) {
	var results []Tag
	err := tag.db.WithContext(ctx).Where("id IN (?)", ids).Find(&results).Error
	return results, err
}

// 原生sql 进行 操作类操作
// func (tag *TagModel) DoExec(cxt context.Context) error {
// 	return tag.db.WithContext(cxt).Exec("update XX from XX where XX =? ", 1).Error
// }

// 事务处理示例,注意事务不能跨库
func (tag *TagModel) TransactionDeal(id string) error {
	err := tag.db.Transaction(func(tx *gorm.DB) error {
		var result Tag
		err := tx.Where("id=?", id).First(&result).Error
		if err != nil {
			//return err
			return errors.New("查询出错")
		}

		err = tx.Exec("delete from XX where id = ? ", id).Error
		return err

	})
	return err
}
