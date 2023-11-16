package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetLog(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	l, err := Log("/var/log/festivals-fileserver/info.log")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get log")
		servertools.RespondError(w, http.StatusBadRequest, "Failed to get log")
		return
	}
	servertools.RespondString(w, http.StatusOK, l)
}

func GetTraceLog(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	l, err := Log("/var/log/festivals-fileserver/trace.log")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get log")
		servertools.RespondError(w, http.StatusBadRequest, "Failed to get log")
		return
	}
	servertools.RespondString(w, http.StatusOK, l)
}

func Log(location string) (string, error) {

	l, err := os.ReadFile(location)
	if err != nil {
		return "", errors.New("Failed to read log file at: '" + location + "' with error: " + err.Error())
	}
	return string(l), nil
}
