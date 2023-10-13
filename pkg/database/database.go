package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// 专门用来封装 gorm的

// 定义全局变量
var DB *gorm.DB
var SQLDB *sql.DB

// 其实最主要就是使用 gorm.Open
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   _logger,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		TranslateError:                           false,

		ConnPool:  nil,
		Dialector: nil,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = SQLDB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

}
