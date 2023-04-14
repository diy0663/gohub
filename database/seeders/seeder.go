package seeders

import "github.com/diy0663/gohub/pkg/seed"

func Initialize() {
	//  触发所在包内的所有init 方法

	// 指定顺序
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})

}
