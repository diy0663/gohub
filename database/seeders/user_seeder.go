package seeders

import (
	"fmt"

	"github.com/diy0663/gohub/database/factories"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/diy0663/gohub/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	// 注意,第二个参数只是定义, 并不是执行
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		users := factories.MakeUsers(12)

		result := db.Create(&users)
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))

	})
}
