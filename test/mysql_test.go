package test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestMysql(t *testing.T) {
	db, err := gorm.Open("mysql", "root:000@tcp(127.0.0.1:3306)/db_apiserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		t.Log("连接mysql数据库失败")
	}
	b := db.HasTable("tb_users")
	t.Log(b)
	defer db.Close()
}
