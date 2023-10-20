package seeders

import (
	"fmt"

	"github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/gohub/database/factories"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	// 添加 到  存放所有seeder 的 全局变量 seeders
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		usersData := factories.MakeUsers(10)
		// todo  批量创建用户（注意批量创建不会调用模型钩子） ??
		result := db.Table("users").Create(&usersData)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
