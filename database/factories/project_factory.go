package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/gohub/app/models/project"
)

func MakeProjects(count int) []project.Project {
	var objs []project.Project

	// 设置唯一性，如 Project 模型的某个字段需要唯一 就得取消本注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		projectModel := project.Project{
			// 使用facker 包去构造大部分字段的数据
			Name: faker.Name(),
		}
		objs = append(objs, projectModel)
	}
	return objs
}
