package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"github.com/diy0663/go_project_packages/config"
	grom_logger "github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/gohub/pkg/database"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// 初始化数据库链接,被main 使用
func SetupDB() {

	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		// 构建 DSN 信息
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

	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	// database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	database.Connect(dbConfig, grom_logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 迁移生成数据表结构
	// database.DB.AutoMigrate(&userModel.User{})
	//database.DB.AutoMigrate(&project.Project{})
}
