package db

import (
	"fmt"
	"github.com/HouLPcode/user-management/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Mysql   *gorm.DB
}

var DB *Database

func (db *Database) init() {
	DB.Mysql = openDB(config.NewConfig())
}

func openDB(cfg *config.Config) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		cfg.MysqlConfig.UserName,
		cfg.MysqlConfig.Password,
		cfg.MysqlConfig.Host,
		cfg.MysqlConfig.Database,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		//log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(true)
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

func (db *Database) Close() {
	DB.Mysql.Close()
}

