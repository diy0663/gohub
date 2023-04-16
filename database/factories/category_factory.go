package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/gohub/app/models/category"
)

func MakeCategories(count int) []category.Category {
	var objs []category.Category

	// 设置唯一性，如 Category 模型的某个字段需要唯一 就得取消本注释
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		categoryModel := category.Category{
			// 使用facker 包去构造大部分字段的数据
			Name:        faker.Username(),
			Description: faker.Sentence(),
		}
		objs = append(objs, categoryModel)
	}
	return objs
}
