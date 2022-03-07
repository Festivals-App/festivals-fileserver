package handler

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/rs/zerolog/log"
)

func GetLog(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	l, err := Log("/var/log/festivals-fileserver/info.log")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get log")
		respondError(w, http.StatusBadRequest, "Failed to get log")
		return
	}
	respondString(w, http.StatusOK, l)
}

func Log(location string) (string, error) {

	l, err := ioutil.ReadFile(location)
	if err != nil {
		return "", errors.New("Failed to read log file at: '" + location + "' with error: " + err.Error())
	}
	return string(l), nil
}
