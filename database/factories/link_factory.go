package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/gohub/app/models/link"
)

func MakeLinks(count int) []link.Link {
	var objs []link.Link

	// 设置唯一性，如 Link 模型的某个字段需要唯一 就得取消本注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		linkModel := link.Link{
			// 使用facker 包去构造大部分字段的数据
			// FIXME()
			//	Name: faker.Name(),
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, linkModel)
	}
	return objs
}
