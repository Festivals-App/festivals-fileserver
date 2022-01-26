package main

import (
	"os"
	"strconv"
	"time"

	"github.com/Festivals-App/festivals-fileserver/server"
	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/Festivals-App/festivals-gateway/server/heartbeat"
)

func main() {
	conf := config.DefaultConfig()
	if len(os.Args) > 1 {
		conf = config.ParseConfig(os.Args[1])
	}

	serverInstance := &server.Server{}
	serverInstance.Initialize(conf)
	go sendHeartbeat(conf)
	serverInstance.Run(conf.ServiceBindAddress + ":" + strconv.Itoa(conf.ServicePort))
}

func sendHeartbeat(conf *config.Config) {
	for {
		timer := time.After(time.Second * 2)
		<-timer
		var beat *heartbeat.Heartbeat = &heartbeat.Heartbeat{Service: "festivals-fileserver", Host: conf.ServiceBindAddress, Port: conf.ServicePort, Available: true}
		heartbeat.SendHeartbeat(conf.LoversEar, beat)
	}
}
