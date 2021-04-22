package initialize

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func Initialize() {
	var config string
	flag.StringVar(&config, "c", "config.yml", "choose config file.")
	flag.Parse()

	v := viper.New()
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	cfgApp := v.Sub("application")
	if cfgApp == nil {
		panic("No found application in the configuration")
	}
	Application = initApplication(cfgApp)

	cfgDatabase := v.Sub("database")
	if cfgDatabase == nil {
		panic("No found database in the configuration")
	}
	Database = initDatabase(cfgDatabase)

	if cfgRedis := v.Sub("redis"); cfgRedis != nil {
		RedisConf = initRedis(cfgRedis)
	}
}

