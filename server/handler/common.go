package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Festivals-App/festivals-fileserver/server/manipulate"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {

	resultMap := map[string]interface{}{"data": payload}
	response, err := json.Marshal(resultMap)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to retrieve content type")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(response))
}

// respondJSON makes the response with payload as json format
func respondFile(w http.ResponseWriter, file *os.File) {

	// calculate content size
	fileInfo, err := file.Stat()
	if err != nil || fileInfo == nil {
		respondError(w, 404, "Failed to retrieve content length")
		return
	}
	size := fileInfo.Size()

	// calculate content type dynamically
	contentType, err := manipulate.GetFileContentType(file)
	if err != nil {
		respondError(w, 404, "Failed to retrieve content type")
		return
	}

	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	w.Header().Set("Content-Type", contentType)
	_, _ = io.Copy(w, file)
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	resultMap := map[string]interface{}{"error": message}
	response, err := json.Marshal(resultMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write([]byte(response))
}
