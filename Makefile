# Makefile for festivals-fileserver

VERSION=development
DATE=$(shell date +"%d-%m-%Y-%H-%M")
REF=refs/tags/development
DEV_PATH_MAC=$(shell echo ~/Library/Containers/org.festivalsapp.project)
export

build:
	go build -ldflags="-X 'github.com/Festivals-App/festivals-fileserver/server/status.ServerVersion=$(VERSION)' -X 'github.com/Festivals-App/festivals-fileserver/server/status.BuildTime=$(DATE)' -X 'github.com/Festivals-App/festivals-fileserver/server/status.GitRef=$(REF)'" -o festivals-fileserver main.go

install:

	mkdir -p $(DEV_PATH_MAC)/usr/local/bin
	mkdir -p $(DEV_PATH_MAC)/etc
	mkdir -p $(DEV_PATH_MAC)/var/log
	mkdir -p $(DEV_PATH_MAC)/usr/local/festivals-fileserver
	mkdir -p $(DEV_PATH_MAC)/srv/festivals-fileserver/images/resized
	mkdir -p $(DEV_PATH_MAC)/srv/festivals-fileserver/pdf

	cp operation/local/ca.crt  $(DEV_PATH_MAC)/usr/local/festivals-fileserver/ca.crt
	cp operation/local/server.crt  $(DEV_PATH_MAC)/usr/local/festivals-fileserver/server.crt
	cp operation/local/server.key  $(DEV_PATH_MAC)/usr/local/festivals-fileserver/server.key
	cp festivals-fileserver $(DEV_PATH_MAC)/usr/local/bin/festivals-fileserver
	chmod +x $(DEV_PATH_MAC)/usr/local/bin/festivals-fileserver
	cp operation/local/config_template_dev.toml $(DEV_PATH_MAC)/etc/festivals-fileserver.conf

run:
	./festivals-fileserver --container="$(DEV_PATH_MAC)"

run-dev:
	$(DEV_PATH_MAC)/usr/local/bin/festivals-identity-server --container="$(DEV_PATH_MAC)" &
	$(DEV_PATH_MAC)/usr/local/bin/festivals-gateway --container="$(DEV_PATH_MAC)" &

stop-dev:
	killall festivals-gateway
	killall festivals-identity-server

stop:
	killall festivals-fileserver

clean:
	rm -r festivals-fileserver