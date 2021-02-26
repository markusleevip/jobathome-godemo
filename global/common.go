package global

import (
	"database/sql"
	"go-server/initialize"
	"gorm.io/gorm"
)

const (
	Version = "1.0.0"
)

var (
	DBSource   string
	DBDriver   string
	DBName     string
	GDB        *gorm.DB
	JwtSecret  string
	JwtTimeout string
)

func Initialize() {
	JwtSecret = initialize.Application.JwtSecret
	JwtTimeout = initialize.Application.JwtTimeout
}

var Cfg = DefaultConfig()

type DBConfig struct {
	Driver string
	DB     *sql.DB
}

type Config struct {
	db *DBConfig
}

// SetDb 设置单个db
func (c *Config) SetDb(db *DBConfig) {
	c.db = db
}

// GetDb 获取单个db
func (c *Config) GetDb() *DBConfig {
	return c.db
}

func DefaultConfig() *Config {
	return &Config{}
}
