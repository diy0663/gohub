package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/gohub/app/models/project"
)

//todo  需要手动保存之后完成自动import, 检查注意要引入的包是否正确,可能会有重名的

func MakeProjects(count int) []project.Project {

	var objs []project.Project

	// 设置唯一性，如 Project 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		projectModel := project.Project{
			// todo 补上对应字段以及faker赋值
			// todo 例如 Name:     faker.Name(),
			Name:     faker.Name(),
			Comments: faker.Sentence(),
			//FIXME()
		}
		objs = append(objs, projectModel)
	}

	return objs
}
