package global

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
	"go-server/initialize"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
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
	Redis      *redis.Client
	ResPath    = ""
	Logger     *zap.Logger
	Cache      *cache.Cache
)

func init() {
	// 默认缓存时间及清理过期缓存间隔时间
	Cache = cache.New(3*time.Minute, 10*time.Minute)
}

func Initialize() {
	JwtSecret = initialize.Application.JwtSecret
	BaseUrl = initialize.Application.BaseUrl
	ResPath = initialize.Application.ResPath

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
