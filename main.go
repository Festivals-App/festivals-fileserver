package main

import (
	server "github.com/Festivals-App/festivals-fileserver/server"
	config "github.com/Festivals-App/festivals-fileserver/server/config"
	"os"
	"strconv"
)

func main() {

	conf := config.DefaultConfig()
	if len(os.Args) > 1 {
		conf = config.ParseConfig(os.Args[1])
	}

	imageserver := &server.Server{}
	imageserver.Initialize(conf)
	imageserver.Run(":" + strconv.Itoa(conf.ServicePort))
}
