package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
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

func TableNameByStruct(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}

// 另一种数据库链接的写法,外层可以在最开始的时候初始化,定义全局变量配合sync.Once 使用
func MustNewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, *sql.DB) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), opts...)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
	// todo 在这里暂时不设置最大连接数,最大空闲连接数,链接连接时间
	return db, sqlDB
}
