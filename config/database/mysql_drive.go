package database

import (
	"database/sql"
	"go-server/global"
	"go-server/initialize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// Mysql mysql配置结构体
type Mysql struct {
}

func (e *Mysql) Setup() {

	global.DBSource = e.GetConnect()
	db, err := sql.Open("mysql", e.GetConnect())
	if err != nil {
		log.Fatal(err)
	}
	global.Cfg.SetDb(&global.DBConfig{
		Driver: e.GetDriver(),
		DB:     db,
	})
	global.GDB, err = e.Open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if global.GDB.Error != nil {
		log.Fatal(global.GDB.Error)
	}

}

// Open 打开数据库连接
func (e *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// GetConnect 获取数据库连接
func (e *Mysql) GetConnect() string {
	return initialize.Database.Source
}

// GetDriver 获取连接
func (e *Mysql) GetDriver() string {
	return initialize.Database.Driver
}
