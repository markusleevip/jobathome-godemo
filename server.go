package main

import (
	"fmt"
	"go-server/config"
	"go-server/global"
	"go-server/initialize"
)

func main() {

	initialize.Initialize()
	global.Initialize()
	config.Initialize()
	fmt.Printf("%v",global.Cfg.GetDb())
	
}
