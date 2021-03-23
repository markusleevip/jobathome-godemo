package global

import (
	"database/sql"
	"go-server/initialize"
	"gorm.io/gorm"
	"strconv"
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
	JwtTimeout = 129600
	BaseUrl    = ""
)

func Initialize() {
	JwtSecret = initialize.Application.JwtSecret
	BaseUrl = initialize.Application.BaseUrl

	if timeout, err := strconv.Atoi(initialize.Application.JwtTimeout); err != nil {
	} else {
		JwtTimeout = timeout
	}
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
