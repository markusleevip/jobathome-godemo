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
	global.GDB.AutoMigrate(&models.Menu{})
	global.GDB.AutoMigrate(&models.ProjectExp{})
	global.GDB.AutoMigrate(&models.JobExp{})
	global.GDB.AutoMigrate(&models.Education{})
	global.GDB.AutoMigrate(&models.Resume{})
	global.GDB.AutoMigrate(&models.Friend{})

}
