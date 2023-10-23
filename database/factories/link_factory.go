package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/gohub/app/models/link"
)

//todo  需要手动保存之后完成自动import, 检查注意要引入的包是否正确,可能会有重名的

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	// 设置唯一性，如 Link 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		linkModel := link.Link{
			// todo 补上对应字段以及faker赋值
			// todo 例如 Name:     faker.Name(),
			Name: faker.Name(),
			Url:  faker.URL(),
		}
		objs = append(objs, linkModel)
	}

	return objs
}
