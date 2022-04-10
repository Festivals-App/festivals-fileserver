package handler

import (
	"net/http"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/Festivals-App/festivals-gateway/server/update"
)

func MakeUpdate(conf *config.Config, w http.ResponseWriter, _ *http.Request) {

	respondCode(w, http.StatusAccepted)
	go update.RunUpdate("https://github.com/Festivals-App/festivals-fileserver/releases/latest", "/usr/local/festivals-fileserver/update.sh")
}
