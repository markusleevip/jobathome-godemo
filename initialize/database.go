package initialize

import "github.com/spf13/viper"

type database struct {
	Driver string
	Source string
}

func initDatabase(cfg *viper.Viper) *database {

	db := &database{
		Driver: cfg.GetString("driver"),
		Source: cfg.GetString("source"),
	}
	return db
}

var Database = new(database)
