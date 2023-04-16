package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"github.com/diy0663/gohub/app/models/category"
	"github.com/diy0663/gohub/app/models/project"
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() {

	// 根据.env 里配置的默认数据库类型,做数据库连接处理
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database of " + config.Get("database.connection") + " connection is  not supported"))
	}
	// 传入配置即可用gorm的方式去开启连接,第二个参数是logger ,可以用gorm自带的logger, 也可以再封装一个基于zap的专用数据库logger,要注意实现 Interface 这个接口

	//database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	database.Connect(dbConfig, logger.NewGormLogger())

	//todo 如何判断数据库连接是否真的成功

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	//本地开发环境才允许使用数据库自动迁移,生产环境不推荐,因为有风险, 且生成的字段类型取值长度并不准确
	if config.Get("app.env") == "local" {
		database.DB.AutoMigrate(&user.User{})
		database.DB.AutoMigrate(&project.Project{})
		database.DB.AutoMigrate(&category.Category{})
	}
}
