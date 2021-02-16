package database

import (
	"database/sql"
	"gorm.io/gorm"
)

// Database 数据库配置
type Database interface {
	Setup()
	Open(conn  *sql.DB, cfg *gorm.Config) (db *gorm.DB, err error)
	GetConnect() string
	GetDriver() string
}

