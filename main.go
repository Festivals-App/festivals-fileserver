package main

import (
	config "github.com/Festivals-App/festivals-fileserver/config"
	server "github.com/Festivals-App/festivals-fileserver/server"
)

func main() {

	conf := config.GetConfig()

	imageserver := &server.Server{}
	imageserver.Initialize(conf)
	imageserver.Run(":1910")
}
