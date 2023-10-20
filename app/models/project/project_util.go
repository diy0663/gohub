package project

import "github.com/diy0663/gohub/pkg/database"

// todo 保存之后自动import
// util 里面存的都是直接对数据表的查询操作, 是函数, 不是结构体实现的方法

func Get(idstr string) (project Project) {
	database.DB.Where("id", idstr).First(&project)
	return
}

func GetBy(field, value string) (project Project) {
	database.DB.Where("? = ?", field, value).First(&project)
	return
}

func All() (projects []Project) {
	database.DB.Find(&projects)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Project{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
