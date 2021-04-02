package config

import (
	"fmt"
	"go-server/config/database"
	"go-server/config/server"
	"go-server/initialize"
)

func Initialize() {
	fmt.Println("Welcome to", initialize.Application.Name)
	initRedis()
	database.InitDataBase(initialize.Database.Driver)
	server.InitServer()
}
