package factories

import "github.com/diy0663/gohub/app/models/role"

//todo  需要手动保存之后完成自动import, 检查注意要引入的包是否正确,可能会有重名的

func MakeRoles() []role.Role {

	var objs []role.Role

	// 设置唯一性，如 Role 模型的某个字段需要唯一，即可取消注释
	//  faker.SetGenerateUniqueValues(true)
	roleMap := []string{"超级管理员", "普通管理员", "测试", "运营", "客服"}

	for _, value := range roleMap {
		roleModel := role.Role{Name: value}
		objs = append(objs, roleModel)
	}

	return objs
}
