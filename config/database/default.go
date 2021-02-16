package database

import (
	"go-server/app/models"
	"go-server/global"
)

// Setup 配置数据库
func InitDataBase(driver string) {

	dbType := driver
	if dbType == "mysql" {
		var db = new(Mysql)
		db.Setup()
	}

	global.GDB.AutoMigrate(&models.Account{})


}
