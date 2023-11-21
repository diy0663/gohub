package category

import (
	"context"
	"fmt"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryModel struct {
	db *gorm.DB
}

// 在使用 NewCategoryModel 的时候可以把 database.DB 传进来
func NewCategoryModel(db *gorm.DB) *CategoryModel {
	return &CategoryModel{
		//
		db: db,
	}
}

func (category *CategoryModel) FindOne(ctx context.Context, id int64) (*Category, error) {
	var result Category
	err := category.db.WithContext(ctx).Where("id =?", id).First(&result).Error
	return &result, err
}

func (category *CategoryModel) Insert(ctx context.Context, data *Category) error {
	return category.db.WithContext(ctx).Create(data).Error
}

func (category *CategoryModel) Update(ctx context.Context, data *Category) error {
	return category.db.WithContext(ctx).Save(data).Error
}

func (category *CategoryModel) UpdateFiels(ctx context.Context, id int64, data map[string]interface{}) error {
	return category.db.WithContext(ctx).Model(&Category{}).Where("id =?", id).Updates(data).Error
}

func (category *CategoryModel) FindByIds(ctx context.Context, ids []int64) ([]Category, error) {
	var results []Category
	err := category.db.WithContext(ctx).Where("id IN (?)", ids).Find(&results).Error
	return results, err
}

// 原生sql 进行 操作类操作
// func (category *CategoryModel) DoExec(cxt context.Context) error {
// 	return category.db.WithContext(cxt).Exec("update XX from XX where XX =? ", 1).Error
// }

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
func Paginate(c *gin.Context, perPage int) (categories []*Category, paging paginator.Paging) {

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

func DeleteWithTopic(categoryId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var category Category
		err := tx.Where("id=?", categoryId).First(&category).Error
		if err != nil {
			return err
		}

		err = tx.Exec("delete from topics where category_id = ? ", categoryId).Error
		return err

	})
	return err
}
