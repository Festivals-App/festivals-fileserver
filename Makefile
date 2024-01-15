# Makefile for festivals-fileserver

VERSION=development
DATE=$(shell date +"%d-%m-%Y-%H-%M")
REF=refs/tags/development
export

build:
	go build -ldflags="-X 'github.com/Festivals-App/festivals-fileserver/server/status.ServerVersion=$(VERSION)' -X 'github.com/Festivals-App/festivals-fileserver/server/status.BuildTime=$(DATE)' -X 'github.com/Festivals-App/festivals-fileserver/server/status.GitRef=$(REF)'" -o festivals-fileserver main.go

install:
	cp festivals-fileserver /usr/local/bin/festivals-fileserver
	cp config_template.toml /etc/festivals-fileserver.conf
	mkdir -p /srv/festivals-fileserver/images/resized
	mkdir -p /srv/festivals-fileserver/pdf
	cp operation/service_template.service /etc/systemd/system/festivals-fileserver.service

update:
	systemctl stop festivals-fileserver
	cp festivals-server /usr/local/bin/festivals-fileserver
	systemctl start festivals-fileserver

uninstall:
	systemctl stop festivals-fileserver
	rm /usr/local/bin/festivals-fileserver
	rm /etc/festivals-fileserver.conf
	rm /etc/systemd/system/festivals-fileserver.service
	rm -r /srv/festivals-fileserver

run:
	./festivals-fileserver

stop:
	killall festivals-fileserver

clean:
	rm -r festivals-fileserver