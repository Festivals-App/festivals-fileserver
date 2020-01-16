package main

import (
	config "github.com/Phisto/eventusfileserver/config"
	server "github.com/Phisto/eventusfileserver/server"
)

func main() {

	conf := config.GetConfig()

	imageserver := &server.Server{}
	imageserver.Initialize(conf)
	imageserver.Run(":1910")
}
