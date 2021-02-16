package initialize

import "github.com/spf13/viper"

type application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	JwtTimeout	  string
}

func initApplication(cfg *viper.Viper) *application {
	return &application{
		Host:          cfg.GetString("host"),
		Port:          cfg.GetString("port"),
		Name:          cfg.GetString("name"),
		JwtSecret:     cfg.GetString("jwtSecret"),
		JwtTimeout:    cfg.GetString("jwtTimeout"),
	}
}

var Application = new(application)

