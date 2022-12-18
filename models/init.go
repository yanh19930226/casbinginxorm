package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 数据库连接
var DB *xorm.Engine

func Setup() {
	var err error
	//数据库链接
	DB, err = xorm.NewEngine("mysql", "root:66^^66@tcp(121.43.34.54:3306)/rbac_db?charset=utf8")
	if err != nil {
		log.Printf("创建DB连接错误: %v\n", err)
		return
	}

	// 控制台打印SQL语句
	DB.ShowSQL(true)
	DB.Logger().SetLevel(2)
}
