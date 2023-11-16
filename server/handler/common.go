package handler

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Festivals-App/festivals-fileserver/server/manipulate"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

// respondFile makes the response with payload as json format
func respondFile(w http.ResponseWriter, file *os.File) {

	// calculate content size
	fileInfo, err := file.Stat()
	if err != nil || fileInfo == nil {
		log.Error().Err(err).Msg("Failed to read file stats for file: '" + file.Name() + "'")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	size := fileInfo.Size()

	// calculate content type dynamically
	contentType, err := manipulate.GetFileContentType(file)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve content type for file: '" + file.Name() + "'")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	w.Header().Set("Content-Type", contentType)

	_, err = io.Copy(w, file)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send write file to response")
	}
}
