package main

import (
	"go-server/config"
	"go-server/global"
	"go-server/initialize"
)

func main() {

	initialize.Initialize()
	global.Initialize()
	config.Initialize()

}
