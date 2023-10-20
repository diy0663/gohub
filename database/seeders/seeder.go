package seeders

import "github.com/diy0663/gohub/pkg/seed"

func InitAllSeeder() {

	// 触发本目录的其他init
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
