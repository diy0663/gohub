package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/diy0663/gohub/app/models/topic"
)

func MakeTopics(count int) []topic.Topic {
	var objs []topic.Topic

	// 设置唯一性，如 Topic 模型的某个字段需要唯一 就得取消本注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			// 使用facker 包去构造大部分字段的数据
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: "3",
			UserID:     "7",
			//	Name: faker.Name(),
		}
		objs = append(objs, topicModel)
	}
	return objs
}
