package initialize

import "github.com/spf13/viper"

type application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	JwtTimeout    string
	BaseUrl       string
	ResPath       string
}

func initApplication(cfg *viper.Viper) *application {
	return &application{
		Host:       cfg.GetString("host"),
		Port:       cfg.GetString("port"),
		Name:       cfg.GetString("name"),
		JwtSecret:  cfg.GetString("jwtSecret"),
		JwtTimeout: cfg.GetString("jwtTimeout"),
		BaseUrl:    cfg.GetString("baseUrl"),
		ResPath:    cfg.GetString("resPath"),
	}
}

var Application = new(application)
