package config

import (
	"github.com/go-redis/redis"
	"go-server/global"
	"go-server/initialize"
	"log"
)

func initRedis() {
	conf := initialize.RedisConf
	if conf.Enable != true {
		log.Printf("redisConf.enable=%v", conf.Enable)
		return
	}
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Println(err)
	} else {
		log.Println("init redis")
		global.Redis = client
	}
}
