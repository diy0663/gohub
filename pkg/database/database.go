package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// 使用gorm

var DB *gorm.DB

// todo ? ???
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

}
