package initialize

import "github.com/spf13/viper"

type redis struct {
	DB       int
	Addr     string
	Password string
	Enable   bool
}

func initRedis(cfg *viper.Viper) *redis {
	redis := &redis{
		DB:       cfg.GetInt("db"),
		Addr:     cfg.GetString("addr"),
		Password: cfg.GetString("password"),
		Enable:   cfg.GetBool("enable"),
	}
	return redis
}

var RedisConf = new(redis)
